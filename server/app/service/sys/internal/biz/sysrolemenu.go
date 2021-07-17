package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/service/sys/internal/model"
)

func (uc *AdminUsecase) DeleteRoleMenu(ctx context.Context, token string, req *pb.DeleteRoleMenuRequest) (reply *pb.ApiReply, err error) {
	dp, err := meta.GetDataPermissions(ctx)
	if err != nil {
		return
	}
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
