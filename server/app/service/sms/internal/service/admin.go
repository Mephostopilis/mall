package service

import (
	"context"

	pb "edu/api/sms"
	"edu/service/sms/internal/biz"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
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

func (s *AdminService) ListCoupon(ctx context.Context, req *pb.ListCouponRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListCouponPage(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		any, err1 := ptypes.MarshalAny(it)
		if err1 != nil {
			err = err1
			return
		}
		list = append(list, any)
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(count),
		Data:    list,
	}
	return
}

func (s *AdminService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetCoupon(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
		Data:    list,
	}
	return
}

func (s *AdminService) CreateCoupon(ctx context.Context, req *pb.Coupon) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateCoupon(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateCoupon(ctx context.Context, req *pb.Coupon) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateCoupon(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteCoupon(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
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
