// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: transmitter.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TransmitterService_Connect_FullMethodName    = "/transmitter.TransmitterService/Connect"
	TransmitterService_StreamData_FullMethodName = "/transmitter.TransmitterService/StreamData"
)

// TransmitterServiceClient is the client API for TransmitterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransmitterServiceClient interface {
	Connect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NewConnection, error)
	StreamData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TransmitterService_StreamDataClient, error)
}

type transmitterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransmitterServiceClient(cc grpc.ClientConnInterface) TransmitterServiceClient {
	return &transmitterServiceClient{cc}
}

func (c *transmitterServiceClient) Connect(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NewConnection, error) {
	out := new(NewConnection)
	err := c.cc.Invoke(ctx, TransmitterService_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transmitterServiceClient) StreamData(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (TransmitterService_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransmitterService_ServiceDesc.Streams[0], TransmitterService_StreamData_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transmitterServiceStreamDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransmitterService_StreamDataClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type transmitterServiceStreamDataClient struct {
	grpc.ClientStream
}

func (x *transmitterServiceStreamDataClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransmitterServiceServer is the server API for TransmitterService service.
// All implementations must embed UnimplementedTransmitterServiceServer
// for forward compatibility
type TransmitterServiceServer interface {
	Connect(context.Context, *emptypb.Empty) (*NewConnection, error)
	StreamData(*emptypb.Empty, TransmitterService_StreamDataServer) error
	//mustEmbedUnimplementedTransmitterServiceServer()
}

// UnimplementedTransmitterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransmitterServiceServer struct {
}

func (UnimplementedTransmitterServiceServer) Connect(context.Context, *emptypb.Empty) (*NewConnection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedTransmitterServiceServer) StreamData(*emptypb.Empty, TransmitterService_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedTransmitterServiceServer) mustEmbedUnimplementedTransmitterServiceServer() {}

// UnsafeTransmitterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransmitterServiceServer will
// result in compilation errors.
type UnsafeTransmitterServiceServer interface {
	mustEmbedUnimplementedTransmitterServiceServer()
}

func RegisterTransmitterServiceServer(s grpc.ServiceRegistrar, srv TransmitterServiceServer) {
	s.RegisterService(&TransmitterService_ServiceDesc, srv)
}

func _TransmitterService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransmitterServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransmitterService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransmitterServiceServer).Connect(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransmitterService_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransmitterServiceServer).StreamData(m, &transmitterServiceStreamDataServer{stream})
}

type TransmitterService_StreamDataServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type transmitterServiceStreamDataServer struct {
	grpc.ServerStream
}

func (x *transmitterServiceStreamDataServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// TransmitterService_ServiceDesc is the grpc.ServiceDesc for TransmitterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransmitterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transmitter.TransmitterService",
	HandlerType: (*TransmitterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _TransmitterService_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamData",
			Handler:       _TransmitterService_StreamData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transmitter.proto",
}