// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/peers.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PeerType int32

const (
	PeerType_UNKNOWN_PEER_TYPE PeerType = 0
	PeerType_ELECTABLE         PeerType = 2 // This peer is involved in RAFT consensus and log replication
	PeerType_LISTENER          PeerType = 3 // Not involved in consensus
)

// Enum value maps for PeerType.
var (
	PeerType_name = map[int32]string{
		0: "UNKNOWN_PEER_TYPE",
		2: "ELECTABLE",
		3: "LISTENER",
	}
	PeerType_value = map[string]int32{
		"UNKNOWN_PEER_TYPE": 0,
		"ELECTABLE":         2,
		"LISTENER":          3,
	}
)

func (x PeerType) Enum() *PeerType {
	p := new(PeerType)
	*p = x
	return p
}

func (x PeerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_peers_proto_enumTypes[0].Descriptor()
}

func (PeerType) Type() protoreflect.EnumType {
	return &file_proto_peers_proto_enumTypes[0]
}

func (x PeerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerType.Descriptor instead.
func (PeerType) EnumDescriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{0}
}

type PeerState int32

const (
	PeerState_UNKNOWN_PEER_STATE PeerState = 0
	PeerState_BUSY               PeerState = 1 // Peer is doing work
	PeerState_IDLE               PeerState = 2 // Peer is ready to work
	PeerState_PAUSED             PeerState = 3 // Peer is awaiting human touch
)

// Enum value maps for PeerState.
var (
	PeerState_name = map[int32]string{
		0: "UNKNOWN_PEER_STATE",
		1: "BUSY",
		2: "IDLE",
		3: "PAUSED",
	}
	PeerState_value = map[string]int32{
		"UNKNOWN_PEER_STATE": 0,
		"BUSY":               1,
		"IDLE":               2,
		"PAUSED":             3,
	}
)

func (x PeerState) Enum() *PeerState {
	p := new(PeerState)
	*p = x
	return p
}

func (x PeerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerState) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_peers_proto_enumTypes[1].Descriptor()
}

func (PeerState) Type() protoreflect.EnumType {
	return &file_proto_peers_proto_enumTypes[1]
}

func (x PeerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerState.Descriptor instead.
func (PeerState) EnumDescriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{1}
}

