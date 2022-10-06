// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: service/activity/activity.proto

package activity

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	CreateActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*CreateActivityResp, error)
	UpdateActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*BaseActivityResp, error)
	GetActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityResp, error)
	DeleteActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*BaseActivityResp, error)
	GetActivityList(ctx context.Context, in *ActivityListReq, opts ...grpc.CallOption) (*ActivityListResp, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) CreateActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*CreateActivityResp, error) {
	out := new(CreateActivityResp)
	err := c.cc.Invoke(ctx, "/activity.Activity/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UpdateActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*BaseActivityResp, error) {
	out := new(BaseActivityResp)
	err := c.cc.Invoke(ctx, "/activity.Activity/UpdateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*ActivityResp, error) {
	out := new(ActivityResp)
	err := c.cc.Invoke(ctx, "/activity.Activity/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteActivity(ctx context.Context, in *ActivityReq, opts ...grpc.CallOption) (*BaseActivityResp, error) {
	out := new(BaseActivityResp)
	err := c.cc.Invoke(ctx, "/activity.Activity/DeleteActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivityList(ctx context.Context, in *ActivityListReq, opts ...grpc.CallOption) (*ActivityListResp, error) {
	out := new(ActivityListResp)
	err := c.cc.Invoke(ctx, "/activity.Activity/GetActivityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	CreateActivity(context.Context, *ActivityReq) (*CreateActivityResp, error)
	UpdateActivity(context.Context, *ActivityReq) (*BaseActivityResp, error)
	GetActivity(context.Context, *ActivityReq) (*ActivityResp, error)
	DeleteActivity(context.Context, *ActivityReq) (*BaseActivityResp, error)
	GetActivityList(context.Context, *ActivityListReq) (*ActivityListResp, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) CreateActivity(context.Context, *ActivityReq) (*CreateActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServer) UpdateActivity(context.Context, *ActivityReq) (*BaseActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (UnimplementedActivityServer) GetActivity(context.Context, *ActivityReq) (*ActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServer) DeleteActivity(context.Context, *ActivityReq) (*BaseActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServer) GetActivityList(context.Context, *ActivityListReq) (*ActivityListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityList not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.Activity/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.Activity/UpdateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UpdateActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.Activity/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.Activity/DeleteActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteActivity(ctx, req.(*ActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.Activity/GetActivityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivityList(ctx, req.(*ActivityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _Activity_CreateActivity_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _Activity_UpdateActivity_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _Activity_GetActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _Activity_DeleteActivity_Handler,
		},
		{
			MethodName: "GetActivityList",
			Handler:    _Activity_GetActivityList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/activity/activity.proto",
}
