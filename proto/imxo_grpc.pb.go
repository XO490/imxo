// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: imxo.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SenderClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Sender_SendMessageClient, error)
}

type senderClient struct {
	cc grpc.ClientConnInterface
}

func NewSenderClient(cc grpc.ClientConnInterface) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Sender_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sender_ServiceDesc.Streams[0], "/imxo.Sender/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &senderSendMessageClient{stream}
	return x, nil
}

type Sender_SendMessageClient interface {
	Send(*FromClient) error
	Recv() (*FromServer, error)
	grpc.ClientStream
}

type senderSendMessageClient struct {
	grpc.ClientStream
}

func (x *senderSendMessageClient) Send(m *FromClient) error {
	return x.ClientStream.SendMsg(m)
}

func (x *senderSendMessageClient) Recv() (*FromServer, error) {
	m := new(FromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SenderServer is the server API for Sender service.
// All implementations must embed UnimplementedSenderServer
// for forward compatibility
type SenderServer interface {
	SendMessage(Sender_SendMessageServer) error
	mustEmbedUnimplementedSenderServer()
}

// UnimplementedSenderServer must be embedded to have forward compatible implementations.
type UnimplementedSenderServer struct {
}

func (UnimplementedSenderServer) SendMessage(Sender_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedSenderServer) mustEmbedUnimplementedSenderServer() {}

// UnsafeSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SenderServer will
// result in compilation errors.
type UnsafeSenderServer interface {
	mustEmbedUnimplementedSenderServer()
}

func RegisterSenderServer(s grpc.ServiceRegistrar, srv SenderServer) {
	s.RegisterService(&Sender_ServiceDesc, srv)
}

func _Sender_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SenderServer).SendMessage(&senderSendMessageServer{stream})
}

type Sender_SendMessageServer interface {
	Send(*FromServer) error
	Recv() (*FromClient, error)
	grpc.ServerStream
}

type senderSendMessageServer struct {
	grpc.ServerStream
}

func (x *senderSendMessageServer) Send(m *FromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *senderSendMessageServer) Recv() (*FromClient, error) {
	m := new(FromClient)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Sender_ServiceDesc is the grpc.ServiceDesc for Sender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imxo.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _Sender_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "imxo.proto",
}