// package prpc sends and receives proto messages over pubsub
package prpc 

import (
  "fmt"
  "context"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
  pb "github.com/smartin015/peerprint/pubsub/proto"
  "google.golang.org/protobuf/proto"
)

type PRPC struct {
  ID string
  ps *pubsub.PubSub
  topics map[string]*pubsub.Topic
  cbs Callbacks
}

func New(id string, ps *pubsub.PubSub) *PRPC {
  return &PRPC{
    ID: id,
    ps: ps,
    topics: make(map[string]*pubsub.Topic),
    cbs: nil,
  }
}

func (p *PRPC) RegisterCallbacks(c Callbacks) {
  p.cbs = c
}

func (p *PRPC) handleSub(ctx context.Context, sub *pubsub.Subscription) {
	for {
		m, err := sub.Next(ctx)
		if err != nil {
			panic(err)
		}
    peer := m.ReceivedFrom.String()
    if peer == p.ID {
      continue // Ignore messages coming from ourselves
    }
    var msg proto.Message
    if err := proto.Unmarshal(m.Message.Data, msg); err != nil {
      panic(err)
    }
    p.Recv(ctx, sub.Topic(), peer, msg)
	}
}

func (p *PRPC) JoinTopic(ctx context.Context, topic string) error {
  if _, ok := p.topics[topic]; ok {
    return fmt.Errorf("Already subscribed to topic %v", topic)
  }
  t, err := p.ps.Join(topic)
  if err != nil {
    return fmt.Errorf("pubsub.Join() failure: %w", err)
  }
  sub, err := t.Subscribe()
  if err != nil {
    panic(err)
  }
  p.topics[topic] = t
  go p.handleSub(ctx, sub)
  return nil
}

func (p *PRPC) LeaveTopic(topic string) error {
  return fmt.Errorf("Unimplemented")
}

func (p *PRPC) Close() error {return nil}

func (p *PRPC) Publish(ctx context.Context, topic string, req interface{}) error {
  reqp, ok := req.(proto.Message)
  if !ok {
    return fmt.Errorf("prpc.Publish() type assertion failed")
  }
  msg, err := proto.Marshal(reqp)
  if err != nil {
    return fmt.Errorf("prpc.Publish() marshal error:", err)
  }
	t, ok := p.topics[topic]
	if !ok {
		return fmt.Errorf("attempted to publish to topic %s without first calling JoinTopic()", topic)
	}

  if err = t.Publish(ctx, msg); err != nil {
    return fmt.Errorf("prpc.Publish() publish error:", err)
  }
  return nil
}


// Callbacks contains callback functions for handling various incoming pubsub messages. It's up to the callee to determine whether they should handle the message.
type Callbacks interface {
  OnPollPeersRequest(string, string, *pb.PollPeersRequest)
  OnPollPeersResponse(string, string, *pb.PollPeersResponse)
  OnAssignmentRequest(string, string, *pb.AssignmentRequest)
  OnAssignmentResponse(string, string, *pb.AssignmentResponse)
  OnSetJobRequest(string, string, *pb.SetJobRequest)
  OnDeleteJobRequest(string, string, *pb.DeleteJobRequest)
  OnAcquireJobRequest(string, string, *pb.AcquireJobRequest)
  OnReleaseJobRequest(string, string, *pb.ReleaseJobRequest)
  OnJobMutationResponse(string, string, *pb.JobMutationResponse)
  OnGetJobsRequest(string, string, *pb.GetJobsRequest)
  OnGetJobsResponse(string, string, *pb.GetJobsResponse)
}

func (p *PRPC) Recv(ctx context.Context, topic string, peer string, resp interface{}) error {
  switch v := resp.(type) {
    // proto/peers.proto
    case *pb.AssignmentRequest:
      p.cbs.OnAssignmentRequest(topic, peer, v)
    case *pb.AssignmentResponse:
      p.cbs.OnAssignmentResponse(topic, peer, v)
    case *pb.PollPeersRequest:
      p.cbs.OnPollPeersRequest(topic, peer, v)
    case *pb.PollPeersResponse:
      p.cbs.OnPollPeersResponse(topic, peer, v)

    // proto/jobs.proto
    case *pb.SetJobRequest:
      p.cbs.OnSetJobRequest(topic, peer, v)
    case *pb.DeleteJobRequest:
      p.cbs.OnDeleteJobRequest(topic, peer, v)
    case *pb.AcquireJobRequest:
      p.cbs.OnAcquireJobRequest(topic, peer, v)
    case *pb.ReleaseJobRequest:
      p.cbs.OnReleaseJobRequest(topic, peer, v)
    case *pb.JobMutationResponse:
      p.cbs.OnJobMutationResponse(topic, peer, v)
    case *pb.GetJobsRequest:
      p.cbs.OnGetJobsRequest(topic, peer, v)
    case *pb.GetJobsResponse:
      p.cbs.OnGetJobsResponse(topic, peer, v)

    default:
      return fmt.Errorf("Received unknown type response on topic")
  }
  return nil
}
