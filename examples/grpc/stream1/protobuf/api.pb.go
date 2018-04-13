// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Request
	Response
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

type Request struct {
	Req string `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type Response struct {
	Res string `protobuf:"bytes,2,opt,name=res" json:"res,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterType((*Response)(nil), "api.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Message service

type MessageClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (Message_ChatClient, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Message_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Message_serviceDesc.Streams[0], c.cc, "/api.Message/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageChatClient{stream}
	return x, nil
}

type Message_ChatClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type messageChatClient struct {
	grpc.ClientStream
}

func (x *messageChatClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageChatClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Message service

type MessageServer interface {
	Chat(Message_ChatServer) error
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServer).Chat(&messageChatServer{stream})
}

type Message_ChatServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type messageChatServer struct {
	grpc.ServerStream
}

func (x *messageChatServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageChatServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Message",
	HandlerType: (*MessageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Message_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x92, 0xe6, 0x62, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x4a, 0x2d, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x21, 0xb2, 0xc5, 0x12, 0x4c, 0x30, 0xd9, 0x62, 0x23, 0x13, 0x2e, 0x76, 0xdf, 0xd4, 0xe2,
	0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x4d, 0x2e, 0x16, 0xe7, 0x8c, 0xc4, 0x12, 0x21, 0x1e, 0x3d, 0x90,
	0xf1, 0x50, 0x03, 0xa5, 0x78, 0xa1, 0x3c, 0x88, 0x09, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49,
	0x6c, 0x60, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0x32, 0x26, 0x62, 0x89, 0x00,
	0x00, 0x00,
}
