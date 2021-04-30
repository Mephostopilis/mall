package biz

import (
	"context"
	"fmt"
	"strings"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
)

func (uc *AdminUsecase) GetSetting(ctx context.Context, req *pb.GetSettingRequest) (reply *pb.SysSetting, err error) {
	filter := model.SysSetting{}
	set, err := uc.d.GetSysSetting(ctx, &filter)
	if err != nil {
		return
	}
	reply = &pb.SysSetting{
		Name: set.Name,
		Logo: set.Logo,
	}
	return
}

func (uc *AdminUsecase) CreateSetting(ctx context.Context, req *pb.SysSetting) (err error) {
	var d model.SysSetting
	d.Logo = req.Logo
	d.Name = req.Name
	if d.Logo != "" {
		if !strings.HasPrefix(d.Logo, "http") {
			d.Logo = fmt.Sprintf("http://%s", d.Logo)
		}
	}
	_, e := uc.d.UpdateSysSetting(ctx, &d)
	if e != nil {
		return
	}
	return
}
