package biz

import (
	"context"

	uuidpb "edu/api/uuid"
	"edu/service/oms/internal/dao"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type GreeterUsecase struct {
	d     dao.Dao
	uuidc uuidpb.UUIDClient
	log   *log.Helper
}

func NewGreeterUsecase(logger log.Logger, d dao.Dao, r *registry.Registry) (*GreeterUsecase, error) {
	cc, err := uuidpb.NewUUID(context.Background(), grpc.WithDiscovery(r))
	if err != nil {
		return nil, err
	}
	return &GreeterUsecase{
		d:     d,
		uuidc: cc,
		log:   log.NewHelper("usecase/greeter", logger),
	}, nil
}
