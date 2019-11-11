// Code generated by protoc-gen-go. DO NOT EDIT.
// source: charbalancer.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BalanceStringRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceStringRequest) Reset()         { *m = BalanceStringRequest{} }
func (m *BalanceStringRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceStringRequest) ProtoMessage()    {}
func (*BalanceStringRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7315142929410128, []int{0}
}

func (m *BalanceStringRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceStringRequest.Unmarshal(m, b)
}
func (m *BalanceStringRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceStringRequest.Marshal(b, m, deterministic)
}
func (m *BalanceStringRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceStringRequest.Merge(m, src)
}
func (m *BalanceStringRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceStringRequest.Size(m)
}
func (m *BalanceStringRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceStringRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceStringRequest proto.InternalMessageInfo

func (m *BalanceStringRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type BalanceStringResponse struct {
	Balanced             bool     `protobuf:"varint,1,opt,name=balanced,proto3" json:"balanced,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceStringResponse) Reset()         { *m = BalanceStringResponse{} }
func (m *BalanceStringResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceStringResponse) ProtoMessage()    {}
func (*BalanceStringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7315142929410128, []int{1}
}

func (m *BalanceStringResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceStringResponse.Unmarshal(m, b)
}
func (m *BalanceStringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceStringResponse.Marshal(b, m, deterministic)
}
func (m *BalanceStringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceStringResponse.Merge(m, src)
}
func (m *BalanceStringResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceStringResponse.Size(m)
}
func (m *BalanceStringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceStringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceStringResponse proto.InternalMessageInfo

func (m *BalanceStringResponse) GetBalanced() bool {
	if m != nil {
		return m.Balanced
	}
	return false
}

func init() {
	proto.RegisterType((*BalanceStringRequest)(nil), "proto.BalanceStringRequest")
	proto.RegisterType((*BalanceStringResponse)(nil), "proto.BalanceStringResponse")
}

func init() { proto.RegisterFile("charbalancer.proto", fileDescriptor_7315142929410128) }

var fileDescriptor_7315142929410128 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2c,
	0x4a, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x4a, 0x5a, 0x5c, 0x22, 0x4e, 0x10, 0x89, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x98, 0x4b, 0x14, 0x4d, 0x6d, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0xd4, 0xf4, 0x14, 0xb0, 0x06, 0x8e, 0x20, 0x38, 0xdf,
	0x28, 0x86, 0x8b, 0xc7, 0x39, 0x23, 0xb1, 0x08, 0xaa, 0xb1, 0x48, 0xc8, 0x87, 0x8b, 0x17, 0xc5,
	0x10, 0x21, 0x69, 0x88, 0x83, 0xf4, 0xb0, 0x39, 0x43, 0x4a, 0x06, 0xbb, 0x24, 0xc4, 0x5e, 0x25,
	0x06, 0x27, 0x6d, 0x2e, 0xf5, 0xe4, 0xfc, 0x5c, 0xbd, 0xa4, 0x9c, 0xc4, 0xec, 0xd4, 0xec, 0x8c,
	0xc4, 0x3c, 0x3d, 0x14, 0x9f, 0x26, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x40, 0x8c, 0x08, 0x60, 0x4c,
	0x62, 0x03, 0x33, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x61, 0xbb, 0x67, 0x0f, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CharBalancerClient is the client API for CharBalancer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CharBalancerClient interface {
	BalanceString(ctx context.Context, in *BalanceStringRequest, opts ...grpc.CallOption) (*BalanceStringResponse, error)
}

type charBalancerClient struct {
	cc *grpc.ClientConn
}

func NewCharBalancerClient(cc *grpc.ClientConn) CharBalancerClient {
	return &charBalancerClient{cc}
}

func (c *charBalancerClient) BalanceString(ctx context.Context, in *BalanceStringRequest, opts ...grpc.CallOption) (*BalanceStringResponse, error) {
	out := new(BalanceStringResponse)
	err := c.cc.Invoke(ctx, "/proto.CharBalancer/BalanceString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharBalancerServer is the server API for CharBalancer service.
type CharBalancerServer interface {
	BalanceString(context.Context, *BalanceStringRequest) (*BalanceStringResponse, error)
}

// UnimplementedCharBalancerServer can be embedded to have forward compatible implementations.
type UnimplementedCharBalancerServer struct {
}

func (*UnimplementedCharBalancerServer) BalanceString(ctx context.Context, req *BalanceStringRequest) (*BalanceStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceString not implemented")
}

func RegisterCharBalancerServer(s *grpc.Server, srv CharBalancerServer) {
	s.RegisterService(&_CharBalancer_serviceDesc, srv)
}

func _CharBalancer_BalanceString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharBalancerServer).BalanceString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CharBalancer/BalanceString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharBalancerServer).BalanceString(ctx, req.(*BalanceStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CharBalancer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CharBalancer",
	HandlerType: (*CharBalancerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceString",
			Handler:    _CharBalancer_BalanceString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "charbalancer.proto",
}