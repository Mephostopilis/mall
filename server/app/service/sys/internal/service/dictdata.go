package service

import (
	"context"

	pb "edu/api/sys/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/dict/data/list [get]
// @Security Bearer
func (s *AdminService) GetDictDataList(ctx context.Context, req *pb.GetDictDataListRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.GetDictDataPage(ctx, req)
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
		Code:      0,
		Message:   "OK",
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		Count:     int32(count),
		Data:      list,
	}
	return
}

// @Summary 通过字典类型获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictType path int true "dictType"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/databyType/{dictType} [get]
// @Security Bearer
func (s *AdminService) GetDictDataListByDictType(ctx context.Context, req *pb.GetDictDataListByDictTypeRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.GetDictDataListByDictType(ctx, req)
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
		Count:   int32(len(list)),
		Data:    list,
	}
	return
}

// @Summary 通过编码获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictCode path int true "字典编码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security Bearer
func (s *AdminService) GetDictData(ctx context.Context, req *pb.GetDictDataRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.GetDictData(ctx, req)
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

// @Summary 添加字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body pb.DictData true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (s *AdminService) InsertDictData(ctx context.Context, req *pb.DictData) (reply *pb.ApiReply, err error) {
	if err = s.uc.InsertDictData(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body pb.DictData true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [put]
// @Security Bearer
func (s *AdminService) UpdateDictData(ctx context.Context, req *pb.DictData) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateDictData(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode path int true "dictCode"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/data/{dictCode} [delete]
func (s *AdminService) DeleteDictData(ctx context.Context, req *pb.DeleteDictDataRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteDictData(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
