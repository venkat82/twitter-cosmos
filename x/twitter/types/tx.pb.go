// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: twitter/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type MsgAddTweet struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *MsgAddTweet) Reset()         { *m = MsgAddTweet{} }
func (m *MsgAddTweet) String() string { return proto.CompactTextString(m) }
func (*MsgAddTweet) ProtoMessage()    {}
func (*MsgAddTweet) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{0}
}
func (m *MsgAddTweet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTweet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTweet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTweet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTweet.Merge(m, src)
}
func (m *MsgAddTweet) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTweet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTweet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTweet proto.InternalMessageInfo

func (m *MsgAddTweet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddTweet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MsgAddTweetResponse struct {
}

func (m *MsgAddTweetResponse) Reset()         { *m = MsgAddTweetResponse{} }
func (m *MsgAddTweetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddTweetResponse) ProtoMessage()    {}
func (*MsgAddTweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{1}
}
func (m *MsgAddTweetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddTweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddTweetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddTweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddTweetResponse.Merge(m, src)
}
func (m *MsgAddTweetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddTweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddTweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddTweetResponse proto.InternalMessageInfo

type MsgAddFollower struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	FollowerId string `protobuf:"bytes,2,opt,name=followerId,proto3" json:"followerId,omitempty"`
}

func (m *MsgAddFollower) Reset()         { *m = MsgAddFollower{} }
func (m *MsgAddFollower) String() string { return proto.CompactTextString(m) }
func (*MsgAddFollower) ProtoMessage()    {}
func (*MsgAddFollower) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{2}
}
func (m *MsgAddFollower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddFollower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddFollower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddFollower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddFollower.Merge(m, src)
}
func (m *MsgAddFollower) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddFollower) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddFollower.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddFollower proto.InternalMessageInfo

func (m *MsgAddFollower) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddFollower) GetFollowerId() string {
	if m != nil {
		return m.FollowerId
	}
	return ""
}

type MsgAddFollowerResponse struct {
}

