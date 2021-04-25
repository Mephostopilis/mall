package service

import (
	"context"

	pb "edu/api/cms"
	"edu/service/cms/internal/biz"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc *biz.GreeterUsecase
}

func NewAdminService(uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{}
}

func (s *AdminService) ListHelp(ctx context.Context, req *pb.ListHelpRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{Code: 0}, nil
}
func (s *AdminService) GetHelp(ctx context.Context, req *pb.GetHelpRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateHelp(ctx context.Context, req *pb.Help) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) UpdateHelp(context.Context, *pb.Help) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) DeleteHelp(context.Context, *pb.DeleteHelpRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) ListHelpCategory(context.Context, *pb.ListHelpCategoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) GetHelpCategory(context.Context, *pb.GetHelpCategoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) CreateHelpCategory(context.Context, *pb.HelpCategory) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) UpdateHelpCategory(ctx context.Context, req *pb.HelpCategory) (reply *pb.ApiReply, err error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) DeleteHelpCategory(context.Context, *pb.DeleteHelpCategoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
