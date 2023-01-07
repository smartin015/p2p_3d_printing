package server

import (
  "context"
  "time"
  "fmt"
  "google.golang.org/protobuf/proto"
  pb "github.com/smartin015/peerprint/p2pgit/pkg/proto"
	"github.com/smartin015/peerprint/p2pgit/pkg/transport"
	"github.com/smartin015/peerprint/p2pgit/pkg/storage"
	"github.com/smartin015/peerprint/p2pgit/pkg/log"
)

const (
  PeerPrintProtocol = "peerprint@0.0.1"
  DefaultTopic = "default"
  SyncPeriod = 5 * time.Second
  TargetAdminCount = 3
)

type Opts struct {
  AccessionDelay time.Duration
  StatusPeriod time.Duration
}

type Interface interface {
  GetService() interface{}
  Run(context.Context)
  WaitUntilReady()
  SetRecord(r *pb.Record) error
  SetGrant(g *pb.Grant) error
  ID() string
}

type Server struct {
  opts Opts
  t transport.Interface
  s storage.Interface
  mirror string
  l *log.Sublog
  roleChanged chan struct{}

  // Tickers for periodic network activity
  publishStatusTicker *time.Ticker
  status *pb.PeerStatus

  // Data for specific states
  handshake *handshake
  listener *listener
  electable *electable
  leader *leader
}

func New(t transport.Interface, s storage.Interface, opts *Opts, l *log.Sublog) *Server {
  srv := &Server{
    t: t,
    s: s,
    l: l,
    roleChanged: make(chan struct{}),
    publishStatusTicker: time.NewTicker(opts.StatusPeriod),
    status: &pb.PeerStatus{
      Type: pb.PeerType_UNKNOWN_PEER_TYPE, // Unknown until handshake is complete
    },
  }
  srv.handshake = &handshake{
    base: srv,
    accessionDelay: opts.AccessionDelay,
    l: log.New("Handshake", l),
  }
  srv.listener = &listener{
    base: srv,
    l: log.New("Listener", l),
  }
  srv.electable = &electable{
    base: srv,
    l: log.New("Electable", l),
  }
  srv.leader = &leader{
    base: srv,
    ticker: time.NewTicker(Heartbeat),
    l: log.New("Leader", l),
  }

  if err := s.SetPubKey(t.ID(), t.PubKey()); err != nil {
    panic(fmt.Errorf("Failed to store our pubkey: %w", err))
  }

  return srv
}

func (s *Server) GetService() interface{} {
  return &PeerPrintService{
    base: s,
  }
}

func (s *Server) ID() string {
  return s.t.ID()
}

func (s *Server) SetGrant(g *pb.Grant) error {
  if ok, err := s.s.IsAdmin(s.ID()); err != nil && ok {
    if sig, err := s.sign(g); err != nil {
      return err
    } else if err := s.storeGrantFromPeer(s.ID(), g, sig); err != nil {
      return err
    }
  }
  return s.t.Publish(DefaultTopic, g)
}

func (s *Server) SetRecord(r *pb.Record) error {
  if ok, err := s.s.IsAdmin(s.ID()); err != nil && ok {
    if sig, err := s.sign(r); err != nil {
      return err
    } else if err := s.storeRecordFromPeer(s.ID(), r, sig); err != nil {
      return err
    }
  }
  return s.t.Publish(DefaultTopic, r)
}

func (s *Server) sendStatus() error {
  return s.t.Publish(StatusTopic, s.status)
}

func (s *Server) changeRole(t pb.PeerType) {
  s.status.Type = t
  select {
  case s.roleChanged<- struct{}{}:
  default:
  }
}

func (s *Server) partialSync() {
  n, err := s.t.GetRandomNeighbor()
  if err != nil {
    s.l.Error("GetRandomNeighbor: %w", err)
    return
  }
  rep := &pb.GetStateResponse{}
  if err := s.t.Call(n, "GetState", &pb.GetStateRequest{}, rep); err != nil {
    s.l.Error("GetState: %w", err)
    return
  }
  cnt := 0
  nerr := 0
  for _, r := range rep.Records {
    if err := s.s.SetSignedRecord(r); err != nil {
      s.l.Error("SetSignedRecord: %w", err)
      nerr += 1
    } else {
      cnt += 1
    }
  }
  for _, g := range rep.Grants {
    if err := s.s.SetSignedGrant(g); err != nil {
      s.l.Error("SetSignedGrant: %w", err)
      nerr += 1
    } else {
      cnt += 1
    }
  }
}

func (s *Server) sign(m proto.Message) (*pb.Signature, error) {
  if b, err := proto.Marshal(m); err != nil {
    return nil, fmt.Errorf("sign() marshal error: %w", err)
  } else if sig, err := s.t.Sign(b); err != nil {
    return nil, fmt.Errorf("sign() crypto error: %w", err)
  } else {
    return &pb.Signature{
      Signer: s.t.ID(),
      Data: sig,
    }, nil
  }
}

func (s *Server) verify(m proto.Message, sig *pb.Signature) (bool, error) {
  k, err := s.s.GetPubKey(sig.Signer)
  if err != nil {
    return false, fmt.Errorf("verify() get key error: %w", err)
  }
  b, err := proto.Marshal(m)
  if err != nil {
    return false, fmt.Errorf("verify() marshal error: %w", err)
  }

  return k.Verify(b, sig.Data)
}

func (s *Server) storeGrantFromPeer(peer string, g *pb.Grant, sig *pb.Signature) error {
  if peerAdmin, err := s.s.IsAdmin(peer); err != nil {
    return fmt.Errorf("IsAdmin error: %w", err)
  } else if !peerAdmin {
    s.l.Info("Ignoring grant from non-admin peer %s", peer)
    return nil
  }
  if err := s.s.SetSignedGrant(&pb.SignedGrant{
    Grant: g,
    Signature: sig,
  }); err != nil {
    return fmt.Errorf("SetSignedGrant error: %w", err)
  }
  s.l.Info("Stored grant (from %s)", peer)
  return nil
}

func (s *Server) storeRecordFromPeer(peer string, r *pb.Record, sig *pb.Signature) error {
  if peerAdmin, err := s.s.IsAdmin(peer); err != nil {
    return fmt.Errorf("IsAdmin error: %w", err)
  } else if !peerAdmin {
    s.l.Info("Ignoring record from non-admin peer %s", peer)
    return nil
  }
  if err := s.s.SetSignedRecord(&pb.SignedRecord{
    Record: r,
    Signature: sig,
  }); err != nil {
    return fmt.Errorf("SetSignedRecord error: %w", err)
  }
  s.l.Info("Stored record (from %s)", peer)
  return nil
}

func (s *Server) Run(ctx context.Context) {
  s.t.Run(ctx)
  s.l.Info("Running server main loop")

  s.handshake.Init()
  for {
    switch s.status.Type {
    case pb.PeerType_UNKNOWN_PEER_TYPE:
      s.handshake.Step(ctx)
    case pb.PeerType_LEADER:
      s.leader.Step(ctx)
    case pb.PeerType_ELECTABLE:
      s.electable.Step(ctx)
    case pb.PeerType_LISTENER:
      s.listener.Step(ctx)
    default:
      panic("Unknown status type")
    }
  }
}

func (s *Server) WaitUntilReady() {
  for {
    select {
    case <-s.roleChanged:
      if s.status.Type != pb.PeerType_UNKNOWN_PEER_TYPE {
        return
      }
    }
  }
}
