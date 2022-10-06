// Code generated by goctl. DO NOT EDIT!
// Source: seckill.proto

package seckillclient

import (
	"context"

	"seckill/service/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseSeckillResp = seckill.BaseSeckillResp
	SeckillReq      = seckill.SeckillReq

	Seckill interface {
		Seckill(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error)
		Seckill2(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error)
	}

	defaultSeckill struct {
		cli zrpc.Client
	}
)

func NewSeckill(cli zrpc.Client) Seckill {
	return &defaultSeckill{
		cli: cli,
	}
}

func (m *defaultSeckill) Seckill(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error) {
	client := seckill.NewSeckillClient(m.cli.Conn())
	return client.Seckill(ctx, in, opts...)
}

func (m *defaultSeckill) Seckill2(ctx context.Context, in *SeckillReq, opts ...grpc.CallOption) (*BaseSeckillResp, error) {
	client := seckill.NewSeckillClient(m.cli.Conn())
	return client.Seckill2(ctx, in, opts...)
}
