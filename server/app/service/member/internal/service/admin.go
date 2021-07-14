package service

import (
	"context"

	pb "edu/api/member"
	"edu/service/member/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.MemberUsecase
}

func NewAdminService(logger log.Logger, uc *biz.MemberUsecase) *AdminService {
	return &AdminService{
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
		uc:  uc,
	}
}

func (s *AdminService) ListMember(ctx context.Context, req *pb.ListMemberRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateMember(ctx context.Context, req *pb.Member) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateMember(ctx context.Context, req *pb.Member) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) ListMemberAssets(ctx context.Context, req *pb.ListMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetMemberAssets(ctx context.Context, req *pb.GetMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateMemberAssets(ctx context.Context, req *pb.MemberAssets) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateMemberAssets(ctx context.Context, req *pb.MemberAssets) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteMemberAssets(ctx context.Context, req *pb.DeleteMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
