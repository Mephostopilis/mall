package service

import (
	"context"

	pb "edu/api/sso/v1"
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

// Authorize 应用授权
func (s *ApiService) Authorize(ctx context.Context, req *pb.AuthorizeReq) (reply *pb.AuthorizeResp, err error) {
	reply, err = s.uc.HandleAuthorizeRequest(ctx, req)
	return
}

// Token bm
func (s *ApiService) Token(ctx context.Context, req *pb.TokenReq) (reply *pb.TokenResp, err error) {
	reply, err = s.uc.HandleTokenRequest(ctx, req)
	return
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
