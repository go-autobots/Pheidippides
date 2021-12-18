// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// PheidiQueueClient is the client API for PheidiQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PheidiQueueClient interface {
	SendTo(ctx context.Context, in *Pheidippides, opts ...grpc.CallOption) (*SendResponse, error)
	GetFrom(ctx context.Context, in *Pheidippides, opts ...grpc.CallOption) (*PhiediResponse, error)
}

type pheidiQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewPheidiQueueClient(cc grpc.ClientConnInterface) PheidiQueueClient {
	return &pheidiQueueClient{cc}
}

func (c *pheidiQueueClient) SendTo(ctx context.Context, in *Pheidippides, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/pheidippides.v1.PheidiQueue/SendTo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pheidiQueueClient) GetFrom(ctx context.Context, in *Pheidippides, opts ...grpc.CallOption) (*PhiediResponse, error) {
	out := new(PhiediResponse)
	err := c.cc.Invoke(ctx, "/pheidippides.v1.PheidiQueue/GetFrom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PheidiQueueServer is the server API for PheidiQueue service.
// All implementations must embed UnimplementedPheidiQueueServer
// for forward compatibility
type PheidiQueueServer interface {
	SendTo(context.Context, *Pheidippides) (*SendResponse, error)
	GetFrom(context.Context, *Pheidippides) (*PhiediResponse, error)
	mustEmbedUnimplementedPheidiQueueServer()
}

// UnimplementedPheidiQueueServer must be embedded to have forward compatible implementations.
type UnimplementedPheidiQueueServer struct {
}

func (UnimplementedPheidiQueueServer) SendTo(context.Context, *Pheidippides) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTo not implemented")
}
func (UnimplementedPheidiQueueServer) GetFrom(context.Context, *Pheidippides) (*PhiediResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrom not implemented")
}
func (UnimplementedPheidiQueueServer) mustEmbedUnimplementedPheidiQueueServer() {}

// UnsafePheidiQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PheidiQueueServer will
// result in compilation errors.
type UnsafePheidiQueueServer interface {
	mustEmbedUnimplementedPheidiQueueServer()
}

func RegisterPheidiQueueServer(s grpc.ServiceRegistrar, srv PheidiQueueServer) {
	s.RegisterService(&PheidiQueue_ServiceDesc, srv)
}

func _PheidiQueue_SendTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pheidippides)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheidiQueueServer).SendTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheidippides.v1.PheidiQueue/SendTo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheidiQueueServer).SendTo(ctx, req.(*Pheidippides))
	}
	return interceptor(ctx, in, info, handler)
}

func _PheidiQueue_GetFrom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pheidippides)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PheidiQueueServer).GetFrom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pheidippides.v1.PheidiQueue/GetFrom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PheidiQueueServer).GetFrom(ctx, req.(*Pheidippides))
	}
	return interceptor(ctx, in, info, handler)
}

// PheidiQueue_ServiceDesc is the grpc.ServiceDesc for PheidiQueue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PheidiQueue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pheidippides.v1.PheidiQueue",
	HandlerType: (*PheidiQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTo",
			Handler:    _PheidiQueue_SendTo_Handler,
		},
		{
			MethodName: "GetFrom",
			Handler:    _PheidiQueue_GetFrom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pheidiqueue/v1/queue.proto",
}
