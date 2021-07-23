package service

import (
	"context"

	pb "edu/api/sso/v1"

	"edu/service/sso/internal/biz"
)

type SsoService struct {
	pb.UnimplementedSsoServer
	uc    *biz.GreeterUsecase
	jwtuc *biz.JWTUsecase
}

func NewSsoService(uc *biz.GreeterUsecase, jwtuc *biz.JWTUsecase) *SsoService {
	return &SsoService{
		uc:    uc,
		jwtuc: jwtuc,
	}
}

func (s *SsoService) Introspect(ctx context.Context, req *pb.IntrospectReq) (reply *pb.IntrospectResp, err error) {
	ti, err := s.jwtuc.ValidationMidToken(ctx, req.AccessToken)
	if err != nil {
		return
	}
	reply = &pb.IntrospectResp{
		DataScope: ti.DataScope,
		UserId:    ti.UserId,
		RoleId:    ti.RoleId,
		RoleKey:   ti.RoleKey,
		// DeptId:    ti.DeptId,
		// PostId:    ti.PostId,
	}
	return
}
