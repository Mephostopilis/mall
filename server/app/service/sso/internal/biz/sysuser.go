package biz

import (
	"context"

	pb "edu/api/sso/v1"
	uuidpb "edu/api/uuid"
	"edu/service/sso/internal/model"
)

func (uc *GreeterUsecase) CreateSysUser(ctx context.Context, req *pb.SysUser) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	sysuser := model.SysUser{
		SysUserId: model.SysUserId{UserId: id.Id},
		SysUserB:  model.SysUserB{},
		LoginM:    model.LoginM{},
	}
	_, err = uc.repo.InsertSysUser(&sysuser)
	if err != nil {
		return
	}
	return
}
