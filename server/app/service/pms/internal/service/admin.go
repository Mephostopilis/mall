package service

import (
	"context"

	pb "edu/api/pms"
	"edu/service/pms/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.GreeterUsecase
}

func NewAdminService(logger log.Logger, uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{
		log: log.NewHelper("service/admin", logger),
		uc:  uc,
	}
}

func (s *AdminService) ListAlbum(ctx context.Context, req *pb.ListAlbumRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateAlbum(ctx context.Context, req *pb.Album) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateAlbum(ctx context.Context, req *pb.Album) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteAlbum(ctx context.Context, req *pb.DeleteAlbumRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) ListAlbumPic(ctx context.Context, req *pb.ListAlbumPicRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetAlbumPic(ctx context.Context, req *pb.GetAlbumPicRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateAlbumPic(ctx context.Context, req *pb.AlbumPic) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateAlbumPic(ctx context.Context, req *pb.AlbumPic) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteAlbumPic(ctx context.Context, req *pb.DeleteAlbumPicRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
