package biz

import (
	"context"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
)

func (uc *AdminUsecase) DeleteRoleMenu(ctx context.Context, token string, req *pb.DeleteRoleMenuRequest) (reply *pb.ApiReply, err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(ssopb.DataPermission)
	var t model.SysRoleMenu
	_, err = uc.d.DeleteRoleMenu(&t, int(dp.RoleId), int(req.MenuId))
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
