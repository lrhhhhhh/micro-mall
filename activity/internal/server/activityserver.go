// Code generated by goctl. DO NOT EDIT!
// Source: activity.proto

package server

import (
	"context"

	"activity/internal/logic"
	"activity/internal/svc"
	"activity/service/activity"
)

type ActivityServer struct {
	svcCtx *svc.ServiceContext
	activity.UnimplementedActivityServer
}

func NewActivityServer(svcCtx *svc.ServiceContext) *ActivityServer {
	return &ActivityServer{
		svcCtx: svcCtx,
	}
}

func (s *ActivityServer) CreateActivity(ctx context.Context, in *activity.ActivityReq) (*activity.CreateActivityResp, error) {
	l := logic.NewCreateActivityLogic(ctx, s.svcCtx)
	return l.CreateActivity(in)
}

func (s *ActivityServer) UpdateActivity(ctx context.Context, in *activity.ActivityReq) (*activity.BaseActivityResp, error) {
	l := logic.NewUpdateActivityLogic(ctx, s.svcCtx)
	return l.UpdateActivity(in)
}

func (s *ActivityServer) GetActivity(ctx context.Context, in *activity.ActivityReq) (*activity.ActivityResp, error) {
	l := logic.NewGetActivityLogic(ctx, s.svcCtx)
	return l.GetActivity(in)
}

func (s *ActivityServer) DeleteActivity(ctx context.Context, in *activity.ActivityReq) (*activity.BaseActivityResp, error) {
	l := logic.NewDeleteActivityLogic(ctx, s.svcCtx)
	return l.DeleteActivity(in)
}

func (s *ActivityServer) GetActivityList(ctx context.Context, in *activity.ActivityListReq) (*activity.ActivityListResp, error) {
	l := logic.NewGetActivityListLogic(ctx, s.svcCtx)
	return l.GetActivityList(in)
}
