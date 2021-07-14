package service

import (
	"context"
	pb "edu/api/sso"

	"edu/service/sso/internal/biz"
	"edu/service/sso/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	adao dao.Dao
	log  *log.Helper
	uc   *biz.GreeterUsecase
}

func NewAuthService(logger log.Logger, d dao.Dao, uc *biz.GreeterUsecase) *AuthService {
	log := log.NewHelper(log.With(logger, "module", "service/auth"))
	s := &AuthService{
		adao: d,
		uc:   uc,
		log:  log,
	}
	return s
}

// Authorize 应用授权
func (s *AuthService) Authorize(ctx context.Context, req *pb.AuthorizeReq) (reply *pb.AuthorizeResp, err error) {
	reply, err = s.uc.HandleAuthorizeRequest(ctx, req)
	return
}

// Token bm
func (s *AuthService) Token(ctx context.Context, req *pb.TokenReq) (reply *pb.TokenResp, err error) {
	reply, err = s.uc.HandleTokenRequest(ctx, req)
	return
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (reply *pb.RefreshTokenResp, err error) {
	return s.uc.RefreshToken(ctx, req)
}
