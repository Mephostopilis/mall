package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/biz"
)

type SysService struct {
	pb.UnimplementedSysServer
	uc *biz.AdminUsecase
}

func NewSysService(uc *biz.AdminUsecase) *SysService {
	return &SysService{
		uc: uc,
	}
}

func (s *SysService) GetDept(ctx context.Context, req *pb.GetDeptRequest) (*pb.Dept, error) {
	return s.uc.GetDept(ctx, req.DeptId)
}

func (s *SysService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	return &pb.Post{}, nil
}
