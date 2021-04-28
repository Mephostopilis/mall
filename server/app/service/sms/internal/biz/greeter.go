package biz

import (
	"context"

	pb "edu/api/sms"
	uuidpb "edu/api/uuid"
	"edu/pkg/strings"
	"edu/service/sms/internal/dao"
	"edu/service/sms/internal/model"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GreeterUsecase struct {
	uuidc uuidpb.UUIDClient
	d     dao.Dao
	log   *log.Helper
}

func NewGreeterUsecase(logger log.Logger, d dao.Dao, r *registry.Registry) (*GreeterUsecase, error) {
	cc, err := uuidpb.NewUUID(context.Background(), grpc.WithDiscovery(r))
	if err != nil {
		return nil, err
	}
	return &GreeterUsecase{
		uuidc: cc,
		d:     d,
		log:   log.NewHelper("usecase/greeter", logger),
	}, nil
}

func (uc *GreeterUsecase) ListCouponPage(ctx context.Context, req *pb.ListCouponRequest) (page []*pb.Coupon, total int64, err error) {
	filter := model.SmsCoupon{}
	list, total, err := uc.d.GetSmsCouponPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Coupon, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Coupon{
			Id:        it.Id,
			Name:      it.Name,
			Amount:    it.Amount,
			Count:     it.Count,
			Platform:  it.Platform,
			Type:      it.Type,
			CreatedAt: timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (reply *pb.Coupon, err error) {
	it, err := uc.d.GetSmsCoupon(ctx, &model.SmsCoupon{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.Coupon{
		Id:        it.Id,
		Name:      it.Name,
		Amount:    it.Amount,
		Count:     it.Count,
		Platform:  it.Platform,
		Type:      it.Type,
		CreatedAt: timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateCoupon(ctx context.Context, req *pb.Coupon) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.SmsCoupon{
		Id:       id.Id,
		Name:     req.Name,
		Amount:   req.Amount,
		Count:    req.Count,
		Platform: req.Platform,
		Type:     req.Type,
	}
	_, err = uc.d.CreateSmsCoupon(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateCoupon(ctx context.Context, req *pb.Coupon) (err error) {
	d := model.SmsCoupon{
		Id:       req.Id,
		Name:     req.Name,
		Amount:   req.Amount,
		Count:    req.Count,
		Platform: req.Platform,
		Type:     req.Type,
	}
	_, err = uc.d.UpdateSmsCoupon(ctx, &d, (req.Id))
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteSmsCoupon(ctx, &model.SmsCoupon{}, ids)
	if err != nil {
		return
	}
	return
}
