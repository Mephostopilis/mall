package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/dict/type/list [get]
// @Security Bearer
func (s *AdminService) GetDictTypeList(ctx context.Context, req *pb.GetDictTypeListRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	return s.uc.GetDictTypePage(ctx, token, req)
}

// @Summary 通过字典id获取字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Param dictId path int true "字典类型编码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/dict/type/{dictId} [get]
// @Security Bearer
func (s *AdminService) GetDictType(ctx context.Context, req *pb.GetDictTypeRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.GetDictType(ctx, req)
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

func (s *AdminService) GetDictTypeOptionSelect(ctx context.Context, req *pb.GetDictTypeOptionSelectRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.GetDictTypeListOptionSelect(ctx, req)
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
		Count:   int32(len(result)),
		Data:    list,
	}
	return
}

// @Summary 添加字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body pb.DictType true "data"
// @Success 200 {object} pb.ApiReply "{"code": 200, "message": "添加成功"}"
// @Router /admin/v1/dict/type [post]
// @Security Bearer
func (s *AdminService) InsertDictType(ctx context.Context, req *pb.DictType) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.InsertDictType(ctx, token, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 修改字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body pb.DictType true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/dict/type [put]
// @Security Bearer
func (s *AdminService) UpdateDictType(ctx context.Context, req *pb.DictType) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateDictType(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictId path int true "dictId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/type/{dictId} [delete]
func (s *AdminService) DeleteDictType(ctx context.Context, req *pb.DeleteDictTypeRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteDictType(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
