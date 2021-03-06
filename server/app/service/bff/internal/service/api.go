package service

import (
	"context"

	pb "edu/api/cms/v1"
	"edu/service/cms/internal/biz"
)

type ApiService struct {
	pb.UnimplementedApiServer
	uc *biz.GreeterUsecase
}

func NewCmsService(uc *biz.GreeterUsecase) *ApiService {
	return &ApiService{}
}

func (s *ApiService) ListHelp(ctx context.Context, req *pb.ListHelpRequest) (*pb.ListHelpReply, error) {
	return &pb.ListHelpReply{}, nil
}
func (s *ApiService) GetHelp(ctx context.Context, req *pb.GetHelpRequest) (*pb.GetHelpReply, error) {
	return &pb.GetHelpReply{}, nil
}
