package service

import (
	"context"

	pb "edu/api/sso"
	"edu/service/sso/internal/biz"
	"edu/service/sso/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	pb.UnimplementedApiServer
	dao   dao.Dao
	uc    *biz.GreeterUsecase
	jwtuc *biz.JWTUsecase
	log   *log.Helper
}

func NewApiService(d dao.Dao, jwtuc *biz.JWTUsecase, uc *biz.GreeterUsecase, logger log.Logger) *ApiService {
	log := log.NewHelper(log.With(logger, "module", "service/api"))
	return &ApiService{
		dao:   d,
		jwtuc: jwtuc,
		uc:    uc,
		log:   log,
	}
}

func (s *ApiService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}

func (s *ApiService) InsetSysUserAvatar(ctx context.Context, req *pb.InsetSysUserAvatarRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *ApiService) SysUserUpdatePwd(ctx context.Context, req *pb.SysUserUpdatePwdRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *ApiService) GetSysUserProfile(ctx context.Context, req *pb.GetSysUserProfileRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
