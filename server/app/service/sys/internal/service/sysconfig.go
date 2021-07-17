package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 配置列表数据
// @Description 获取JSON
// @Tags 配置
// @Param configKey query string false "configKey"
// @Param configName query string false "configName"
// @Param configType query string false "configType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/configList [get]
// @Security Bearer
func (s *AdminService) ListConfig(ctx context.Context, req *pb.ListConfigRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	return s.uc.ListConfig(ctx, token, req)
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 配置
// @Param configId path int true "配置编码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/config/{configId} [get]
// @Security Bearer
func (s *AdminService) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetConfig(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
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

// @Summary 获取配置
// @Description 获取JSON
// @Tags 配置
// @Param configKey path int true "configKey"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/configKey/{configKey} [get]
// @Security Bearer
func (s *AdminService) GetConfigByConfigKey(ctx context.Context, req *pb.GetConfigByConfigKeyRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetConfigByConfigKey(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
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

// @Summary 添加配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body models.SysConfig true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (s *AdminService) CreateConfig(ctx context.Context, req *pb.SysConfig) (reply *pb.ApiReply, err error) {
	err = s.uc.CreateConfig(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 修改配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body models.SysConfig true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/config [put]
// @Security Bearer
func (s *AdminService) UpdateConfig(ctx context.Context, req *pb.SysConfig) (reply *pb.ApiReply, err error) {
	err = s.uc.UpdateConfig(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
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
func (s *AdminService) DeleteConfig(ctx context.Context, req *pb.DeleteConfigRequest) (reply *pb.ApiReply, err error) {
	err = s.uc.DeleteConfig(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
