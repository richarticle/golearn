// Code generated by protoc-gen-go.
// source: protocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	Request
	Response
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Numbers []int32 `protobuf:"varint,1,rep,packed,name=numbers" json:"numbers,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Sum int32 `protobuf:"varint,1,opt,name=sum" json:"sum,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Response)(nil), "protocol.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Summation service

type SummationClient interface {
	ComputeSum(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type summationClient struct {
	cc *grpc.ClientConn
}

func NewSummationClient(cc *grpc.ClientConn) SummationClient {
	return &summationClient{cc}
}

func (c *summationClient) ComputeSum(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protocol.Summation/ComputeSum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Summation service

type SummationServer interface {
	ComputeSum(context.Context, *Request) (*Response, error)
}

func RegisterSummationServer(s *grpc.Server, srv SummationServer) {
	s.RegisterService(&_Summation_serviceDesc, srv)
}

func _Summation_ComputeSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummationServer).ComputeSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Summation/ComputeSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummationServer).ComputeSum(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Summation",
	HandlerType: (*SummationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeSum",
			Handler:    _Summation_ComputeSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x84, 0x38, 0x60, 0x7c, 0x25, 0x65, 0x2e, 0xf6, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0xa2,
	0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0x49, 0x86, 0x8b, 0x23, 0x28, 0xb5,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b, 0xb9, 0xb8, 0x34, 0x17, 0xa8, 0x82, 0x11,
	0xa8, 0x02, 0xc4, 0x34, 0x72, 0xe2, 0xe2, 0x0c, 0x2e, 0xcd, 0xcd, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf,
	0x13, 0x32, 0xe5, 0xe2, 0x72, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x05, 0x8a, 0x09, 0x09, 0xea,
	0xc1, 0x2d, 0x86, 0xda, 0x22, 0x25, 0x84, 0x2c, 0x04, 0x31, 0x53, 0x89, 0x21, 0x89, 0x0d, 0x2c,
	0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x92, 0x59, 0x08, 0xfa, 0xa9, 0x00, 0x00, 0x00,
}
