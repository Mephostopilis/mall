package service

import (
	"context"

	pb "edu/api/sms"
	"edu/service/sms/internal/biz"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc *biz.GreeterUsecase
}

func NewAdminService(uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{
		uc: uc,
	}
}

func (s *AdminService) ListCoupon(ctx context.Context, req *pb.ListCouponRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateCoupon(ctx context.Context, req *pb.Coupon) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateCoupon(ctx context.Context, req *pb.Coupon) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) ListCouponHistory(ctx context.Context, req *pb.ListCouponHistoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetCouponHistory(ctx context.Context, req *pb.GetCouponHistoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateCouponHistory(ctx context.Context, req *pb.CouponHistory) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateCouponHistory(ctx context.Context, req *pb.CouponHistory) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteCouponHistory(ctx context.Context, req *pb.DeleteCouponHistoryRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
