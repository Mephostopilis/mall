package service

import (
	"context"

	pb "edu/api/oms"
)

type ApiService struct {
	pb.UnimplementedApiServer
}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (s *ApiService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	return &pb.ListOrderReply{}, nil
}
func (s *ApiService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	return &pb.GetOrderReply{}, nil
}
func (s *ApiService) CreateOrder(ctx context.Context, req *pb.Order) (*pb.CreateOrderReply, error) {
	return &pb.CreateOrderReply{}, nil
}
