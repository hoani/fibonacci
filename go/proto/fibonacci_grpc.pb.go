// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fibonacci

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

// FibonacciClient is the client API for Fibonacci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibonacciClient interface {
	AtIndex(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error)
	GetSequence(ctx context.Context, in *Number, opts ...grpc.CallOption) (Fibonacci_GetSequenceClient, error)
	SumIndicies(ctx context.Context, opts ...grpc.CallOption) (Fibonacci_SumIndiciesClient, error)
	StreamSequence(ctx context.Context, opts ...grpc.CallOption) (Fibonacci_StreamSequenceClient, error)
}

type fibonacciClient struct {
	cc grpc.ClientConnInterface
}

func NewFibonacciClient(cc grpc.ClientConnInterface) FibonacciClient {
	return &fibonacciClient{cc}
}

func (c *fibonacciClient) AtIndex(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/fibonacci.Fibonacci/AtIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fibonacciClient) GetSequence(ctx context.Context, in *Number, opts ...grpc.CallOption) (Fibonacci_GetSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fibonacci_ServiceDesc.Streams[0], "/fibonacci.Fibonacci/GetSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &fibonacciGetSequenceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fibonacci_GetSequenceClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type fibonacciGetSequenceClient struct {
	grpc.ClientStream
}

func (x *fibonacciGetSequenceClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fibonacciClient) SumIndicies(ctx context.Context, opts ...grpc.CallOption) (Fibonacci_SumIndiciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fibonacci_ServiceDesc.Streams[1], "/fibonacci.Fibonacci/SumIndicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &fibonacciSumIndiciesClient{stream}
	return x, nil
}

type Fibonacci_SumIndiciesClient interface {
	Send(*Number) error
	CloseAndRecv() (*Number, error)
	grpc.ClientStream
}

type fibonacciSumIndiciesClient struct {
	grpc.ClientStream
}

func (x *fibonacciSumIndiciesClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fibonacciSumIndiciesClient) CloseAndRecv() (*Number, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fibonacciClient) StreamSequence(ctx context.Context, opts ...grpc.CallOption) (Fibonacci_StreamSequenceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fibonacci_ServiceDesc.Streams[2], "/fibonacci.Fibonacci/StreamSequence", opts...)
	if err != nil {
		return nil, err
	}
	x := &fibonacciStreamSequenceClient{stream}
	return x, nil
}

type Fibonacci_StreamSequenceClient interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ClientStream
}

type fibonacciStreamSequenceClient struct {
	grpc.ClientStream
}

func (x *fibonacciStreamSequenceClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fibonacciStreamSequenceClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FibonacciServer is the server API for Fibonacci service.
// All implementations must embed UnimplementedFibonacciServer
// for forward compatibility
type FibonacciServer interface {
	AtIndex(context.Context, *Number) (*Number, error)
	GetSequence(*Number, Fibonacci_GetSequenceServer) error
	SumIndicies(Fibonacci_SumIndiciesServer) error
	StreamSequence(Fibonacci_StreamSequenceServer) error
	mustEmbedUnimplementedFibonacciServer()
}

// UnimplementedFibonacciServer must be embedded to have forward compatible implementations.
type UnimplementedFibonacciServer struct {
}

func (UnimplementedFibonacciServer) AtIndex(context.Context, *Number) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtIndex not implemented")
}
func (UnimplementedFibonacciServer) GetSequence(*Number, Fibonacci_GetSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSequence not implemented")
}
func (UnimplementedFibonacciServer) SumIndicies(Fibonacci_SumIndiciesServer) error {
	return status.Errorf(codes.Unimplemented, "method SumIndicies not implemented")
}
func (UnimplementedFibonacciServer) StreamSequence(Fibonacci_StreamSequenceServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSequence not implemented")
}
func (UnimplementedFibonacciServer) mustEmbedUnimplementedFibonacciServer() {}

// UnsafeFibonacciServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibonacciServer will
// result in compilation errors.
type UnsafeFibonacciServer interface {
	mustEmbedUnimplementedFibonacciServer()
}

func RegisterFibonacciServer(s grpc.ServiceRegistrar, srv FibonacciServer) {
	s.RegisterService(&Fibonacci_ServiceDesc, srv)
}

func _Fibonacci_AtIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibonacciServer).AtIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fibonacci.Fibonacci/AtIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibonacciServer).AtIndex(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fibonacci_GetSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Number)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FibonacciServer).GetSequence(m, &fibonacciGetSequenceServer{stream})
}

type Fibonacci_GetSequenceServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type fibonacciGetSequenceServer struct {
	grpc.ServerStream
}

func (x *fibonacciGetSequenceServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func _Fibonacci_SumIndicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FibonacciServer).SumIndicies(&fibonacciSumIndiciesServer{stream})
}

type Fibonacci_SumIndiciesServer interface {
	SendAndClose(*Number) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type fibonacciSumIndiciesServer struct {
	grpc.ServerStream
}

func (x *fibonacciSumIndiciesServer) SendAndClose(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fibonacciSumIndiciesServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fibonacci_StreamSequence_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FibonacciServer).StreamSequence(&fibonacciStreamSequenceServer{stream})
}

type Fibonacci_StreamSequenceServer interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type fibonacciStreamSequenceServer struct {
	grpc.ServerStream
}

func (x *fibonacciStreamSequenceServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fibonacciStreamSequenceServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Fibonacci_ServiceDesc is the grpc.ServiceDesc for Fibonacci service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fibonacci_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fibonacci.Fibonacci",
	HandlerType: (*FibonacciServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AtIndex",
			Handler:    _Fibonacci_AtIndex_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSequence",
			Handler:       _Fibonacci_GetSequence_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SumIndicies",
			Handler:       _Fibonacci_SumIndicies_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSequence",
			Handler:       _Fibonacci_StreamSequence_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/fibonacci.proto",
}
