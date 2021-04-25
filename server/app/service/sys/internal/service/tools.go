package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 分页列表数据
// @Description 生成表分页列表
// @Tags 工具 - 生成表
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/tools/v1/sys/tables/page [get]
func (s *AdminService) GetSysTableList(ctx context.Context, req *pb.GetSysTableListRequest) (reply *pb.ApiReply, err error) {
	tables, count, err := s.uc.GetSysTableList(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(tables); i++ {
		it := tables[i]
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

// @Summary 获取配置
// @Description 获取JSON
// @Tags 工具 - 生成表
// @Param configKey path int true "configKey"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/tools/v1/sys/tables/info/{tableId} [get]
// @Security Bearer
func (s *AdminService) GetSysTables(ctx context.Context, req *pb.GetSysTablesRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.GetSysTables(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
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

// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 - 生成表
// @Accept  application/json
// @Product application/json
// @Param tables query string false "tableName / 数据表名称"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/tools/v1/sys/tables/info [post]
// @Security Bearer
func (s *AdminService) InsertSysTable(ctx context.Context, req *pb.InsertSysTableRequest) (reply *pb.ApiReply, err error) {
	tok, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.InsertSysTable(ctx, tok, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Reason:  "OK",
		Message: "OK",
	}
	return
}

// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 - 生成表
// @Accept  application/json
// @Product application/json
// @Param data body pb.SysTable true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/tools/v1/sys/tables/info [put]
// @Security Bearer
func (s *AdminService) UpdateSysTable(ctx context.Context, req *pb.SysTable) (reply *pb.ApiReply, err error) {
	tok, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.UpdateSysTable(ctx, tok, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 删除表结构
// @Description 删除表结构
// @Tags 工具 - 生成表
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/tools/v1/sys/tables/info/{tableId} [delete]
func (s *AdminService) DeleteSysTables(ctx context.Context, req *pb.DeleteSysTablesRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteSysTables(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}
