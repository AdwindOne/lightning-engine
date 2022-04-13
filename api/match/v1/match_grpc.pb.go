// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/match/v1/match.proto

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

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderReply, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderReply, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderReply, error) {
	out := new(AddOrderReply)
	err := c.cc.Invoke(ctx, "/api.match.v1.MatchService/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderReply, error) {
	out := new(CancelOrderReply)
	err := c.cc.Invoke(ctx, "/api.match.v1.MatchService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	AddOrder(context.Context, *AddOrderRequest) (*AddOrderReply, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderReply, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) AddOrder(context.Context, *AddOrderRequest) (*AddOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedMatchServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.match.v1.MatchService/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).AddOrder(ctx, req.(*AddOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.match.v1.MatchService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.match.v1.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _MatchService_AddOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _MatchService_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/match/v1/match.proto",
}