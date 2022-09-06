// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/peers.proto

package proto // import "github.com/smartin015/peerprint/pubsub/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PeerType int32

const (
	PeerType_UNKNOWN_PEER_TYPE PeerType = 0
	PeerType_LEADER            PeerType = 1
	PeerType_CANDIDATE         PeerType = 2
	PeerType_FOLLOWER          PeerType = 3
	PeerType_LISTENER          PeerType = 4
)

var PeerType_name = map[int32]string{
	0: "UNKNOWN_PEER_TYPE",
	1: "LEADER",
	2: "CANDIDATE",
	3: "FOLLOWER",
	4: "LISTENER",
}
var PeerType_value = map[string]int32{
	"UNKNOWN_PEER_TYPE": 0,
	"LEADER":            1,
	"CANDIDATE":         2,
	"FOLLOWER":          3,
	"LISTENER":          4,
}

func (x PeerType) String() string {
	return proto.EnumName(PeerType_name, int32(x))
}
func (PeerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{0}
}

type PeerState int32

const (
	PeerState_UNKNOWN_PEER_STATE PeerState = 0
	PeerState_BUSY               PeerState = 1
	PeerState_IDLE               PeerState = 2
	PeerState_PAUSED             PeerState = 3
)

var PeerState_name = map[int32]string{
	0: "UNKNOWN_PEER_STATE",
	1: "BUSY",
	2: "IDLE",
	3: "PAUSED",
}
var PeerState_value = map[string]int32{
	"UNKNOWN_PEER_STATE": 0,
	"BUSY":               1,
	"IDLE":               2,
	"PAUSED":             3,
}

func (x PeerState) String() string {
	return proto.EnumName(PeerState_name, int32(x))
}
func (PeerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{1}
}

type PeerStatus struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string    `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Type                 PeerType  `protobuf:"varint,3,opt,name=type,proto3,enum=rpc.PeerType" json:"type,omitempty"`
	State                PeerState `protobuf:"varint,4,opt,name=state,proto3,enum=rpc.PeerState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PeerStatus) Reset()         { *m = PeerStatus{} }
func (m *PeerStatus) String() string { return proto.CompactTextString(m) }
func (*PeerStatus) ProtoMessage()    {}
func (*PeerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{0}
}
func (m *PeerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerStatus.Unmarshal(m, b)
}
func (m *PeerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerStatus.Marshal(b, m, deterministic)
}
func (dst *PeerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerStatus.Merge(dst, src)
}
func (m *PeerStatus) XXX_Size() int {
	return xxx_messageInfo_PeerStatus.Size(m)
}
func (m *PeerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PeerStatus proto.InternalMessageInfo

func (m *PeerStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerStatus) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PeerStatus) GetType() PeerType {
	if m != nil {
		return m.Type
	}
	return PeerType_UNKNOWN_PEER_TYPE
}

func (m *PeerStatus) GetState() PeerState {
	if m != nil {
		return m.State
	}
	return PeerState_UNKNOWN_PEER_STATE
}

