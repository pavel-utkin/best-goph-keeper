// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: internal/api/proto/gophkeeper.proto

package gophkeeper

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

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	HandleAuthentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error)
	HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error)
	HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error)
	HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) HandleAuthentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandleAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandleRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error) {
	out := new(CreateTextResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandleCreateText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error) {
	out := new(GetListTextResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandleGetListText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error) {
	out := new(GetNodeTextResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandleGetNodeText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/api.Gophkeeper/HandlePing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility
type GophkeeperServer interface {
	HandleAuthentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error)
	HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error)
	HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error)
	HandlePing(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophkeeperServer struct {
}

func (UnimplementedGophkeeperServer) HandleAuthentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleAuthentication not implemented")
}
func (UnimplementedGophkeeperServer) HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRegistration not implemented")
}
func (UnimplementedGophkeeperServer) HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetListText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetNodeText not implemented")
}
func (UnimplementedGophkeeperServer) HandlePing(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePing not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_HandleAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandleAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleAuthentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandleRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleRegistration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleCreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleCreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandleCreateText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleCreateText(ctx, req.(*CreateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetListText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetListText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandleGetListText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetListText(ctx, req.(*GetListTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetNodeText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandleGetNodeText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, req.(*GetNodeTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandlePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandlePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Gophkeeper/HandlePing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandlePing(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleAuthentication",
			Handler:    _Gophkeeper_HandleAuthentication_Handler,
		},
		{
			MethodName: "HandleRegistration",
			Handler:    _Gophkeeper_HandleRegistration_Handler,
		},
		{
			MethodName: "HandleCreateText",
			Handler:    _Gophkeeper_HandleCreateText_Handler,
		},
		{
			MethodName: "HandleGetListText",
			Handler:    _Gophkeeper_HandleGetListText_Handler,
		},
		{
			MethodName: "HandleGetNodeText",
			Handler:    _Gophkeeper_HandleGetNodeText_Handler,
		},
		{
			MethodName: "HandlePing",
			Handler:    _Gophkeeper_HandlePing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/proto/gophkeeper.proto",
}
