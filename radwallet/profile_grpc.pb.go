// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: profile.proto

package radwallet

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

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	UpdateProfile(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	GetProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error)
	GetProfileInternal(ctx context.Context, in *ExchangeDataRequest, opts ...grpc.CallOption) (*ExchangeUser, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) UpdateProfile(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/radhark.radwallet.Profile/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/radhark.radwallet.Profile/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileInternal(ctx context.Context, in *ExchangeDataRequest, opts ...grpc.CallOption) (*ExchangeUser, error) {
	out := new(ExchangeUser)
	err := c.cc.Invoke(ctx, "/radhark.radwallet.Profile/GetProfileInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	UpdateProfile(context.Context, *ProfileUpdateRequest) (*ProfileResponse, error)
	GetProfile(context.Context, *empty.Empty) (*ProfileResponse, error)
	GetProfileInternal(context.Context, *ExchangeDataRequest) (*ExchangeUser, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) UpdateProfile(context.Context, *ProfileUpdateRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServer) GetProfile(context.Context, *empty.Empty) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServer) GetProfileInternal(context.Context, *ExchangeDataRequest) (*ExchangeUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileInternal not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/radhark.radwallet.Profile/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateProfile(ctx, req.(*ProfileUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/radhark.radwallet.Profile/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/radhark.radwallet.Profile/GetProfileInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileInternal(ctx, req.(*ExchangeDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "radhark.radwallet.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateProfile",
			Handler:    _Profile_UpdateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Profile_GetProfile_Handler,
		},
		{
			MethodName: "GetProfileInternal",
			Handler:    _Profile_GetProfileInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
