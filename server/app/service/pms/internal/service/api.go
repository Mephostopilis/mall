package service

import (
	"context"

	pb "edu/api/pms"
	"edu/service/pms/internal/biz"
)

type ApiService struct {
	pb.UnimplementedApiServer
	uc *biz.GreeterUsecase
}

func NewApiService(uc *biz.GreeterUsecase) *ApiService {
	return &ApiService{
		uc: uc,
	}
}

func (s *ApiService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
func (s *ApiService) ListAlbum(ctx context.Context, req *pb.ListAlbumRequest) (*pb.ListAlbumReply, error) {
	return &pb.ListAlbumReply{}, nil
}
func (s *ApiService) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.GetAlbumReply, error) {
	return &pb.GetAlbumReply{}, nil
}
