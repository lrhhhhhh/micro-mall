// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: service/stock/stock.proto

package stock

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

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockClient interface {
	DeductStock(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error)
	DeductStockRevert(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error)
	DeductStockFast(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error)
	DeductStockFastRevert(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) DeductStock(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error) {
	out := new(DeductStockResp)
	err := c.cc.Invoke(ctx, "/stock.stock/deductStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) DeductStockRevert(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error) {
	out := new(DeductStockResp)
	err := c.cc.Invoke(ctx, "/stock.stock/deductStockRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) DeductStockFast(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error) {
	out := new(DeductStockResp)
	err := c.cc.Invoke(ctx, "/stock.stock/deductStockFast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) DeductStockFastRevert(ctx context.Context, in *DeductStockReq, opts ...grpc.CallOption) (*DeductStockResp, error) {
	out := new(DeductStockResp)
	err := c.cc.Invoke(ctx, "/stock.stock/deductStockFastRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
// All implementations must embed UnimplementedStockServer
// for forward compatibility
type StockServer interface {
	DeductStock(context.Context, *DeductStockReq) (*DeductStockResp, error)
	DeductStockRevert(context.Context, *DeductStockReq) (*DeductStockResp, error)
	DeductStockFast(context.Context, *DeductStockReq) (*DeductStockResp, error)
	DeductStockFastRevert(context.Context, *DeductStockReq) (*DeductStockResp, error)
	mustEmbedUnimplementedStockServer()
}

// UnimplementedStockServer must be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (UnimplementedStockServer) DeductStock(context.Context, *DeductStockReq) (*DeductStockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStock not implemented")
}
func (UnimplementedStockServer) DeductStockRevert(context.Context, *DeductStockReq) (*DeductStockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStockRevert not implemented")
}
func (UnimplementedStockServer) DeductStockFast(context.Context, *DeductStockReq) (*DeductStockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStockFast not implemented")
}
func (UnimplementedStockServer) DeductStockFastRevert(context.Context, *DeductStockReq) (*DeductStockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStockFastRevert not implemented")
}
func (UnimplementedStockServer) mustEmbedUnimplementedStockServer() {}

// UnsafeStockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServer will
// result in compilation errors.
type UnsafeStockServer interface {
	mustEmbedUnimplementedStockServer()
}

func RegisterStockServer(s grpc.ServiceRegistrar, srv StockServer) {
	s.RegisterService(&Stock_ServiceDesc, srv)
}

func _Stock_DeductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/deductStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductStock(ctx, req.(*DeductStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_DeductStockRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductStockRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/deductStockRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductStockRevert(ctx, req.(*DeductStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_DeductStockFast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductStockFast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/deductStockFast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductStockFast(ctx, req.(*DeductStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_DeductStockFastRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductStockFastRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock.stock/deductStockFastRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductStockFastRevert(ctx, req.(*DeductStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Stock_ServiceDesc is the grpc.ServiceDesc for Stock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock.stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "deductStock",
			Handler:    _Stock_DeductStock_Handler,
		},
		{
			MethodName: "deductStockRevert",
			Handler:    _Stock_DeductStockRevert_Handler,
		},
		{
			MethodName: "deductStockFast",
			Handler:    _Stock_DeductStockFast_Handler,
		},
		{
			MethodName: "deductStockFastRevert",
			Handler:    _Stock_DeductStockFastRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/stock/stock.proto",
}