// Queue is an entry in the Registry describing a distributed queue.
type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User-readable name for the queue
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A user-readable description of the queue's purpose
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	// A web URL (or IPFS CID) for more information on the queue
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// String used for MDNS or DHT rendezvous between peers
	Rendezvous string `protobuf:"bytes,4,opt,name=rendezvous,proto3" json:"rendezvous,omitempty"`
	// List of peers which participate in RAFT leadership and
	// peer assignment (i.e. PeerType.ELECTABLE)
	TrustedPeers []string `protobuf:"bytes,5,rep,name=trustedPeers,proto3" json:"trustedPeers,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{0}
}

func (x *Queue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Queue) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Queue) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Queue) GetRendezvous() string {
	if x != nil {
		return x.Rendezvous
	}
	return ""
}

func (x *Queue) GetTrustedPeers() []string {
	if x != nil {
		return x.TrustedPeers
	}
	return nil
}

// Registry is a listing Queues with some information on
// the list maintainers.
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp of when the registry was written
	Created uint64 `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	// A web URL (or IPFS CID) for more information about the registry
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// List of joinable queues.
	Queues []*Queue `protobuf:"bytes,3,rep,name=queues,proto3" json:"queues,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{1}
}

func (x *Registry) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Registry) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Registry) GetQueues() []*Queue {
	if x != nil {
		return x.Queues
	}
	return nil
}

// PeerStatus is sent by a peer to indicate its role and what it's doing.
type PeerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key ID of the peer
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The topic of queue items on which this peer operates
	// Analgous to a "shard" in other distributed systems, but named so
	// because pubsub is the backing implementation.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// Who this peer thinks is the leader. This is used
	// to inspect queue coherency and peer readiness
	Leader string `protobuf:"bytes,3,opt,name=leader,proto3" json:"leader,omitempty"`
	// The role of the peer
	Type PeerType `protobuf:"varint,4,opt,name=type,proto3,enum=rpc.PeerType" json:"type,omitempty"`
	// What the peer is up to right now
	State PeerState `protobuf:"varint,5,opt,name=state,proto3,enum=rpc.PeerState" json:"state,omitempty"`
}

func (x *PeerStatus) Reset() {
	*x = PeerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerStatus) ProtoMessage() {}

func (x *PeerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerStatus.ProtoReflect.Descriptor instead.
func (*PeerStatus) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{2}
}

func (x *PeerStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeerStatus) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PeerStatus) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *PeerStatus) GetType() PeerType {
	if x != nil {
		return x.Type
	}
	return PeerType_UNKNOWN_PEER_TYPE
}

func (x *PeerStatus) GetState() PeerState {
	if x != nil {
		return x.State
	}
	return PeerState_UNKNOWN_PEER_STATE
}

// Leader is sent when a new leader is asserted on the topic
type Leader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Leader) Reset() {
	*x = Leader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leader) ProtoMessage() {}

func (x *Leader) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leader.ProtoReflect.Descriptor instead.
func (*Leader) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{3}
}

func (x *Leader) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Published after polling to update all clients on state of peers
type PeersSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Estimate may not be fully accurate unless variance indicates such
	PeerEstimate int64 `protobuf:"varint,1,opt,name=peer_estimate,json=peerEstimate,proto3" json:"peer_estimate,omitempty"`
	// Statistical variance of the peer estimate - 0 for exact population count.
	// Note that census taking (choosing k from N) is a binomial distribution
	// https://en.wikipedia.org/wiki/Binomial_distribution
	Variance float64 `protobuf:"fixed64,2,opt,name=variance,proto3" json:"variance,omitempty"`
	// A sample of peers - this is bounded and not necessarily the whole population
	Sample []*PeerStatus `protobuf:"bytes,3,rep,name=sample,proto3" json:"sample,omitempty"`
}

func (x *PeersSummary) Reset() {
	*x = PeersSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersSummary) ProtoMessage() {}

func (x *PeersSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersSummary.ProtoReflect.Descriptor instead.
func (*PeersSummary) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{4}
}

func (x *PeersSummary) GetPeerEstimate() int64 {
	if x != nil {
		return x.PeerEstimate
	}
	return 0
}

func (x *PeersSummary) GetVariance() float64 {
	if x != nil {
		return x.Variance
	}
	return 0
}

func (x *PeersSummary) GetSample() []*PeerStatus {
	if x != nil {
		return x.Sample
	}
	return nil
}

// Probabilistically poll all peers on a topic for census-taking and reporting
type PollPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       string  `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Probability float64 `protobuf:"fixed64,2,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (x *PollPeersRequest) Reset() {
	*x = PollPeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollPeersRequest) ProtoMessage() {}

func (x *PollPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollPeersRequest.ProtoReflect.Descriptor instead.
func (*PollPeersRequest) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{5}
}

func (x *PollPeersRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PollPeersRequest) GetProbability() float64 {
	if x != nil {
		return x.Probability
	}
	return 0
}

// PollPeersResponse is published by peers with probability given in PollPeersRequest
type PollPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *PeerStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PollPeersResponse) Reset() {
	*x = PollPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollPeersResponse) ProtoMessage() {}

func (x *PollPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollPeersResponse.ProtoReflect.Descriptor instead.
func (*PollPeersResponse) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{6}
}

func (x *PollPeersResponse) GetStatus() *PeerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Peer requesting assignment to a particular role. This is typically
// sent to the "assignment" topic
type AssignmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssignmentRequest) Reset() {
	*x = AssignmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignmentRequest) ProtoMessage() {}

func (x *AssignmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignmentRequest.ProtoReflect.Descriptor instead.
func (*AssignmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{7}
}

// Matches https://pkg.go.dev/github.com/libp2p/go-libp2p/core/peer#AddrInfo
// using string encoding for id and addrs.
type AddrInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addrs []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *AddrInfo) Reset() {
	*x = AddrInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddrInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddrInfo) ProtoMessage() {}

func (x *AddrInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddrInfo.ProtoReflect.Descriptor instead.
func (*AddrInfo) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{8}
}

func (x *AddrInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddrInfo) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

// Get address(es) of raft peers on the network
type RaftAddrsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Network details for the requester
	AddrInfo *AddrInfo `protobuf:"bytes,1,opt,name=addr_info,json=addrInfo,proto3" json:"addr_info,omitempty"`
}

func (x *RaftAddrsRequest) Reset() {
	*x = RaftAddrsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftAddrsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftAddrsRequest) ProtoMessage() {}

func (x *RaftAddrsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftAddrsRequest.ProtoReflect.Descriptor instead.
func (*RaftAddrsRequest) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{9}
}

func (x *RaftAddrsRequest) GetAddrInfo() *AddrInfo {
	if x != nil {
		return x.AddrInfo
	}
	return nil
}

// Peers self-report their addresses when RaftAddrsRequest is sent.
type RaftAddrsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of peers present in the consensus group
	AddrInfo *AddrInfo `protobuf:"bytes,1,opt,name=addr_info,json=addrInfo,proto3" json:"addr_info,omitempty"`
}

func (x *RaftAddrsResponse) Reset() {
	*x = RaftAddrsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftAddrsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftAddrsResponse) ProtoMessage() {}

func (x *RaftAddrsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftAddrsResponse.ProtoReflect.Descriptor instead.
func (*RaftAddrsResponse) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{10}
}

func (x *RaftAddrsResponse) GetAddrInfo() *AddrInfo {
	if x != nil {
		return x.AddrInfo
	}
	return nil
}

// A response sent by the leader of the topic to assign a peer to some work
type AssignmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the topic ("shard") to subscribe and listen on
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// ID of the peer being assigned to
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the leader of the topic
	LeaderId string `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	// Role of the peer on this topic
	Type PeerType `protobuf:"varint,4,opt,name=type,proto3,enum=rpc.PeerType" json:"type,omitempty"`
}

