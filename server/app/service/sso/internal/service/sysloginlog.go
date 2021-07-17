package service

import (
	"context"

	pb "edu/api/sso/v1"
	"edu/service/sso/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
// @Router /admin/v1/loginloglist [get]
// @Security Bearer
func (s *AdminService) ListLoginLog(ctx context.Context, req *pb.ListLoginLogRequest) (reply *pb.ApiReply, err error) {

	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysLoginLog
	data.Username = req.Username
	data.Status = req.Status
	data.Ipaddr = req.Ipaddr
	result, count, err := s.adao.GetLoginLogPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.LoginLog{
			InfoId:        int32(it.InfoId),
			Username:      it.Username,
			Status:        it.Status,
			Ipaddr:        it.Ipaddr,
			LoginLocation: it.LoginLocation,
			Browser:       it.Browser,
			Os:            it.Os,
			Platform:      it.Platform,
			LoginTime:     timestamppb.New(it.LoginTime),
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

// @Summary 通过编码获取登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Param infoId path int true "infoId"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/loginlog/{infoId} [get]
// @Security Bearer
func (s *AdminService) GetLoginLog(ctx context.Context, req *pb.GetLoginLogRequest) (reply *pb.ApiReply, err error) {
	var LoginLog model.SysLoginLog
	LoginLog.InfoId = int(req.InfoId)
	it, err := s.adao.GetLoginLog(&LoginLog)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.LoginLog{
		InfoId:        int32(it.InfoId),
		Username:      it.Username,
		Status:        it.Status,
		Ipaddr:        it.Ipaddr,
		LoginLocation: it.LoginLocation,
		Browser:       it.Browser,
		Os:            it.Os,
		Platform:      it.Platform,
		LoginTime:     timestamppb.New(it.LoginTime),
	}
	any, err1 := ptypes.MarshalAny(d)
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

// @Summary 添加登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Accept  application/json
// @Product application/json
// @Param data body pb.LoginLog true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/loginlog [post]
// @Security Bearer
func (s *AdminService) CreateLoginLog(ctx context.Context, req *pb.LoginLog) (reply *pb.ApiReply, err error) {
	var data model.SysLoginLog
	_, err = s.adao.CreateLoginLog(&data)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 修改登录日志
// @Description 获取JSON
// @Tags 登录日志
// @Accept  application/json
// @Product application/json
// @Param data body pb.LoginLog true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/loginlog [put]
// @Security Bearer
func (s *AdminService) UpdateLoginLog(ctx context.Context, req *pb.LoginLog) (reply *pb.ApiReply, err error) {
	var data model.SysLoginLog
	it, err := s.adao.UpdateLoginLog(&data, data.InfoId)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.LoginLog{
		InfoId:        int32(it.InfoId),
		Username:      it.Username,
		Status:        it.Status,
		Ipaddr:        it.Ipaddr,
		LoginLocation: it.LoginLocation,
		Browser:       it.Browser,
		Os:            it.Os,
		Platform:      it.Platform,
	}
	any, err1 := ptypes.MarshalAny(d)
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

// @Summary 批量删除登录日志
// @Description 删除数据
// @Tags 登录日志
// @Param infoId path string true "以逗号（,）分割的infoId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/loginlog/{infoId} [delete]
func (s *AdminService) DeleteLoginLog(c context.Context, req *pb.DeleteLoginLogRequest) (reply *pb.ApiReply, err error) {
	var data model.SysLoginLog
	data.UpdateBy = ""
	IDS := make([]int, 0)
	result, err := s.adao.BatchDeleteLoginLog(&data, IDS)
	if err != nil {
		return
	}
	if result {
		reply = &pb.ApiReply{
			Code:    0,
			Message: "OK",
		}
	}
	return
}
