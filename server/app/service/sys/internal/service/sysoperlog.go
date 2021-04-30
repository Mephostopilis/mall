package service

import (
	"context"

	pb "edu/api/sys/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 登录日志列表
// @Description 获取JSON
// @Tags 登录日志
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/operloglist [get]
// @Security Bearer
func (s *AdminService) ListOperLog(ctx context.Context, req *pb.ListOperLogRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListOperLogPage(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		any, err1 := ptypes.MarshalAny(it)
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

// @Summary 通过编码获取登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Param infoId path int true "infoId"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/operlog/{infoId} [get]
// @Security Bearer
func (s *AdminService) GetOperLog(ctx context.Context, req *pb.GetOperLogRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetOperLog(ctx, req)
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

// @Summary 添加操作日志
// @Description 获取JSON
// @Tags 操作日志
// @Accept  application/json
// @Product application/json
// @Param data body models.SysOperLog true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/operlog [post]
// @Security Bearer
func (s *AdminService) CreateOperLog(ctx context.Context, req *pb.OperLog) (reply *pb.ApiReply, err error) {
	err = s.uc.CreateOperLog(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

func (s *AdminService) UpdateOperLog(context.Context, *pb.OperLog) (reply *pb.ApiReply, err error) {
	return
}

// @Summary 批量删除操作日志
// @Description 删除数据
// @Tags 操作日志
// @Param operId path string true "以逗号（,）分割的operId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/operlog/{operId} [delete]
func (s *AdminService) DeleteOperLog(ctx context.Context, req *pb.DeleteOperLogRequest) (reply *pb.ApiReply, err error) {
	// var data models.SysOperLog
	// data.UpdateBy = jwt.GetUserIdStr(ctx)
	// IDS := jwt.IdsStrToIdsIntGroup("operId", c)
	// result, err := s.adao.BatchDeleteSysOperLog(&data, IDS)
	// if err != nil {
	// 	return
	// }
	// reply = &pb.ApiReply{
	// 	Code:    0,
	// 	Message: "OK",
	// }
	return
}
