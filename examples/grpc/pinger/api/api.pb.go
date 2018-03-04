// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GreetMessage
	RequestDate
	DateMessage
*/
package api

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

type GreetMessage struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetMessage) Reset()                    { *m = GreetMessage{} }
func (m *GreetMessage) String() string            { return proto.CompactTextString(m) }
func (*GreetMessage) ProtoMessage()               {}
func (*GreetMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GreetMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type RequestDate struct {
	Req string `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
}

func (m *RequestDate) Reset()                    { *m = RequestDate{} }
func (m *RequestDate) String() string            { return proto.CompactTextString(m) }
func (*RequestDate) ProtoMessage()               {}
func (*RequestDate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestDate) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type DateMessage struct {
	Day   int32 `protobuf:"varint,1,opt,name=day" json:"day,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Year  int32 `protobuf:"varint,3,opt,name=year" json:"year,omitempty"`
}

func (m *DateMessage) Reset()                    { *m = DateMessage{} }
func (m *DateMessage) String() string            { return proto.CompactTextString(m) }
func (*DateMessage) ProtoMessage()               {}
func (*DateMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DateMessage) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *DateMessage) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *DateMessage) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func init() {
	proto.RegisterType((*GreetMessage)(nil), "api.GreetMessage")
	proto.RegisterType((*RequestDate)(nil), "api.RequestDate")
	proto.RegisterType((*DateMessage)(nil), "api.DateMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ping service

type PingClient interface {
	CheckIn(ctx context.Context, in *GreetMessage, opts ...grpc.CallOption) (*GreetMessage, error)
	GetDate(ctx context.Context, in *RequestDate, opts ...grpc.CallOption) (*DateMessage, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) CheckIn(ctx context.Context, in *GreetMessage, opts ...grpc.CallOption) (*GreetMessage, error) {
	out := new(GreetMessage)
	err := grpc.Invoke(ctx, "/api.Ping/CheckIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingClient) GetDate(ctx context.Context, in *RequestDate, opts ...grpc.CallOption) (*DateMessage, error) {
	out := new(DateMessage)
	err := grpc.Invoke(ctx, "/api.Ping/GetDate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingServer interface {
	CheckIn(context.Context, *GreetMessage) (*GreetMessage, error)
	GetDate(context.Context, *RequestDate) (*DateMessage, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_CheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).CheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Ping/CheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).CheckIn(ctx, req.(*GreetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ping_GetDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).GetDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Ping/GetDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).GetDate(ctx, req.(*RequestDate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckIn",
			Handler:    _Ping_CheckIn_Handler,
		},
		{
			MethodName: "GetDate",
			Handler:    _Ping_GetDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xd2, 0xe2, 0xe2, 0x71, 0x2f,
	0x4a, 0x4d, 0x2d, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe2, 0xe2, 0x48, 0x07,
	0xf1, 0x33, 0xf3, 0xd2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x79, 0x2e,
	0xee, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x97, 0xc4, 0x92, 0x54, 0x21, 0x01, 0x2e, 0xe6,
	0xa2, 0xd4, 0x42, 0xa8, 0x2a, 0x10, 0x53, 0xc9, 0x93, 0x8b, 0x1b, 0x24, 0x03, 0x33, 0x4b, 0x80,
	0x8b, 0x39, 0x25, 0xb1, 0x12, 0xac, 0x80, 0x35, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0xcd, 0xcd,
	0xcf, 0x2b, 0xc9, 0x90, 0x60, 0x02, 0x8b, 0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x95, 0xa9, 0x89,
	0x45, 0x12, 0xcc, 0x60, 0x41, 0x30, 0xdb, 0x28, 0x8b, 0x8b, 0x25, 0x20, 0x33, 0x2f, 0x5d, 0xc8,
	0x90, 0x8b, 0xdd, 0x39, 0x23, 0x35, 0x39, 0xdb, 0x33, 0x4f, 0x48, 0x50, 0x0f, 0xe4, 0x76, 0x64,
	0xd7, 0x4a, 0x61, 0x0a, 0x29, 0x31, 0x08, 0xe9, 0x73, 0xb1, 0xbb, 0xa7, 0x42, 0x9d, 0x08, 0x96,
	0x47, 0x72, 0xb4, 0x14, 0x44, 0x04, 0xc9, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xf0, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x6f, 0x41, 0x78, 0x1c, 0x01, 0x00, 0x00,
}
