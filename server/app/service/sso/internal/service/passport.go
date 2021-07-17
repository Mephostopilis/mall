package service

import (
	"context"

	pb "edu/api/sso/v1"
	"edu/pkg/captcha"
	"edu/pkg/meta"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ApiService) GenerateCaptchaHandler(ctx context.Context, req *pb.GenerateCaptchaHandlerRequest) (reply *pb.GenerateCaptchaHandlerReply, err error) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		return
	}
	reply = &pb.GenerateCaptchaHandlerReply{
		Id:   id,
		B64S: b64s,
	}
	return
}

// SMSCode 获取短信验证吗
func (s *ApiService) SMSCode(ctx context.Context, req *pb.SMSCodeReq) (reply *pb.SMSCodeResp, err error) {
	return
}

//Register 只针对本应用
func (s *ApiService) Register(ctx context.Context, req *pb.RegisterReq) (reply *pb.RegisterResp, err error) {
	id, err := s.uc.GenID(ctx)
	if err != nil {
		return
	}
	// _, err = s.dao.cre(ctx, &model.SysUser{
	// 	Username:     req.Username,
	// 	Password:     req.Password,
	// 	Uid:          fmt.Sprintf("%d", id),
	// 	Email:        req.Mail,
	// 	Mobile:       req.Mobile,
	// 	Birthday:     time.Now(),
	// 	LastSignInAt: time.Now(),
	// 	Privilege:    "0",
	// 	CreateBy:     "0",
	// 	UpdateBy:     "0",
	// })
	if err != nil {
		return
	}
	reply = &pb.RegisterResp{
		Mid:    uint64(id),
		Mail:   req.Mail,
		Mobile: req.Mobile,
	}
	return
}

// Login for
func (s *ApiService) Login(ctx context.Context, req *pb.LoginRequest) (reply *pb.LoginReply, err error) {
	token, expire, err := s.jwtuc.LoginHandler(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.LoginReply{
		Token:  token,
		Expire: timestamppb.New(expire),
	}
	return
}

// Logout for
func (s *ApiService) Logout(ctx context.Context, req *pb.LogoutReq) (reply *pb.LogoutResp, err error) {
	// resp, err := s.ssoClient.Logout(ctx, &ssopb.LogoutReq{})
	// if err != nil {
	// 	return
	// }
	// reply = &pb.LogoutResp{
	// 	Content: "OK",
	// }
	return
}

func (s *ApiService) LoginGoogle(ctx context.Context, req *pb.LoginGoogleReq) (resp *pb.LoginGoogleResp, err error) {
	return
}

func (s *ApiService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (reply *pb.RefreshTokenReply, err error) {
	tokenstr, err := meta.GetToken(ctx)
	token, expire, err := s.jwtuc.RefreshHandler(ctx, tokenstr)
	if err != nil {
		return
	}
	reply = &pb.RefreshTokenReply{
		Token:  token,
		Expire: timestamppb.New(expire),
	}
	return
}
