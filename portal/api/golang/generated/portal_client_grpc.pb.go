// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: portal_client.proto

package golang

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

// KurtosisPortalClientClient is the client API for KurtosisPortalClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KurtosisPortalClientClient interface {
	// To check availability
	Ping(ctx context.Context, in *PortalPing, opts ...grpc.CallOption) (*PortalPong, error)
	// SwitchContext switches the current context used by Kurtosis.
	//
	// If the new context is a dual-backend-context, it connects to it automatically using the connection information
	// linked to the context Right now, it is expected that the remote environment is running a Kurtosis Portal Server
	// on port 9720
	SwitchContext(ctx context.Context, in *SwitchContextArgs, opts ...grpc.CallOption) (*SwitchContextResponse, error)
	// TODO: Raw endpoint to forward a port from server to client. This is very low level, in the future the portal
	//  should accept higher level info, like (enclave_id, service_id, port_id) and automatically find the ephemeral
	//  port number.
	ForwardPort(ctx context.Context, in *ForwardPortArgs, opts ...grpc.CallOption) (*ForwardPortResponse, error)
}

type kurtosisPortalClientClient struct {
	cc grpc.ClientConnInterface
}

func NewKurtosisPortalClientClient(cc grpc.ClientConnInterface) KurtosisPortalClientClient {
	return &kurtosisPortalClientClient{cc}
}

func (c *kurtosisPortalClientClient) Ping(ctx context.Context, in *PortalPing, opts ...grpc.CallOption) (*PortalPong, error) {
	out := new(PortalPong)
	err := c.cc.Invoke(ctx, "/kurtosis_portal_daemon.KurtosisPortalClient/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisPortalClientClient) SwitchContext(ctx context.Context, in *SwitchContextArgs, opts ...grpc.CallOption) (*SwitchContextResponse, error) {
	out := new(SwitchContextResponse)
	err := c.cc.Invoke(ctx, "/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kurtosisPortalClientClient) ForwardPort(ctx context.Context, in *ForwardPortArgs, opts ...grpc.CallOption) (*ForwardPortResponse, error) {
	out := new(ForwardPortResponse)
	err := c.cc.Invoke(ctx, "/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KurtosisPortalClientServer is the server API for KurtosisPortalClient service.
// All implementations should embed UnimplementedKurtosisPortalClientServer
// for forward compatibility
type KurtosisPortalClientServer interface {
	// To check availability
	Ping(context.Context, *PortalPing) (*PortalPong, error)
	// SwitchContext switches the current context used by Kurtosis.
	//
	// If the new context is a dual-backend-context, it connects to it automatically using the connection information
	// linked to the context Right now, it is expected that the remote environment is running a Kurtosis Portal Server
	// on port 9720
	SwitchContext(context.Context, *SwitchContextArgs) (*SwitchContextResponse, error)
	// TODO: Raw endpoint to forward a port from server to client. This is very low level, in the future the portal
	//  should accept higher level info, like (enclave_id, service_id, port_id) and automatically find the ephemeral
	//  port number.
	ForwardPort(context.Context, *ForwardPortArgs) (*ForwardPortResponse, error)
}

// UnimplementedKurtosisPortalClientServer should be embedded to have forward compatible implementations.
type UnimplementedKurtosisPortalClientServer struct {
}

func (UnimplementedKurtosisPortalClientServer) Ping(context.Context, *PortalPing) (*PortalPong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedKurtosisPortalClientServer) SwitchContext(context.Context, *SwitchContextArgs) (*SwitchContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchContext not implemented")
}
func (UnimplementedKurtosisPortalClientServer) ForwardPort(context.Context, *ForwardPortArgs) (*ForwardPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardPort not implemented")
}

// UnsafeKurtosisPortalClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KurtosisPortalClientServer will
// result in compilation errors.
type UnsafeKurtosisPortalClientServer interface {
	mustEmbedUnimplementedKurtosisPortalClientServer()
}

func RegisterKurtosisPortalClientServer(s grpc.ServiceRegistrar, srv KurtosisPortalClientServer) {
	s.RegisterService(&KurtosisPortalClient_ServiceDesc, srv)
}

func _KurtosisPortalClient_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortalPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisPortalClientServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kurtosis_portal_daemon.KurtosisPortalClient/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisPortalClientServer).Ping(ctx, req.(*PortalPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisPortalClient_SwitchContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchContextArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisPortalClientServer).SwitchContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kurtosis_portal_daemon.KurtosisPortalClient/SwitchContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisPortalClientServer).SwitchContext(ctx, req.(*SwitchContextArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _KurtosisPortalClient_ForwardPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardPortArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KurtosisPortalClientServer).ForwardPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kurtosis_portal_daemon.KurtosisPortalClient/ForwardPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KurtosisPortalClientServer).ForwardPort(ctx, req.(*ForwardPortArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// KurtosisPortalClient_ServiceDesc is the grpc.ServiceDesc for KurtosisPortalClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KurtosisPortalClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kurtosis_portal_daemon.KurtosisPortalClient",
	HandlerType: (*KurtosisPortalClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _KurtosisPortalClient_Ping_Handler,
		},
		{
			MethodName: "SwitchContext",
			Handler:    _KurtosisPortalClient_SwitchContext_Handler,
		},
		{
			MethodName: "ForwardPort",
			Handler:    _KurtosisPortalClient_ForwardPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal_client.proto",
}
