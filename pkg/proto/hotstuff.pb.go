// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hotstuff.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/relab/gorums"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	GENERIC    Type = 0
	PREPARE    Type = 1
	PRE_COMMIT Type = 2
	COMMIT     Type = 3
	DECIDE     Type = 4
	NEW_VIEW   Type = 5
)

var Type_name = map[int32]string{
	0: "GENERIC",
	1: "PREPARE",
	2: "PRE_COMMIT",
	3: "COMMIT",
	4: "DECIDE",
	5: "NEW_VIEW",
}

var Type_value = map[string]int32{
	"GENERIC":    0,
	"PREPARE":    1,
	"PRE_COMMIT": 2,
	"COMMIT":     3,
	"DECIDE":     4,
	"NEW_VIEW":   5,
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8517aa0e19c54851, []int{0}
}

type Msg struct {
	Type                 Type        `protobuf:"varint,1,opt,name=Type,proto3,enum=proto.Type" json:"Type,omitempty"`
	ViewNumber           int32       `protobuf:"varint,2,opt,name=ViewNumber,proto3" json:"ViewNumber,omitempty"`
	Node                 *HSNode     `protobuf:"bytes,3,opt,name=Node,proto3" json:"Node,omitempty"`
	Justify              *QuorumCert `protobuf:"bytes,4,opt,name=Justify,proto3" json:"Justify,omitempty"`
	PartialSig           string      `protobuf:"bytes,5,opt,name=PartialSig,proto3" json:"PartialSig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Msg) Reset()      { *m = Msg{} }
func (*Msg) ProtoMessage() {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_8517aa0e19c54851, []int{0}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetType() Type {
	if m != nil {
		return m.Type
	}
	return GENERIC
}

func (m *Msg) GetViewNumber() int32 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *Msg) GetNode() *HSNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Msg) GetJustify() *QuorumCert {
	if m != nil {
		return m.Justify
	}
	return nil
}

func (m *Msg) GetPartialSig() string {
	if m != nil {
		return m.PartialSig
	}
	return ""
}

type QuorumCert struct {
	Type                 Type     `protobuf:"varint,1,opt,name=Type,proto3,enum=proto.Type" json:"Type,omitempty"`
	ViewNumber           int32    `protobuf:"varint,2,opt,name=ViewNumber,proto3" json:"ViewNumber,omitempty"`
	Node                 *HSNode  `protobuf:"bytes,3,opt,name=Node,proto3" json:"Node,omitempty"`
	Sig                  string   `protobuf:"bytes,4,opt,name=Sig,proto3" json:"Sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuorumCert) Reset()      { *m = QuorumCert{} }
func (*QuorumCert) ProtoMessage() {}
func (*QuorumCert) Descriptor() ([]byte, []int) {
	return fileDescriptor_8517aa0e19c54851, []int{1}
}
func (m *QuorumCert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuorumCert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuorumCert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuorumCert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuorumCert.Merge(m, src)
}
func (m *QuorumCert) XXX_Size() int {
	return m.Size()
}
func (m *QuorumCert) XXX_DiscardUnknown() {
	xxx_messageInfo_QuorumCert.DiscardUnknown(m)
}

var xxx_messageInfo_QuorumCert proto.InternalMessageInfo

func (m *QuorumCert) GetType() Type {
	if m != nil {
		return m.Type
	}
	return GENERIC
}

func (m *QuorumCert) GetViewNumber() int32 {
	if m != nil {
		return m.ViewNumber
	}
	return 0
}

func (m *QuorumCert) GetNode() *HSNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *QuorumCert) GetSig() string {
	if m != nil {
		return m.Sig
	}
	return ""
}

