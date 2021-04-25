package service

import (
	"context"

	pb "edu/api/sms"
	"edu/service/sms/internal/biz"
)

type ApiService struct {
	pb.UnimplementedApiServer
}

func NewApiService(uc *biz.GreeterUsecase) *ApiService {
	return &ApiService{}
}

func (s *ApiService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
func (s *ApiService) ListCoupon(ctx context.Context, req *pb.ListCouponRequest) (*pb.ListCouponReply, error) {
	return &pb.ListCouponReply{}, nil
}
func (s *ApiService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (*pb.GetCouponReply, error) {
	return &pb.GetCouponReply{}, nil
}