// Probabilistically poll all peers on a topic for census-taking and reporting
type PollPeersRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Probability          float64  `protobuf:"fixed64,2,opt,name=probability,proto3" json:"probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollPeersRequest) Reset()         { *m = PollPeersRequest{} }
func (m *PollPeersRequest) String() string { return proto.CompactTextString(m) }
func (*PollPeersRequest) ProtoMessage()    {}
func (*PollPeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{1}
}
func (m *PollPeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollPeersRequest.Unmarshal(m, b)
}
func (m *PollPeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollPeersRequest.Marshal(b, m, deterministic)
}
func (dst *PollPeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollPeersRequest.Merge(dst, src)
}
func (m *PollPeersRequest) XXX_Size() int {
	return xxx_messageInfo_PollPeersRequest.Size(m)
}
func (m *PollPeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PollPeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PollPeersRequest proto.InternalMessageInfo

func (m *PollPeersRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PollPeersRequest) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

// PollPeersResponse is published by peers with probability given in PollPeersRequest
type PollPeersResponse struct {
	Status               *PeerStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PollPeersResponse) Reset()         { *m = PollPeersResponse{} }
func (m *PollPeersResponse) String() string { return proto.CompactTextString(m) }
func (*PollPeersResponse) ProtoMessage()    {}
func (*PollPeersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{2}
}
func (m *PollPeersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollPeersResponse.Unmarshal(m, b)
}
func (m *PollPeersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollPeersResponse.Marshal(b, m, deterministic)
}
func (dst *PollPeersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollPeersResponse.Merge(dst, src)
}
func (m *PollPeersResponse) XXX_Size() int {
	return xxx_messageInfo_PollPeersResponse.Size(m)
}
func (m *PollPeersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PollPeersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PollPeersResponse proto.InternalMessageInfo

func (m *PollPeersResponse) GetStatus() *PeerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Peer requesting assignment to a particular role. This is typically
// sent to the "assignment" topic
type AssignmentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignmentRequest) Reset()         { *m = AssignmentRequest{} }
func (m *AssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*AssignmentRequest) ProtoMessage()    {}
func (*AssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{3}
}
func (m *AssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignmentRequest.Unmarshal(m, b)
}
func (m *AssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignmentRequest.Marshal(b, m, deterministic)
}
func (dst *AssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignmentRequest.Merge(dst, src)
}
func (m *AssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_AssignmentRequest.Size(m)
}
func (m *AssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignmentRequest proto.InternalMessageInfo

// A response sent by the leader of the topic to assign a peer to some work
type AssignmentResponse struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	LeaderId             string   `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Type                 PeerType `protobuf:"varint,4,opt,name=type,proto3,enum=rpc.PeerType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignmentResponse) Reset()         { *m = AssignmentResponse{} }
func (m *AssignmentResponse) String() string { return proto.CompactTextString(m) }
func (*AssignmentResponse) ProtoMessage()    {}
func (*AssignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_peers_985b03ba3cb95b4e, []int{4}
}
func (m *AssignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignmentResponse.Unmarshal(m, b)
}
func (m *AssignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignmentResponse.Marshal(b, m, deterministic)
}
func (dst *AssignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignmentResponse.Merge(dst, src)
}
func (m *AssignmentResponse) XXX_Size() int {
	return xxx_messageInfo_AssignmentResponse.Size(m)
}
func (m *AssignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssignmentResponse proto.InternalMessageInfo

func (m *AssignmentResponse) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *AssignmentResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssignmentResponse) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *AssignmentResponse) GetType() PeerType {
	if m != nil {
		return m.Type
	}
	return PeerType_UNKNOWN_PEER_TYPE
}

func init() {
	proto.RegisterType((*PeerStatus)(nil), "rpc.PeerStatus")
	proto.RegisterType((*PollPeersRequest)(nil), "rpc.PollPeersRequest")
	proto.RegisterType((*PollPeersResponse)(nil), "rpc.PollPeersResponse")
	proto.RegisterType((*AssignmentRequest)(nil), "rpc.AssignmentRequest")
	proto.RegisterType((*AssignmentResponse)(nil), "rpc.AssignmentResponse")
	proto.RegisterEnum("rpc.PeerType", PeerType_name, PeerType_value)
	proto.RegisterEnum("rpc.PeerState", PeerState_name, PeerState_value)
}

func init() { proto.RegisterFile("proto/peers.proto", fileDescriptor_peers_985b03ba3cb95b4e) }

var fileDescriptor_peers_985b03ba3cb95b4e = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x97, 0x34, 0xab, 0x9a, 0x3b, 0x56, 0x5c, 0xf3, 0x47, 0x95, 0x78, 0x29, 0x11, 0x12,
	0xd3, 0x84, 0x52, 0x18, 0xe2, 0x8d, 0x97, 0x6c, 0x31, 0x52, 0x21, 0x4a, 0x23, 0x27, 0xd5, 0x18,
	0x2f, 0x55, 0xd2, 0x58, 0xc3, 0x52, 0x9b, 0x18, 0xdb, 0x79, 0x28, 0x7c, 0x79, 0x14, 0xb7, 0x84,
	0x4c, 0x82, 0xa7, 0xd8, 0xe7, 0x5e, 0xfd, 0xee, 0xc9, 0xf1, 0x85, 0x89, 0x90, 0xb5, 0xae, 0xe7,
	0x82, 0x31, 0xa9, 0x7c, 0x73, 0xc6, 0x03, 0x29, 0x36, 0xde, 0x2f, 0x80, 0x84, 0x31, 0x99, 0xea,
	0x5c, 0x37, 0x0a, 0x8f, 0xc1, 0xe6, 0xe5, 0xd4, 0x9a, 0x59, 0x17, 0x2e, 0xb5, 0x79, 0x89, 0x9f,
	0xc2, 0xa9, 0xae, 0x05, 0xdf, 0x4c, 0x6d, 0x23, 0x1d, 0x2e, 0xf8, 0x25, 0x38, 0x7a, 0x2f, 0xd8,
	0x74, 0x30, 0xb3, 0x2e, 0xc6, 0x57, 0xe7, 0xbe, 0x14, 0x1b, 0xbf, 0x85, 0x64, 0x7b, 0xc1, 0xa8,
	0x29, 0xe1, 0x57, 0x70, 0xaa, 0x74, 0xae, 0xd9, 0xd4, 0x31, 0x3d, 0xe3, 0xae, 0xa7, 0x1d, 0xc4,
	0xe8, 0xa1, 0xe8, 0x7d, 0x06, 0x94, 0xd4, 0xdb, 0x6d, 0xab, 0x2b, 0xca, 0x7e, 0x34, 0x4c, 0xe9,
	0xbf, 0x23, 0xad, 0xfe, 0xc8, 0x19, 0x9c, 0x09, 0x59, 0x17, 0x79, 0xc1, 0xb7, 0x5c, 0xef, 0x8d,
	0x1d, 0x8b, 0xf6, 0x25, 0xef, 0x23, 0x4c, 0x7a, 0x2c, 0x25, 0xea, 0x4a, 0x31, 0xfc, 0x1a, 0x86,
	0xca, 0xfc, 0x99, 0xa1, 0x9d, 0x5d, 0x3d, 0x7e, 0xe0, 0xa3, 0x51, 0xf4, 0x58, 0xf6, 0x9e, 0xc0,
	0x24, 0x50, 0x8a, 0xdf, 0x57, 0x3b, 0x56, 0xe9, 0xa3, 0x15, 0xef, 0x27, 0xe0, 0xbe, 0x78, 0x64,
	0xfe, 0xdb, 0xe0, 0x21, 0x39, 0xbb, 0x4b, 0xee, 0x05, 0xb8, 0x5b, 0x96, 0x97, 0x4c, 0xae, 0x79,
	0x69, 0x82, 0x72, 0xe9, 0xe8, 0x20, 0x2c, 0xca, 0x2e, 0x40, 0xe7, 0xbf, 0x01, 0x5e, 0x7e, 0x85,
	0xd1, 0x1f, 0x05, 0x3f, 0x83, 0xc9, 0x2a, 0xfe, 0x12, 0x2f, 0x6f, 0xe3, 0x75, 0x42, 0x08, 0x5d,
	0x67, 0x77, 0x09, 0x41, 0x27, 0x18, 0x60, 0x18, 0x91, 0x20, 0x24, 0x14, 0x59, 0xf8, 0x1c, 0xdc,
	0x9b, 0x20, 0x0e, 0x17, 0x61, 0x90, 0x11, 0x64, 0xe3, 0x47, 0x30, 0xfa, 0xb4, 0x8c, 0xa2, 0xe5,
	0x2d, 0xa1, 0x68, 0xd0, 0xde, 0xa2, 0x45, 0x9a, 0x91, 0x98, 0x50, 0xe4, 0x5c, 0xde, 0x80, 0xdb,
	0x3d, 0x04, 0x7e, 0x0e, 0xf8, 0x01, 0x3a, 0xcd, 0x5a, 0xc0, 0x09, 0x1e, 0x81, 0x73, 0xbd, 0x4a,
	0xef, 0x90, 0xd5, 0x9e, 0x16, 0x61, 0xd4, 0x42, 0x01, 0x86, 0x49, 0xb0, 0x4a, 0x49, 0x88, 0x06,
	0xd7, 0xfe, 0xb7, 0x37, 0xf7, 0x5c, 0x7f, 0x6f, 0x0a, 0x7f, 0x53, 0xef, 0xe6, 0x6a, 0x97, 0x4b,
	0xcd, 0xab, 0xb7, 0xef, 0x3e, 0x98, 0x05, 0x13, 0x92, 0x57, 0x7a, 0x2e, 0x9a, 0x42, 0x35, 0xc5,
	0xdc, 0xec, 0x5a, 0x31, 0x34, 0x9f, 0xf7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x18, 0x47,
	0x73, 0x87, 0x02, 0x00, 0x00,
}
