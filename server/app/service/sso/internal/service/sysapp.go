package service

import (
	"context"

	pb "edu/api/sso"
)

func (s *AdminService) GetSsoApp(ctx context.Context, req *pb.GetSsoAppRequest) (reply *pb.ApiReply, err error) {
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) GetSsoAppList(ctx context.Context, req *pb.GetSsoAppListRequest) (reply *pb.ApiReply, err error) {
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) InsertSsoApp(ctx context.Context, req *pb.App) (reply *pb.ApiReply, err error) {
	return
}

func (s *AdminService) UpdateSsoApp(ctx context.Context, req *pb.App) (reply *pb.ApiReply, err error) {
	return
}

func (s *AdminService) DeleteSsoApp(ctx context.Context, req *pb.DeleteSsoAppRequest) (reply *pb.ApiReply, err error) {
	return
}
