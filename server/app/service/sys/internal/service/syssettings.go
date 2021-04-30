package service

import (
	"context"

	pb "edu/api/sys/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// GetSetting 获取设置
// @Summary 查询系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/setting [get]
func (s *AdminService) GetSetting(ctx context.Context, req *pb.GetSettingRequest) (reply *pb.ApiReply, err error) {
	set, err := s.uc.GetSetting(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(set)
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

// @Summary 更新或提交系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Param data body pb.CreateSettingRequest true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/system/setting [post]
func (s *AdminService) CreateSetting(ctx context.Context, req *pb.SysSetting) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateSetting(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}
