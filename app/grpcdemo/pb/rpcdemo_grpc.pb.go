// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: manifest/protobuf/rpcdemo.proto

package pb

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

const (
	RpcDemo_Reply_FullMethodName = "/pb.RpcDemo/Reply"
)

// RpcDemoClient is the client API for RpcDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcDemoClient interface {
	Reply(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*ReplyResponse, error)
}

type rpcDemoClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcDemoClient(cc grpc.ClientConnInterface) RpcDemoClient {
	return &rpcDemoClient{cc}
}

func (c *rpcDemoClient) Reply(ctx context.Context, in *ReplyRequest, opts ...grpc.CallOption) (*ReplyResponse, error) {
	out := new(ReplyResponse)
	err := c.cc.Invoke(ctx, RpcDemo_Reply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcDemoServer is the server API for RpcDemo service.
// All implementations must embed UnimplementedRpcDemoServer
// for forward compatibility
type RpcDemoServer interface {
	Reply(context.Context, *ReplyRequest) (*ReplyResponse, error)
	mustEmbedUnimplementedRpcDemoServer()
}

// UnimplementedRpcDemoServer must be embedded to have forward compatible implementations.
type UnimplementedRpcDemoServer struct {
}

func (UnimplementedRpcDemoServer) Reply(context.Context, *ReplyRequest) (*ReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reply not implemented")
}
func (UnimplementedRpcDemoServer) mustEmbedUnimplementedRpcDemoServer() {}

// UnsafeRpcDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcDemoServer will
// result in compilation errors.
type UnsafeRpcDemoServer interface {
	mustEmbedUnimplementedRpcDemoServer()
}

func RegisterRpcDemoServer(s grpc.ServiceRegistrar, srv RpcDemoServer) {
	s.RegisterService(&RpcDemo_ServiceDesc, srv)
}

func _RpcDemo_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcDemoServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RpcDemo_Reply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcDemoServer).Reply(ctx, req.(*ReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcDemo_ServiceDesc is the grpc.ServiceDesc for RpcDemo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcDemo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RpcDemo",
	HandlerType: (*RpcDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reply",
			Handler:    _RpcDemo_Reply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manifest/protobuf/rpcdemo.proto",
}
