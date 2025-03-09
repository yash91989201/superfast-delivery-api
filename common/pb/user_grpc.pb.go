// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_CreateProfile_FullMethodName         = "/pb.UserService/CreateProfile"
	UserService_GetProfile_FullMethodName            = "/pb.UserService/GetProfile"
	UserService_UpdateProfile_FullMethodName         = "/pb.UserService/UpdateProfile"
	UserService_DeleteProfile_FullMethodName         = "/pb.UserService/DeleteProfile"
	UserService_CreateDeliveryAddress_FullMethodName = "/pb.UserService/CreateDeliveryAddress"
	UserService_GetDeliveryAddress_FullMethodName    = "/pb.UserService/GetDeliveryAddress"
	UserService_ListDeliveryAddress_FullMethodName   = "/pb.UserService/ListDeliveryAddress"
	UserService_UpdateDeliveryAddress_FullMethodName = "/pb.UserService/UpdateDeliveryAddress"
	UserService_DeleteDeliveryAddress_FullMethodName = "/pb.UserService/DeleteDeliveryAddress"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateProfile(ctx context.Context, in *CreateProfileReq, opts ...grpc.CallOption) (*Profile, error)
	GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*Profile, error)
	DeleteProfile(ctx context.Context, in *DeleteProfileReq, opts ...grpc.CallOption) (*EmptyRes, error)
	CreateDeliveryAddress(ctx context.Context, in *CreateDeliveryAddressReq, opts ...grpc.CallOption) (*DeliveryAddress, error)
	GetDeliveryAddress(ctx context.Context, in *GetDeliveryAddressReq, opts ...grpc.CallOption) (*DeliveryAddress, error)
	ListDeliveryAddress(ctx context.Context, in *ListDeliveryAddressReq, opts ...grpc.CallOption) (*ListDeliveryAddressRes, error)
	UpdateDeliveryAddress(ctx context.Context, in *DeliveryAddress, opts ...grpc.CallOption) (*DeliveryAddress, error)
	DeleteDeliveryAddress(ctx context.Context, in *DeleteDeliveryAddressReq, opts ...grpc.CallOption) (*EmptyRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateProfile(ctx context.Context, in *CreateProfileReq, opts ...grpc.CallOption) (*Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profile)
	err := c.cc.Invoke(ctx, UserService_CreateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetProfileReq, opts ...grpc.CallOption) (*Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profile)
	err := c.cc.Invoke(ctx, UserService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Profile)
	err := c.cc.Invoke(ctx, UserService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteProfile(ctx context.Context, in *DeleteProfileReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, UserService_DeleteProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateDeliveryAddress(ctx context.Context, in *CreateDeliveryAddressReq, opts ...grpc.CallOption) (*DeliveryAddress, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryAddress)
	err := c.cc.Invoke(ctx, UserService_CreateDeliveryAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetDeliveryAddress(ctx context.Context, in *GetDeliveryAddressReq, opts ...grpc.CallOption) (*DeliveryAddress, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryAddress)
	err := c.cc.Invoke(ctx, UserService_GetDeliveryAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListDeliveryAddress(ctx context.Context, in *ListDeliveryAddressReq, opts ...grpc.CallOption) (*ListDeliveryAddressRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDeliveryAddressRes)
	err := c.cc.Invoke(ctx, UserService_ListDeliveryAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateDeliveryAddress(ctx context.Context, in *DeliveryAddress, opts ...grpc.CallOption) (*DeliveryAddress, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryAddress)
	err := c.cc.Invoke(ctx, UserService_UpdateDeliveryAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteDeliveryAddress(ctx context.Context, in *DeleteDeliveryAddressReq, opts ...grpc.CallOption) (*EmptyRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyRes)
	err := c.cc.Invoke(ctx, UserService_DeleteDeliveryAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	CreateProfile(context.Context, *CreateProfileReq) (*Profile, error)
	GetProfile(context.Context, *GetProfileReq) (*Profile, error)
	UpdateProfile(context.Context, *UpdateProfileReq) (*Profile, error)
	DeleteProfile(context.Context, *DeleteProfileReq) (*EmptyRes, error)
	CreateDeliveryAddress(context.Context, *CreateDeliveryAddressReq) (*DeliveryAddress, error)
	GetDeliveryAddress(context.Context, *GetDeliveryAddressReq) (*DeliveryAddress, error)
	ListDeliveryAddress(context.Context, *ListDeliveryAddressReq) (*ListDeliveryAddressRes, error)
	UpdateDeliveryAddress(context.Context, *DeliveryAddress) (*DeliveryAddress, error)
	DeleteDeliveryAddress(context.Context, *DeleteDeliveryAddressReq) (*EmptyRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateProfile(context.Context, *CreateProfileReq) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *GetProfileReq) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileReq) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) DeleteProfile(context.Context, *DeleteProfileReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedUserServiceServer) CreateDeliveryAddress(context.Context, *CreateDeliveryAddressReq) (*DeliveryAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeliveryAddress not implemented")
}
func (UnimplementedUserServiceServer) GetDeliveryAddress(context.Context, *GetDeliveryAddressReq) (*DeliveryAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeliveryAddress not implemented")
}
func (UnimplementedUserServiceServer) ListDeliveryAddress(context.Context, *ListDeliveryAddressReq) (*ListDeliveryAddressRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeliveryAddress not implemented")
}
func (UnimplementedUserServiceServer) UpdateDeliveryAddress(context.Context, *DeliveryAddress) (*DeliveryAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveryAddress not implemented")
}
func (UnimplementedUserServiceServer) DeleteDeliveryAddress(context.Context, *DeleteDeliveryAddressReq) (*EmptyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeliveryAddress not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateProfile(ctx, req.(*CreateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*GetProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteProfile(ctx, req.(*DeleteProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeliveryAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateDeliveryAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateDeliveryAddress(ctx, req.(*CreateDeliveryAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeliveryAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetDeliveryAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetDeliveryAddress(ctx, req.(*GetDeliveryAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeliveryAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListDeliveryAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListDeliveryAddress(ctx, req.(*ListDeliveryAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateDeliveryAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateDeliveryAddress(ctx, req.(*DeliveryAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteDeliveryAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeliveryAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteDeliveryAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteDeliveryAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteDeliveryAddress(ctx, req.(*DeleteDeliveryAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _UserService_CreateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _UserService_DeleteProfile_Handler,
		},
		{
			MethodName: "CreateDeliveryAddress",
			Handler:    _UserService_CreateDeliveryAddress_Handler,
		},
		{
			MethodName: "GetDeliveryAddress",
			Handler:    _UserService_GetDeliveryAddress_Handler,
		},
		{
			MethodName: "ListDeliveryAddress",
			Handler:    _UserService_ListDeliveryAddress_Handler,
		},
		{
			MethodName: "UpdateDeliveryAddress",
			Handler:    _UserService_UpdateDeliveryAddress_Handler,
		},
		{
			MethodName: "DeleteDeliveryAddress",
			Handler:    _UserService_DeleteDeliveryAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