func (m *MsgAddFollowerResponse) Reset()         { *m = MsgAddFollowerResponse{} }
func (m *MsgAddFollowerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddFollowerResponse) ProtoMessage()    {}
func (*MsgAddFollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{3}
}
func (m *MsgAddFollowerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddFollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddFollowerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddFollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddFollowerResponse.Merge(m, src)
}
func (m *MsgAddFollowerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddFollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddFollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddFollowerResponse proto.InternalMessageInfo

type MsgFetchFeed struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgFetchFeed) Reset()         { *m = MsgFetchFeed{} }
func (m *MsgFetchFeed) String() string { return proto.CompactTextString(m) }
func (*MsgFetchFeed) ProtoMessage()    {}
func (*MsgFetchFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{4}
}
func (m *MsgFetchFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFetchFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFetchFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFetchFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFetchFeed.Merge(m, src)
}
func (m *MsgFetchFeed) XXX_Size() int {
	return m.Size()
}
func (m *MsgFetchFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFetchFeed.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFetchFeed proto.InternalMessageInfo

func (m *MsgFetchFeed) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgFetchFeedResponse struct {
	FollowerTweets []string `protobuf:"bytes,1,rep,name=followerTweets,proto3" json:"followerTweets,omitempty"`
}

func (m *MsgFetchFeedResponse) Reset()         { *m = MsgFetchFeedResponse{} }
func (m *MsgFetchFeedResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFetchFeedResponse) ProtoMessage()    {}
func (*MsgFetchFeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d09ccf11b0cc5d3, []int{5}
}
func (m *MsgFetchFeedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFetchFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFetchFeedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFetchFeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFetchFeedResponse.Merge(m, src)
}
func (m *MsgFetchFeedResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFetchFeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFetchFeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFetchFeedResponse proto.InternalMessageInfo

func (m *MsgFetchFeedResponse) GetFollowerTweets() []string {
	if m != nil {
		return m.FollowerTweets
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAddTweet)(nil), "venkat.twitter.twitter.MsgAddTweet")
	proto.RegisterType((*MsgAddTweetResponse)(nil), "venkat.twitter.twitter.MsgAddTweetResponse")
	proto.RegisterType((*MsgAddFollower)(nil), "venkat.twitter.twitter.MsgAddFollower")
	proto.RegisterType((*MsgAddFollowerResponse)(nil), "venkat.twitter.twitter.MsgAddFollowerResponse")
	proto.RegisterType((*MsgFetchFeed)(nil), "venkat.twitter.twitter.MsgFetchFeed")
	proto.RegisterType((*MsgFetchFeedResponse)(nil), "venkat.twitter.twitter.MsgFetchFeedResponse")
}

func init() { proto.RegisterFile("twitter/tx.proto", fileDescriptor_7d09ccf11b0cc5d3) }

var fileDescriptor_7d09ccf11b0cc5d3 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x4b,
	0xcd, 0xcb, 0x4e, 0x2c, 0xd1, 0x83, 0x4a, 0xc0, 0x68, 0x25, 0x47, 0x2e, 0x6e, 0xdf, 0xe2, 0x74,
	0xc7, 0x94, 0x94, 0x90, 0xf2, 0xd4, 0xd4, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4,
	0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x17, 0x2c, 0x93, 0x9f, 0x57,
	0x92, 0x9a, 0x57, 0x22, 0xc1, 0x04, 0x95, 0x81, 0x70, 0x95, 0x44, 0xb9, 0x84, 0x91, 0x8c, 0x08,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0xf2, 0xe2, 0xe2, 0x83, 0x08, 0xbb, 0xe5, 0xe7,
	0xe4, 0xe4, 0x97, 0xa7, 0x16, 0xe1, 0x31, 0x5c, 0x8e, 0x8b, 0x2b, 0x0d, 0xaa, 0xca, 0x33, 0x05,
	0x6a, 0x3e, 0x92, 0x88, 0x92, 0x04, 0x97, 0x18, 0xaa, 0x59, 0x70, 0x5b, 0x34, 0xb8, 0x78, 0x7c,
	0x8b, 0xd3, 0xdd, 0x52, 0x4b, 0x92, 0x33, 0xdc, 0x52, 0x53, 0x53, 0x70, 0xdb, 0xa1, 0x64, 0xc7,
	0x25, 0x82, 0xac, 0x12, 0x66, 0x82, 0x90, 0x1a, 0x17, 0x1f, 0xcc, 0x26, 0xb0, 0x07, 0x8a, 0x25,
	0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0xd0, 0x44, 0x8d, 0x56, 0x33, 0x71, 0x31, 0xfb, 0x16, 0xa7,
	0x0b, 0xc5, 0x70, 0x71, 0xc0, 0x83, 0x4b, 0x59, 0x0f, 0x7b, 0xb0, 0xea, 0x21, 0x05, 0x88, 0x94,
	0x36, 0x11, 0x8a, 0xe0, 0xae, 0x49, 0xe5, 0xe2, 0x46, 0x0e, 0x32, 0x35, 0xfc, 0x7a, 0x61, 0xea,
	0xa4, 0xf4, 0x88, 0x53, 0x07, 0xb7, 0x26, 0x9e, 0x8b, 0x13, 0x11, 0x66, 0x2a, 0x78, 0x34, 0xc3,
	0x55, 0x49, 0xe9, 0x10, 0xa3, 0x0a, 0x66, 0x81, 0x93, 0xfd, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x43, 0x4c, 0xd3, 0x87,
	0xa5, 0xd4, 0x0a, 0x38, 0xab, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0x6e, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x8d, 0x4c, 0x9a, 0xcb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddTweet(ctx context.Context, in *MsgAddTweet, opts ...grpc.CallOption) (*MsgAddTweetResponse, error)
	AddFollower(ctx context.Context, in *MsgAddFollower, opts ...grpc.CallOption) (*MsgAddFollowerResponse, error)
	FetchFeed(ctx context.Context, in *MsgFetchFeed, opts ...grpc.CallOption) (*MsgFetchFeedResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddTweet(ctx context.Context, in *MsgAddTweet, opts ...grpc.CallOption) (*MsgAddTweetResponse, error) {
	out := new(MsgAddTweetResponse)
	err := c.cc.Invoke(ctx, "/venkat.twitter.twitter.Msg/AddTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddFollower(ctx context.Context, in *MsgAddFollower, opts ...grpc.CallOption) (*MsgAddFollowerResponse, error) {
	out := new(MsgAddFollowerResponse)
	err := c.cc.Invoke(ctx, "/venkat.twitter.twitter.Msg/AddFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FetchFeed(ctx context.Context, in *MsgFetchFeed, opts ...grpc.CallOption) (*MsgFetchFeedResponse, error) {
	out := new(MsgFetchFeedResponse)
	err := c.cc.Invoke(ctx, "/venkat.twitter.twitter.Msg/FetchFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddTweet(context.Context, *MsgAddTweet) (*MsgAddTweetResponse, error)
	AddFollower(context.Context, *MsgAddFollower) (*MsgAddFollowerResponse, error)
	FetchFeed(context.Context, *MsgFetchFeed) (*MsgFetchFeedResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddTweet(ctx context.Context, req *MsgAddTweet) (*MsgAddTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTweet not implemented")
}
func (*UnimplementedMsgServer) AddFollower(ctx context.Context, req *MsgAddFollower) (*MsgAddFollowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFollower not implemented")
}
func (*UnimplementedMsgServer) FetchFeed(ctx context.Context, req *MsgFetchFeed) (*MsgFetchFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFeed not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddTweet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venkat.twitter.twitter.Msg/AddTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddTweet(ctx, req.(*MsgAddTweet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddFollower)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venkat.twitter.twitter.Msg/AddFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddFollower(ctx, req.(*MsgAddFollower))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FetchFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFetchFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FetchFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/venkat.twitter.twitter.Msg/FetchFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FetchFeed(ctx, req.(*MsgFetchFeed))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "venkat.twitter.twitter.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTweet",
			Handler:    _Msg_AddTweet_Handler,
		},
		{
			MethodName: "AddFollower",
			Handler:    _Msg_AddFollower_Handler,
		},
		{
			MethodName: "FetchFeed",
			Handler:    _Msg_FetchFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "twitter/tx.proto",
}

func (m *MsgAddTweet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTweet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTweet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddTweetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddTweetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddTweetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddFollower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddFollower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddFollower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FollowerId) > 0 {
		i -= len(m.FollowerId)
		copy(dAtA[i:], m.FollowerId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FollowerId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddFollowerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddFollowerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddFollowerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgFetchFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFetchFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFetchFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFetchFeedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFetchFeedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFetchFeedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FollowerTweets) > 0 {
		for iNdEx := len(m.FollowerTweets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FollowerTweets[iNdEx])
			copy(dAtA[i:], m.FollowerTweets[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.FollowerTweets[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddTweet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddTweetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddFollower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.FollowerId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddFollowerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgFetchFeed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFetchFeedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FollowerTweets) > 0 {
		for _, s := range m.FollowerTweets {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddTweet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTweet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTweet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddTweetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddTweetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddTweetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddFollower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddFollower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddFollower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FollowerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FollowerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAddFollowerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddFollowerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddFollowerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgFetchFeed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFetchFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFetchFeed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgFetchFeedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFetchFeedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFetchFeedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FollowerTweets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FollowerTweets = append(m.FollowerTweets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
