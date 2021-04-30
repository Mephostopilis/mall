package biz

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/pkg/strings"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
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

func (uc *AdminUsecase) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (reply *pb.SysConfig, err error) {
	var filter model.SysConfig
	filter.ConfigId = (req.ConfigId)
	it, err := uc.d.GetSysConfig(&filter)
	if err != nil {
		return
	}
	reply = &pb.SysConfig{
		ConfigId:    int32(it.ConfigId),
		ConfigName:  it.ConfigName,
		ConfigKey:   it.ConfigKey,
		ConfigValue: it.ConfigValue,
		ConfigType:  it.ConfigType,
		Remark:      it.Remark,
		CreatedAt:   timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *AdminUsecase) GetConfigByConfigKey(ctx context.Context, req *pb.GetConfigByConfigKeyRequest) (reply *pb.SysConfig, err error) {
	var filter model.SysConfig
	filter.ConfigKey = req.ConfigKey
	it, err := uc.d.GetSysConfig(&filter)
	if err != nil {
		return
	}
	reply = &pb.SysConfig{
		ConfigId:    int32(it.ConfigId),
		ConfigName:  it.ConfigName,
		ConfigKey:   it.ConfigKey,
		ConfigValue: it.ConfigValue,
		ConfigType:  it.ConfigType,
		Remark:      it.Remark,
		CreatedAt:   timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *AdminUsecase) CreateConfig(ctx context.Context, req *pb.SysConfig) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("permision")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	var permission ssopb.DataPermission
	if err = proto.Unmarshal([]byte(v[0]), &permission); err != nil {
		return
	}
	data := model.SysConfig{
		ConfigId:    req.ConfigId,
		ConfigName:  req.ConfigName,
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		ConfigType:  req.ConfigType,
		Remark:      req.Remark,
	}
	data.CreateBy = fmt.Sprint(permission.UserId)
	_, err = uc.d.CreateSysConfig(ctx, &data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateConfig(ctx context.Context, req *pb.SysConfig) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("permision")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	var permission ssopb.DataPermission
	if err = proto.Unmarshal([]byte(v[0]), &permission); err != nil {
		return
	}
	data := model.SysConfig{
		ConfigId:    req.ConfigId,
		ConfigName:  req.ConfigName,
		ConfigKey:   req.ConfigKey,
		ConfigValue: req.ConfigValue,
		ConfigType:  req.ConfigType,
		Remark:      req.Remark,
	}
	data.UpdateBy = fmt.Sprint(permission.UserId)
	_, err = uc.d.UpdateSysConfig(&data, int(data.ConfigId))
	if err != nil {
		return
	}
	return
}

// @Summary 删除配置
// @Description 删除数据
// @Tags 配置
// @Param configId path int true "configId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/config/{configId} [delete]
func (uc *AdminUsecase) DeleteConfig(ctx context.Context, req *pb.DeleteConfigRequest) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("permision")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	var permission ssopb.DataPermission
	if err = proto.Unmarshal([]byte(v[0]), &permission); err != nil {
		return
	}
	var data model.SysConfig
	data.UpdateBy = fmt.Sprint(permission.UserId)
	ids, err := strings.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteSysConfig(&data, ids)
	if err != nil {
		return
	}
	return
}
