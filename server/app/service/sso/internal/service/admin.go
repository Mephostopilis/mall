package service

import (
	"context"
	pb "edu/api/sso/v1"
	"edu/pkg/captcha"
	"edu/pkg/meta"
	"edu/service/sso/internal/biz"
	"edu/service/sso/internal/dao"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	adao  dao.Dao
	log   *log.Helper
	uc    *biz.GreeterUsecase
	jwtuc *biz.JWTUsecase
}

func NewAdminService(logger log.Logger, d dao.Dao, uc *biz.GreeterUsecase, jwtuc *biz.JWTUsecase) *AdminService {
	log := log.NewHelper(log.With(logger, "module", "service/admin"))
	s := &AdminService{
		adao:  d,
		uc:    uc,
		jwtuc: jwtuc,
		log:   log,
	}
	return s
}

// @Summary 生成验证码 HealthCheck shows OK as the ping-pong result.
// @Description 健康状况
// @Accept text/html
// @Produce text/html
// @Success 200 {string} string "OK"
// @Router /admin/public/health [get]
// @BasePath
func (s *AdminService) GenerateCaptchaHandler(ctx context.Context, req *pb.GenerateCaptchaHandlerRequest) (reply *pb.ApiReply, err error) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.GenerateCaptchaHandlerReply{
		Id:   id,
		B64S: b64s,
	}
	any, err1 := ptypes.MarshalAny(d)
	if err1 != nil {
		s.log.Errorf("err = %v", err1)
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

func (s *AdminService) Login(ctx context.Context, req *pb.LoginRequest) (reply *pb.ApiReply, err error) {
	token, expire, err := s.jwtuc.LoginHandler(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.LoginReply{
		Token:  token,
		Expire: timestamppb.New(expire),
	}
	any, err1 := ptypes.MarshalAny(d)
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

func (s *AdminService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	token, expire, err := s.jwtuc.RefreshHandler(ctx, token)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.RefreshTokenReply{
		Token:  token,
		Expire: timestamppb.New(expire),
	}
	any, err1 := ptypes.MarshalAny(d)
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

// GetInfo 得到用户初始信息
func (s *AdminService) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	pd, err := s.jwtuc.ValidationMidToken(ctx, token)
	if err != nil {
		return
	}
	d, err := s.uc.GetInfo(ctx, pd, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(d)
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
