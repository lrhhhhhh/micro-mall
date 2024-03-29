// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: service/seckill/seckill.proto

package seckill

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

// SeckillClient is the client API for Seckill service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeckillClient interface {
	Seckill(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error)
	Seckill2(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error)
}

type seckillClient struct {
	cc grpc.ClientConnInterface
}

func NewSeckillClient(cc grpc.ClientConnInterface) SeckillClient {
	return &seckillClient{cc}
}

func (c *seckillClient) Seckill(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error) {
	out := new(BaseSeckillResp)
	err := c.cc.Invoke(ctx, "/seckill.Seckill/Seckill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seckillClient) Seckill2(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error) {
	out := new(BaseSeckillResp)
	err := c.cc.Invoke(ctx, "/seckill.Seckill/Seckill2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeckillServer is the server API for Seckill service.
// All implementations must embed UnimplementedSeckillServer
// for forward compatibility
type SeckillServer interface {
	Seckill(context.Context, *SeckillReq) (*BaseSeckillResp, error)
	Seckill2(context.Context, *SeckillReq) (*BaseSeckillResp, error)
	mustEmbedUnimplementedSeckillServer()
}

// UnimplementedSeckillServer must be embedded to have forward compatible implementations.
type UnimplementedSeckillServer struct {
}

func (UnimplementedSeckillServer) Seckill(context.Context, *SeckillReq) (*BaseSeckillResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seckill not implemented")
}
func (UnimplementedSeckillServer) Seckill2(context.Context, *SeckillReq) (*BaseSeckillResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seckill2 not implemented")
}
func (UnimplementedSeckillServer) mustEmbedUnimplementedSeckillServer() {}

// UnsafeSeckillServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeckillServer will
// result in compilation errors.
type UnsafeSeckillServer interface {
	mustEmbedUnimplementedSeckillServer()
}

func RegisterSeckillServer(s grpc.ServiceRegistrar, srv SeckillServer) {
	s.RegisterService(&Seckill_ServiceDesc, srv)
}

func _Seckill_Seckill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillServer).Seckill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seckill.Seckill/Seckill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillServer).Seckill(ctx, req.(*SeckillReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seckill_Seckill2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeckillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeckillServer).Seckill2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seckill.Seckill/Seckill2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeckillServer).Seckill2(ctx, req.(*SeckillReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Seckill_ServiceDesc is the grpc.ServiceDesc for Seckill service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seckill_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seckill.Seckill",
	HandlerType: (*SeckillServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Seckill",
			Handler:    _Seckill_Seckill_Handler,
		},
		{
			MethodName: "Seckill2",
			Handler:    _Seckill_Seckill2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/seckill/seckill.proto",
}