type HSNode struct {
	ParentHash           string   `protobuf:"bytes,1,opt,name=ParentHash,proto3" json:"ParentHash,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSNode) Reset()      { *m = HSNode{} }
func (*HSNode) ProtoMessage() {}
func (*HSNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8517aa0e19c54851, []int{2}
}
func (m *HSNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HSNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HSNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HSNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSNode.Merge(m, src)
}
func (m *HSNode) XXX_Size() int {
	return m.Size()
}
func (m *HSNode) XXX_DiscardUnknown() {
	xxx_messageInfo_HSNode.DiscardUnknown(m)
}

var xxx_messageInfo_HSNode proto.InternalMessageInfo

func (m *HSNode) GetParentHash() string {
	if m != nil {
		return m.ParentHash
	}
	return ""
}

func (m *HSNode) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.Type", Type_name, Type_value)
	proto.RegisterType((*Msg)(nil), "proto.Msg")
	proto.RegisterType((*QuorumCert)(nil), "proto.QuorumCert")
	proto.RegisterType((*HSNode)(nil), "proto.HSNode")
}

func init() { proto.RegisterFile("hotstuff.proto", fileDescriptor_8517aa0e19c54851) }

var fileDescriptor_8517aa0e19c54851 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x8d, 0x93, 0x90, 0x09, 0x04, 0xb3, 0x87, 0xca, 0x0a, 0xd2, 0x12, 0xa2, 0x1e,
	0x22, 0x10, 0x8e, 0x14, 0x8e, 0x5c, 0x20, 0xee, 0x8a, 0x04, 0x61, 0x13, 0xb6, 0x55, 0x7a, 0xac,
	0xec, 0x64, 0xe3, 0x58, 0x8a, 0xbb, 0x91, 0xbd, 0x56, 0x95, 0x5b, 0xc5, 0x13, 0xf0, 0x18, 0xbc,
	0x40, 0x79, 0x00, 0x4e, 0x1c, 0x7b, 0xe4, 0x48, 0xcd, 0x85, 0x23, 0x12, 0x2f, 0x80, 0x76, 0xed,
	0x88, 0xf0, 0x02, 0x9c, 0x76, 0xe6, 0x9f, 0x7f, 0x67, 0xbf, 0x1d, 0x0d, 0xb4, 0x57, 0x42, 0xa6,
	0x32, 0x5b, 0x2e, 0xed, 0x4d, 0x22, 0xa4, 0xc0, 0x35, 0x7d, 0x74, 0x8e, 0xc2, 0x48, 0xae, 0xb2,
	0xc0, 0x9e, 0x8b, 0x78, 0x90, 0xf0, 0xb5, 0x1f, 0x0c, 0x42, 0x91, 0x64, 0x71, 0x5a, 0x1e, 0x85,
	0xb9, 0xf3, 0x30, 0x14, 0x22, 0x5c, 0xf3, 0x81, 0xce, 0x82, 0x6c, 0x39, 0xe0, 0xf1, 0x46, 0x6e,
	0x8b, 0x62, 0xef, 0x33, 0x82, 0xaa, 0x9b, 0x86, 0xf8, 0x11, 0x18, 0xa7, 0xdb, 0x0d, 0xb7, 0x50,
	0x17, 0xf5, 0xdb, 0xc3, 0x56, 0x51, 0xb5, 0x95, 0xc4, 0x74, 0x01, 0x13, 0x80, 0x59, 0xc4, 0x2f,
	0xbd, 0x2c, 0x0e, 0x78, 0x62, 0x1d, 0x74, 0x51, 0xbf, 0xc6, 0xf6, 0x14, 0xfc, 0x18, 0x0c, 0x4f,
	0x2c, 0xb8, 0x55, 0xed, 0xa2, 0x7e, 0x6b, 0x78, 0xaf, 0x6c, 0x30, 0x3e, 0x51, 0x22, 0xd3, 0x25,
	0xfc, 0x14, 0x1a, 0x6f, 0xb2, 0x54, 0x46, 0xcb, 0xad, 0x65, 0x68, 0xd7, 0x83, 0xd2, 0xf5, 0x3e,
	0x53, 0xbc, 0x0e, 0x4f, 0x24, 0xdb, 0x39, 0xd4, 0x7b, 0x53, 0x3f, 0x91, 0x91, 0xbf, 0x3e, 0x89,
	0x42, 0xab, 0xd6, 0x45, 0xfd, 0x26, 0xdb, 0x53, 0x7a, 0x1f, 0x10, 0xc0, 0xdf, 0x7b, 0xff, 0x85,
	0xdf, 0x84, 0xaa, 0x62, 0x31, 0x34, 0x8b, 0x0a, 0x7b, 0x23, 0xa8, 0x17, 0x8e, 0x12, 0x97, 0x5f,
	0xc8, 0xb1, 0x9f, 0xae, 0x34, 0x45, 0x81, 0x5b, 0x2a, 0xd8, 0x82, 0x86, 0x23, 0xe2, 0xd8, 0xbf,
	0x58, 0xe8, 0xb7, 0x9b, 0x6c, 0x97, 0x3e, 0x99, 0x15, 0xe4, 0xb8, 0x05, 0x8d, 0xd7, 0xd4, 0xa3,
	0x6c, 0xe2, 0x98, 0x15, 0x95, 0x4c, 0x19, 0x9d, 0xbe, 0x62, 0xd4, 0x44, 0xb8, 0x0d, 0x30, 0x65,
	0xf4, 0xdc, 0x79, 0xe7, 0xba, 0x93, 0x53, 0xf3, 0x00, 0x03, 0xd4, 0xcb, 0xb8, 0xaa, 0xe2, 0x63,
	0xea, 0x4c, 0x8e, 0xa9, 0x69, 0xe0, 0xbb, 0x70, 0xc7, 0xa3, 0x67, 0xe7, 0xb3, 0x09, 0x3d, 0x33,
	0x6b, 0x43, 0x17, 0xda, 0xe3, 0x72, 0x6b, 0xde, 0x72, 0x7f, 0xc1, 0x13, 0xfc, 0x02, 0x9a, 0xa3,
	0x44, 0xf8, 0x8b, 0xb9, 0x9f, 0x4a, 0x0c, 0xe5, 0x0f, 0xdd, 0x34, 0xec, 0xec, 0xc5, 0xbd, 0xc3,
	0xab, 0x6b, 0x0b, 0x7d, 0xba, 0xb6, 0xd0, 0x97, 0xdf, 0xd6, 0xde, 0x80, 0x87, 0x2f, 0xe1, 0xfe,
	0xae, 0x1d, 0xe3, 0x9b, 0x75, 0x34, 0xf7, 0xf1, 0x33, 0x68, 0x78, 0xfc, 0x52, 0xcd, 0xf0, 0x9f,
	0x6e, 0x87, 0x76, 0xb1, 0x70, 0xf6, 0x6e, 0xe1, 0x6c, 0xaa, 0x16, 0x6e, 0x74, 0x74, 0x73, 0x4b,
	0x2a, 0xdf, 0x6e, 0x49, 0xe5, 0x2a, 0x27, 0xe8, 0x6b, 0x4e, 0xd0, 0x4d, 0x4e, 0xd0, 0xf7, 0x9c,
	0xa0, 0x9f, 0x39, 0xa9, 0xfc, 0xca, 0x09, 0xfa, 0xf8, 0x83, 0x54, 0x82, 0xba, 0xbe, 0xf5, 0xfc,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x10, 0xb4, 0x44, 0xf3, 0x02, 0x00, 0x00,
}

func (x Type) String() string {
	s, ok := Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotstuffLeaderClient is the client API for HotstuffLeader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotstuffLeaderClient interface {
	Broadcast(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error)
}

type hotstuffLeaderClient struct {
	cc *grpc.ClientConn
}

func NewHotstuffLeaderClient(cc *grpc.ClientConn) HotstuffLeaderClient {
	return &hotstuffLeaderClient{cc}
}

func (c *hotstuffLeaderClient) Broadcast(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := c.cc.Invoke(ctx, "/proto.HotstuffLeader/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotstuffLeaderServer is the server API for HotstuffLeader service.
type HotstuffLeaderServer interface {
	Broadcast(context.Context, *Msg) (*Msg, error)
}

// UnimplementedHotstuffLeaderServer can be embedded to have forward compatible implementations.
type UnimplementedHotstuffLeaderServer struct {
}

func (*UnimplementedHotstuffLeaderServer) Broadcast(ctx context.Context, req *Msg) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterHotstuffLeaderServer(s *grpc.Server, srv HotstuffLeaderServer) {
	s.RegisterService(&_HotstuffLeader_serviceDesc, srv)
}

func _HotstuffLeader_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotstuffLeaderServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HotstuffLeader/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotstuffLeaderServer).Broadcast(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotstuffLeader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HotstuffLeader",
	HandlerType: (*HotstuffLeaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _HotstuffLeader_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotstuff.proto",
}

// HotstuffReplicaClient is the client API for HotstuffReplica service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotstuffReplicaClient interface {
	NewView(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type hotstuffReplicaClient struct {
	cc *grpc.ClientConn
}

func NewHotstuffReplicaClient(cc *grpc.ClientConn) HotstuffReplicaClient {
	return &hotstuffReplicaClient{cc}
}

func (c *hotstuffReplicaClient) NewView(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.HotstuffReplica/NewView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotstuffReplicaServer is the server API for HotstuffReplica service.
type HotstuffReplicaServer interface {
	NewView(context.Context, *Msg) (*empty.Empty, error)
}

// UnimplementedHotstuffReplicaServer can be embedded to have forward compatible implementations.
type UnimplementedHotstuffReplicaServer struct {
}

func (*UnimplementedHotstuffReplicaServer) NewView(ctx context.Context, req *Msg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewView not implemented")
}

func RegisterHotstuffReplicaServer(s *grpc.Server, srv HotstuffReplicaServer) {
	s.RegisterService(&_HotstuffReplica_serviceDesc, srv)
}

func _HotstuffReplica_NewView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotstuffReplicaServer).NewView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HotstuffReplica/NewView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotstuffReplicaServer).NewView(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotstuffReplica_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HotstuffReplica",
	HandlerType: (*HotstuffReplicaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewView",
			Handler:    _HotstuffReplica_NewView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotstuff.proto",
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PartialSig) > 0 {
		i -= len(m.PartialSig)
		copy(dAtA[i:], m.PartialSig)
		i = encodeVarintHotstuff(dAtA, i, uint64(len(m.PartialSig)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Justify != nil {
		{
			size, err := m.Justify.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHotstuff(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHotstuff(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ViewNumber != 0 {
		i = encodeVarintHotstuff(dAtA, i, uint64(m.ViewNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintHotstuff(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuorumCert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuorumCert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuorumCert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintHotstuff(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x22
	}
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHotstuff(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ViewNumber != 0 {
		i = encodeVarintHotstuff(dAtA, i, uint64(m.ViewNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintHotstuff(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HSNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HSNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HSNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintHotstuff(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarintHotstuff(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHotstuff(dAtA []byte, offset int, v uint64) int {
	offset -= sovHotstuff(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovHotstuff(uint64(m.Type))
	}
	if m.ViewNumber != 0 {
		n += 1 + sovHotstuff(uint64(m.ViewNumber))
	}
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovHotstuff(uint64(l))
	}
	if m.Justify != nil {
		l = m.Justify.Size()
		n += 1 + l + sovHotstuff(uint64(l))
	}
	l = len(m.PartialSig)
	if l > 0 {
		n += 1 + l + sovHotstuff(uint64(l))
	}
	return n
}

func (m *QuorumCert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovHotstuff(uint64(m.Type))
	}
	if m.ViewNumber != 0 {
		n += 1 + sovHotstuff(uint64(m.ViewNumber))
	}
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovHotstuff(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovHotstuff(uint64(l))
	}
	return n
}

func (m *HSNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovHotstuff(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHotstuff(uint64(l))
	}
	return n
}

func sovHotstuff(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHotstuff(x uint64) (n int) {
	return sovHotstuff(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Msg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Msg{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`ViewNumber:` + fmt.Sprintf("%v", this.ViewNumber) + `,`,
		`Node:` + strings.Replace(this.Node.String(), "HSNode", "HSNode", 1) + `,`,
		`Justify:` + strings.Replace(this.Justify.String(), "QuorumCert", "QuorumCert", 1) + `,`,
		`PartialSig:` + fmt.Sprintf("%v", this.PartialSig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuorumCert) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuorumCert{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`ViewNumber:` + fmt.Sprintf("%v", this.ViewNumber) + `,`,
		`Node:` + strings.Replace(this.Node.String(), "HSNode", "HSNode", 1) + `,`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HSNode) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HSNode{`,
		`ParentHash:` + fmt.Sprintf("%v", this.ParentHash) + `,`,
		`Command:` + fmt.Sprintf("%v", this.Command) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHotstuff(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHotstuff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViewNumber", wireType)
			}
			m.ViewNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ViewNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &HSNode{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Justify", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Justify == nil {
				m.Justify = &QuorumCert{}
			}
			if err := m.Justify.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartialSig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartialSig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHotstuff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuorumCert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHotstuff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuorumCert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuorumCert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViewNumber", wireType)
			}
			m.ViewNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ViewNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &HSNode{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHotstuff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HSNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHotstuff
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HSNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HSNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHotstuff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHotstuff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHotstuff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHotstuff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHotstuff(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHotstuff
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHotstuff
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHotstuff
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHotstuff
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHotstuff
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHotstuff        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHotstuff          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHotstuff = fmt.Errorf("proto: unexpected end of group")
)