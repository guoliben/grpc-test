// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: RouteGruid.proto

package hello_grpc_proto

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

// RouteGuidClient is the client API for RouteGuid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuidClient interface {
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RouteChatClient, error)
}

type routeGuidClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuidClient(cc grpc.ClientConnInterface) RouteGuidClient {
	return &routeGuidClient{cc}
}

func (c *routeGuidClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuid_ServiceDesc.Streams[0], "/service.RouteGuid/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuidRouteChatClient{stream}
	return x, nil
}

type RouteGuid_RouteChatClient interface {
	Send(*RouteNoteRequest) error
	Recv() (*RouteNoteResponse, error)
	grpc.ClientStream
}

type routeGuidRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuidRouteChatClient) Send(m *RouteNoteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuidRouteChatClient) Recv() (*RouteNoteResponse, error) {
	m := new(RouteNoteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuidServer is the server API for RouteGuid service.
// All implementations must embed UnimplementedRouteGuidServer
// for forward compatibility
type RouteGuidServer interface {
	RouteChat(RouteGuid_RouteChatServer) error
	mustEmbedUnimplementedRouteGuidServer()
}

// UnimplementedRouteGuidServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuidServer struct {
}

func (UnimplementedRouteGuidServer) RouteChat(RouteGuid_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}
func (UnimplementedRouteGuidServer) mustEmbedUnimplementedRouteGuidServer() {}

// UnsafeRouteGuidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuidServer will
// result in compilation errors.
type UnsafeRouteGuidServer interface {
	mustEmbedUnimplementedRouteGuidServer()
}

func RegisterRouteGuidServer(s grpc.ServiceRegistrar, srv RouteGuidServer) {
	s.RegisterService(&RouteGuid_ServiceDesc, srv)
}

func _RouteGuid_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuidServer).RouteChat(&routeGuidRouteChatServer{stream})
}

type RouteGuid_RouteChatServer interface {
	Send(*RouteNoteResponse) error
	Recv() (*RouteNoteRequest, error)
	grpc.ServerStream
}

type routeGuidRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuidRouteChatServer) Send(m *RouteNoteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuidRouteChatServer) Recv() (*RouteNoteRequest, error) {
	m := new(RouteNoteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuid_ServiceDesc is the grpc.ServiceDesc for RouteGuid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.RouteGuid",
	HandlerType: (*RouteGuidServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuid_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "RouteGruid.proto",
}