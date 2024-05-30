// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/grpc/bfa_protection.proto

package bfa_protection

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BFAProtection_Authorize_FullMethodName       = "/bfa_protection.BFAProtection/Authorize"
	BFAProtection_FlushBucket_FullMethodName     = "/bfa_protection.BFAProtection/FlushBucket"
	BFAProtection_BlacklistAdd_FullMethodName    = "/bfa_protection.BFAProtection/BlacklistAdd"
	BFAProtection_BlacklistRemove_FullMethodName = "/bfa_protection.BFAProtection/BlacklistRemove"
	BFAProtection_WhitelistAdd_FullMethodName    = "/bfa_protection.BFAProtection/WhitelistAdd"
	BFAProtection_WhitelistRemove_FullMethodName = "/bfa_protection.BFAProtection/WhitelistRemove"
)

// BFAProtectionClient is the client API for BFAProtection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFAProtectionClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	FlushBucket(ctx context.Context, in *FlushBucketRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	BlacklistAdd(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	BlacklistRemove(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	WhitelistAdd(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	WhitelistRemove(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bFAProtectionClient struct {
	cc grpc.ClientConnInterface
}

func NewBFAProtectionClient(cc grpc.ClientConnInterface) BFAProtectionClient {
	return &bFAProtectionClient{cc}
}

func (c *bFAProtectionClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, BFAProtection_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFAProtectionClient) FlushBucket(ctx context.Context, in *FlushBucketRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BFAProtection_FlushBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFAProtectionClient) BlacklistAdd(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BFAProtection_BlacklistAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFAProtectionClient) BlacklistRemove(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BFAProtection_BlacklistRemove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFAProtectionClient) WhitelistAdd(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BFAProtection_WhitelistAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFAProtectionClient) WhitelistRemove(ctx context.Context, in *IPRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BFAProtection_WhitelistRemove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFAProtectionServer is the server API for BFAProtection service.
// All implementations must embed UnimplementedBFAProtectionServer
// for forward compatibility
type BFAProtectionServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	FlushBucket(context.Context, *FlushBucketRequest) (*empty.Empty, error)
	BlacklistAdd(context.Context, *IPRequest) (*empty.Empty, error)
	BlacklistRemove(context.Context, *IPRequest) (*empty.Empty, error)
	WhitelistAdd(context.Context, *IPRequest) (*empty.Empty, error)
	WhitelistRemove(context.Context, *IPRequest) (*empty.Empty, error)
	mustEmbedUnimplementedBFAProtectionServer()
}

// UnimplementedBFAProtectionServer must be embedded to have forward compatible implementations.
type UnimplementedBFAProtectionServer struct {
}

func (UnimplementedBFAProtectionServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedBFAProtectionServer) FlushBucket(context.Context, *FlushBucketRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushBucket not implemented")
}
func (UnimplementedBFAProtectionServer) BlacklistAdd(context.Context, *IPRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistAdd not implemented")
}
func (UnimplementedBFAProtectionServer) BlacklistRemove(context.Context, *IPRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistRemove not implemented")
}
func (UnimplementedBFAProtectionServer) WhitelistAdd(context.Context, *IPRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistAdd not implemented")
}
func (UnimplementedBFAProtectionServer) WhitelistRemove(context.Context, *IPRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhitelistRemove not implemented")
}
func (UnimplementedBFAProtectionServer) mustEmbedUnimplementedBFAProtectionServer() {}

// UnsafeBFAProtectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFAProtectionServer will
// result in compilation errors.
type UnsafeBFAProtectionServer interface {
	mustEmbedUnimplementedBFAProtectionServer()
}

func RegisterBFAProtectionServer(s grpc.ServiceRegistrar, srv BFAProtectionServer) {
	s.RegisterService(&BFAProtection_ServiceDesc, srv)
}

func _BFAProtection_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFAProtection_FlushBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).FlushBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_FlushBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).FlushBucket(ctx, req.(*FlushBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFAProtection_BlacklistAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).BlacklistAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_BlacklistAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).BlacklistAdd(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFAProtection_BlacklistRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).BlacklistRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_BlacklistRemove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).BlacklistRemove(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFAProtection_WhitelistAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).WhitelistAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_WhitelistAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).WhitelistAdd(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFAProtection_WhitelistRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFAProtectionServer).WhitelistRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BFAProtection_WhitelistRemove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFAProtectionServer).WhitelistRemove(ctx, req.(*IPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFAProtection_ServiceDesc is the grpc.ServiceDesc for BFAProtection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFAProtection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bfa_protection.BFAProtection",
	HandlerType: (*BFAProtectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _BFAProtection_Authorize_Handler,
		},
		{
			MethodName: "FlushBucket",
			Handler:    _BFAProtection_FlushBucket_Handler,
		},
		{
			MethodName: "BlacklistAdd",
			Handler:    _BFAProtection_BlacklistAdd_Handler,
		},
		{
			MethodName: "BlacklistRemove",
			Handler:    _BFAProtection_BlacklistRemove_Handler,
		},
		{
			MethodName: "WhitelistAdd",
			Handler:    _BFAProtection_WhitelistAdd_Handler,
		},
		{
			MethodName: "WhitelistRemove",
			Handler:    _BFAProtection_WhitelistRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/bfa_protection.proto",
}
