package service

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
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
	return s.uc.GetDictTypeList(ctx, token, req)
}

// @Summary 通过字典id获取字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Param dictId path int true "字典类型编码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/dict/type/{dictId} [get]
// @Security Bearer
func (s *AdminService) GetDictType(ctx context.Context, req *pb.GetDictTypeRequest) (reply *pb.ApiReply, err error) {
	var DictType model.SysDictType
	DictType.DictName = req.DictName
	DictType.DictId = int(req.DictId)
	it, err := s.adao.GetDictType(&DictType)
	if err != nil {
		return
	}

	list := make([]*anypb.Any, 0)
	d := &pb.DictType{
		DictId:   int32(it.DictId),
		DictName: it.DictName,
		DictType: it.DictType,
		Status:   it.Status,
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

func (s *AdminService) GetDictTypeOptionSelect(ctx context.Context, req *pb.GetDictTypeOptionSelectRequest) (reply *pb.ApiReply, err error) {
	var DictType model.SysDictType
	DictType.DictName = req.DictName
	DictType.DictId = int(req.DictId)
	result, err := s.adao.GetDictTypeList(&DictType)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DictType{
			DictId:   int32(it.DictId),
			DictName: it.DictName,
			DictType: it.DictType,
			Status:   it.Status,
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
	return s.uc.InsertDictType(ctx, token, req)
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

	var data model.SysDictType
	data.UpdateBy = fmt.Sprint(permission.UserId)
	it, err := s.adao.UpdateDictType(&data, data.DictId)
	if err != nil {
		return
	}

	list := make([]*anypb.Any, 0)
	d := &pb.DictType{
		DictId:   int32(it.DictId),
		DictName: it.DictName,
		DictType: it.DictType,
		Status:   it.Status,
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

// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictId path int true "dictId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/type/{dictId} [delete]
func (s *AdminService) DeleteDictType(ctx context.Context, req *pb.DeleteDictTypeRequest) (reply *pb.ApiReply, err error) {
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
	var data model.SysDictType
	data.UpdateBy = permission.RoleKey
	// IDS := jwt.IdsStrToIdsIntGroup("dictId", ctx)
	IDS := make([]int, 0)
	_, err = s.adao.BatchDeleteDictType(&data, IDS)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
