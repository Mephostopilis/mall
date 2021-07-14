package service

import (
	"context"

	pb "edu/api/uuid"
	"edu/service/uuid/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UUIDService struct {
	pb.UnimplementedUUIDServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewUUIDService(logger log.Logger, uc *biz.GreeterUsecase) *UUIDService {
	return &UUIDService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service")),
	}
}

func (s *UUIDService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}

func (s *UUIDService) GenID(ctx context.Context, req *pb.GenIDReq) (reply *pb.GenIDResp, err error) {
	id, err := s.uc.GenID(ctx)
	if err != nil {
		return
	}
	reply = &pb.GenIDResp{
		Id: id,
	}
	return
}
