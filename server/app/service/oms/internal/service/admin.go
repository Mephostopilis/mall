package service

import (
	"context"

	pb "edu/api/oms"
	"edu/service/oms/internal/biz"
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

func (s *AdminService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateOrder(ctx context.Context, req *pb.Order) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) ListCompanyAddress(ctx context.Context, req *pb.ListCompanyAddressRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) GetCompanyAddress(ctx context.Context, req *pb.GetCompanyAddressRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) CreateCompanyAddress(ctx context.Context, req *pb.CompanyAddress) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) UpdateCompanyAddress(ctx context.Context, req *pb.CompanyAddress) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *AdminService) DeleteCompanyAddress(ctx context.Context, req *pb.DeleteCompanyAddressRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