func (x *AssignmentResponse) Reset() {
	*x = AssignmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignmentResponse) ProtoMessage() {}

func (x *AssignmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignmentResponse.ProtoReflect.Descriptor instead.
func (*AssignmentResponse) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{11}
}

func (x *AssignmentResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *AssignmentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssignmentResponse) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *AssignmentResponse) GetType() PeerType {
	if x != nil {
		return x.Type
	}
	return PeerType_UNKNOWN_PEER_TYPE
}

// Sent by the newly elected RAFT leader when it wins an election, so others
// know who's in control.
type NewLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key of the new leader
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewLeaderResponse) Reset() {
	*x = NewLeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peers_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLeaderResponse) ProtoMessage() {}

func (x *NewLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peers_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLeaderResponse.ProtoReflect.Descriptor instead.
func (*NewLeaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_peers_proto_rawDescGZIP(), []int{12}
}

func (x *NewLeaderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_peers_proto protoreflect.FileDescriptor

var file_proto_peers_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x22, 0x5a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a,
	0x0a, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x0c,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x65, 0x65, 0x72, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x10, 0x50, 0x6f, 0x6c, 0x6c, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x3c, 0x0a, 0x11, 0x50, 0x6f, 0x6c, 0x6c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x13, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x22, 0x3e, 0x0a, 0x10, 0x52, 0x61, 0x66, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61,
	0x64, 0x64, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x11, 0x52, 0x61, 0x66, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x61, 0x64, 0x64, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7a, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x3e, 0x0a, 0x08, 0x50, 0x65, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x45, 0x4c, 0x45, 0x43, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4c,
	0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x43, 0x0a, 0x09, 0x50, 0x65, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x30, 0x31, 0x35, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_peers_proto_rawDescOnce sync.Once
	file_proto_peers_proto_rawDescData = file_proto_peers_proto_rawDesc
)

func file_proto_peers_proto_rawDescGZIP() []byte {
	file_proto_peers_proto_rawDescOnce.Do(func() {
		file_proto_peers_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_peers_proto_rawDescData)
	})
	return file_proto_peers_proto_rawDescData
}

var file_proto_peers_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_peers_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_peers_proto_goTypes = []interface{}{
	(PeerType)(0),              // 0: rpc.PeerType
	(PeerState)(0),             // 1: rpc.PeerState
	(*Queue)(nil),              // 2: rpc.Queue
	(*Registry)(nil),           // 3: rpc.Registry
	(*PeerStatus)(nil),         // 4: rpc.PeerStatus
	(*Leader)(nil),             // 5: rpc.Leader
	(*PeersSummary)(nil),       // 6: rpc.PeersSummary
	(*PollPeersRequest)(nil),   // 7: rpc.PollPeersRequest
	(*PollPeersResponse)(nil),  // 8: rpc.PollPeersResponse
	(*AssignmentRequest)(nil),  // 9: rpc.AssignmentRequest
	(*AddrInfo)(nil),           // 10: rpc.AddrInfo
	(*RaftAddrsRequest)(nil),   // 11: rpc.RaftAddrsRequest
	(*RaftAddrsResponse)(nil),  // 12: rpc.RaftAddrsResponse
	(*AssignmentResponse)(nil), // 13: rpc.AssignmentResponse
	(*NewLeaderResponse)(nil),  // 14: rpc.NewLeaderResponse
}
var file_proto_peers_proto_depIdxs = []int32{
	2,  // 0: rpc.Registry.queues:type_name -> rpc.Queue
	0,  // 1: rpc.PeerStatus.type:type_name -> rpc.PeerType
	1,  // 2: rpc.PeerStatus.state:type_name -> rpc.PeerState
	4,  // 3: rpc.PeersSummary.sample:type_name -> rpc.PeerStatus
	4,  // 4: rpc.PollPeersResponse.status:type_name -> rpc.PeerStatus
	10, // 5: rpc.RaftAddrsRequest.addr_info:type_name -> rpc.AddrInfo
	10, // 6: rpc.RaftAddrsResponse.addr_info:type_name -> rpc.AddrInfo
	0,  // 7: rpc.AssignmentResponse.type:type_name -> rpc.PeerType
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_peers_proto_init() }
func file_proto_peers_proto_init() {
	if File_proto_peers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_peers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersSummary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollPeersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollPeersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddrInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftAddrsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftAddrsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_peers_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLeaderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_peers_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_peers_proto_goTypes,
		DependencyIndexes: file_proto_peers_proto_depIdxs,
		EnumInfos:         file_proto_peers_proto_enumTypes,
		MessageInfos:      file_proto_peers_proto_msgTypes,
	}.Build()
	File_proto_peers_proto = out.File
	file_proto_peers_proto_rawDesc = nil
	file_proto_peers_proto_goTypes = nil
	file_proto_peers_proto_depIdxs = nil
}
