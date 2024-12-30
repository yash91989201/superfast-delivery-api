// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: shop.proto

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
	ShopService_CreateShop_FullMethodName        = "/pb.ShopService/CreateShop"
	ShopService_UpdateShop_FullMethodName        = "/pb.ShopService/UpdateShop"
	ShopService_UpdateShopAddress_FullMethodName = "/pb.ShopService/UpdateShopAddress"
	ShopService_UpdateShopContact_FullMethodName = "/pb.ShopService/UpdateShopContact"
	ShopService_UpdateShopImages_FullMethodName  = "/pb.ShopService/UpdateShopImages"
	ShopService_UpdateShopTimings_FullMethodName = "/pb.ShopService/UpdateShopTimings"
	ShopService_GetShop_FullMethodName           = "/pb.ShopService/GetShop"
	ShopService_ListShops_FullMethodName         = "/pb.ShopService/ListShops"
	ShopService_DeleteShop_FullMethodName        = "/pb.ShopService/DeleteShop"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	// Create a new shop
	CreateShop(ctx context.Context, in *CreateShopReq, opts ...grpc.CallOption) (*CreateShopRes, error)
	// Update shop details
	UpdateShop(ctx context.Context, in *UpdateShopReq, opts ...grpc.CallOption) (*UpdateShopRes, error)
	// Update shop address
	UpdateShopAddress(ctx context.Context, in *UpdateShopAddressReq, opts ...grpc.CallOption) (*UpdateShopAddressRes, error)
	// Update shop contact
	UpdateShopContact(ctx context.Context, in *UpdateShopContactReq, opts ...grpc.CallOption) (*UpdateShopContactRes, error)
	// Update shop images
	UpdateShopImages(ctx context.Context, in *UpdateShopImagesReq, opts ...grpc.CallOption) (*UpdateShopImagesRes, error)
	// Update shop timings
	UpdateShopTimings(ctx context.Context, in *UpdateShopTimingsReq, opts ...grpc.CallOption) (*UpdateShopTimingsRes, error)
	// Get shop details by ID
	GetShop(ctx context.Context, in *GetShopReq, opts ...grpc.CallOption) (*GetShopRes, error)
	// List all shops with optional filters
	ListShops(ctx context.Context, in *ListShopsReq, opts ...grpc.CallOption) (*ListShopsRes, error)
	// Delete a shop by ID (soft delete)
	DeleteShop(ctx context.Context, in *DeleteShopReq, opts ...grpc.CallOption) (*DeleteShopRes, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) CreateShop(ctx context.Context, in *CreateShopReq, opts ...grpc.CallOption) (*CreateShopRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShopRes)
	err := c.cc.Invoke(ctx, ShopService_CreateShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShop(ctx context.Context, in *UpdateShopReq, opts ...grpc.CallOption) (*UpdateShopRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShopRes)
	err := c.cc.Invoke(ctx, ShopService_UpdateShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShopAddress(ctx context.Context, in *UpdateShopAddressReq, opts ...grpc.CallOption) (*UpdateShopAddressRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShopAddressRes)
	err := c.cc.Invoke(ctx, ShopService_UpdateShopAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShopContact(ctx context.Context, in *UpdateShopContactReq, opts ...grpc.CallOption) (*UpdateShopContactRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShopContactRes)
	err := c.cc.Invoke(ctx, ShopService_UpdateShopContact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShopImages(ctx context.Context, in *UpdateShopImagesReq, opts ...grpc.CallOption) (*UpdateShopImagesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShopImagesRes)
	err := c.cc.Invoke(ctx, ShopService_UpdateShopImages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) UpdateShopTimings(ctx context.Context, in *UpdateShopTimingsReq, opts ...grpc.CallOption) (*UpdateShopTimingsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShopTimingsRes)
	err := c.cc.Invoke(ctx, ShopService_UpdateShopTimings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *GetShopReq, opts ...grpc.CallOption) (*GetShopRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShopRes)
	err := c.cc.Invoke(ctx, ShopService_GetShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) ListShops(ctx context.Context, in *ListShopsReq, opts ...grpc.CallOption) (*ListShopsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListShopsRes)
	err := c.cc.Invoke(ctx, ShopService_ListShops_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteShop(ctx context.Context, in *DeleteShopReq, opts ...grpc.CallOption) (*DeleteShopRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteShopRes)
	err := c.cc.Invoke(ctx, ShopService_DeleteShop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations must embed UnimplementedShopServiceServer
// for forward compatibility.
type ShopServiceServer interface {
	// Create a new shop
	CreateShop(context.Context, *CreateShopReq) (*CreateShopRes, error)
	// Update shop details
	UpdateShop(context.Context, *UpdateShopReq) (*UpdateShopRes, error)
	// Update shop address
	UpdateShopAddress(context.Context, *UpdateShopAddressReq) (*UpdateShopAddressRes, error)
	// Update shop contact
	UpdateShopContact(context.Context, *UpdateShopContactReq) (*UpdateShopContactRes, error)
	// Update shop images
	UpdateShopImages(context.Context, *UpdateShopImagesReq) (*UpdateShopImagesRes, error)
	// Update shop timings
	UpdateShopTimings(context.Context, *UpdateShopTimingsReq) (*UpdateShopTimingsRes, error)
	// Get shop details by ID
	GetShop(context.Context, *GetShopReq) (*GetShopRes, error)
	// List all shops with optional filters
	ListShops(context.Context, *ListShopsReq) (*ListShopsRes, error)
	// Delete a shop by ID (soft delete)
	DeleteShop(context.Context, *DeleteShopReq) (*DeleteShopRes, error)
	mustEmbedUnimplementedShopServiceServer()
}

// UnimplementedShopServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShopServiceServer struct{}

func (UnimplementedShopServiceServer) CreateShop(context.Context, *CreateShopReq) (*CreateShopRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedShopServiceServer) UpdateShop(context.Context, *UpdateShopReq) (*UpdateShopRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShop not implemented")
}
func (UnimplementedShopServiceServer) UpdateShopAddress(context.Context, *UpdateShopAddressReq) (*UpdateShopAddressRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShopAddress not implemented")
}
func (UnimplementedShopServiceServer) UpdateShopContact(context.Context, *UpdateShopContactReq) (*UpdateShopContactRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShopContact not implemented")
}
func (UnimplementedShopServiceServer) UpdateShopImages(context.Context, *UpdateShopImagesReq) (*UpdateShopImagesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShopImages not implemented")
}
func (UnimplementedShopServiceServer) UpdateShopTimings(context.Context, *UpdateShopTimingsReq) (*UpdateShopTimingsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShopTimings not implemented")
}
func (UnimplementedShopServiceServer) GetShop(context.Context, *GetShopReq) (*GetShopRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (UnimplementedShopServiceServer) ListShops(context.Context, *ListShopsReq) (*ListShopsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShops not implemented")
}
func (UnimplementedShopServiceServer) DeleteShop(context.Context, *DeleteShopReq) (*DeleteShopRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShop not implemented")
}
func (UnimplementedShopServiceServer) mustEmbedUnimplementedShopServiceServer() {}
func (UnimplementedShopServiceServer) testEmbeddedByValue()                     {}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	// If the following call pancis, it indicates UnimplementedShopServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_CreateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CreateShop(ctx, req.(*CreateShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShop(ctx, req.(*UpdateShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShopAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShopAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShopAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShopAddress(ctx, req.(*UpdateShopAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShopContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShopContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShopContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShopContact(ctx, req.(*UpdateShopContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShopImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopImagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShopImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShopImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShopImages(ctx, req.(*UpdateShopImagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_UpdateShopTimings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShopTimingsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).UpdateShopTimings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_UpdateShopTimings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).UpdateShopTimings(ctx, req.(*UpdateShopTimingsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*GetShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_ListShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShopsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).ListShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_ListShops_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).ListShops(ctx, req.(*ListShopsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_DeleteShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteShop(ctx, req.(*DeleteShopReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _ShopService_CreateShop_Handler,
		},
		{
			MethodName: "UpdateShop",
			Handler:    _ShopService_UpdateShop_Handler,
		},
		{
			MethodName: "UpdateShopAddress",
			Handler:    _ShopService_UpdateShopAddress_Handler,
		},
		{
			MethodName: "UpdateShopContact",
			Handler:    _ShopService_UpdateShopContact_Handler,
		},
		{
			MethodName: "UpdateShopImages",
			Handler:    _ShopService_UpdateShopImages_Handler,
		},
		{
			MethodName: "UpdateShopTimings",
			Handler:    _ShopService_UpdateShopTimings_Handler,
		},
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "ListShops",
			Handler:    _ShopService_ListShops_Handler,
		},
		{
			MethodName: "DeleteShop",
			Handler:    _ShopService_DeleteShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
