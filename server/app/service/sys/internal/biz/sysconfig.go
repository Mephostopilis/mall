package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) ListConfig(ctx context.Context, token string, req *pb.ListConfigRequest) (reply *pb.ApiReply, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var data model.SysConfig
	data.ConfigKey = req.ConfigKey
	data.ConfigName = req.ConfigName
	data.ConfigType = req.ConfigType
	result, count, err := uc.d.GetSysConfigPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.SysConfig{
			ConfigId:    int32(it.ConfigId),
			ConfigName:  it.ConfigName,
			ConfigKey:   it.ConfigKey,
			ConfigValue: it.ConfigValue,
			ConfigType:  it.ConfigType,
			Remark:      it.Remark,
			CreatedAt:   timestamppb.New(it.CreatedAt),
		}
		any, err1 := ptypes.MarshalAny(d)
		if err1 != nil {
			err = err1
			return
		}
		list = append(list, any)
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(count),
		Data:    list,
	}
	return
}
