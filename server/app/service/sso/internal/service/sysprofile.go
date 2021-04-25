package service

import (
	"context"

	pb "edu/api/sso"

	"edu/pkg/meta"
)

// @Summary 获取个人中心用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/user/profile [get]
// @Security Bearer
func (s *AdminService) GetSysUserProfile(ctx context.Context, req *pb.GetSysUserProfileRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	pd, err := s.jwtuc.ValidationMidToken(ctx, token)
	if err != nil {
		return
	}
	return s.uc.GetSysUserProfile(ctx, pd, req)
}

// @Summary 修改头像
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/user/profileAvatar [post]
func (s *AdminService) InsetSysUserAvatar(ctx context.Context, req *pb.InsetSysUserAvatarRequest) (reply *pb.ApiReply, err error) {
	// token, err := meta.GetToken(ctx)
	// if err != nil {
	// 	return
	// }
	// pd, err := s.jwtuc.ValidationMidToken(ctx, token)
	// if err != nil {
	// 	return
	// }
	return s.uc.InsetSysUserAvatar(ctx, req)
}

// @Summary 修改密码
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/user/profileAvatar [post]
func (s *AdminService) SysUserUpdatePwd(ctx context.Context, req *pb.SysUserUpdatePwdRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	pd, err := s.jwtuc.ValidationMidToken(ctx, token)
	if err != nil {
		return
	}
	err = s.uc.SetSysUserPwd(ctx, pd.UserId, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}
