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

func (uc *AdminUsecase) GetDictTypePage(ctx context.Context, token string, req *pb.GetDictTypeListRequest) (reply *pb.ApiReply, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysDictType
	data.DictId = req.DictId
	data.DictName = req.DictName
	data.DictType = req.DictType
	result, total, err := uc.d.GetDictTypePage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}

	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DictType{
			DictId:    int32(it.DictId),
			DictName:  it.DictName,
			DictType:  it.DictType,
			Status:    it.Status,
			Remark:    it.Remark,
			CreatedAt: timestamppb.New(it.CreatedAt),
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
		Count:   int32(total),
		Data:    list,
	}
	return
}

func (uc *AdminUsecase) GetDictTypeListOptionSelect(ctx context.Context, req *pb.GetDictTypeOptionSelectRequest) (list []*pb.DictType, err error) {
	var filter model.SysDictType
	filter.DictName = req.DictName
	filter.DictId = req.DictId
	result, err := uc.d.GetDictTypeList(&filter)
	if err != nil {
		return
	}
	list = make([]*pb.DictType, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DictType{
			DictId:   int32(it.DictId),
			DictName: it.DictName,
			DictType: it.DictType,
			Status:   it.Status,
		}
		list = append(list, d)
	}
	return
}

func (uc *AdminUsecase) GetDictType(ctx context.Context, req *pb.GetDictTypeRequest) (reply *pb.DictType, err error) {
	var filter model.SysDictType
	filter.DictName = req.DictName
	filter.DictId = req.DictId
	it, err := uc.d.GetDictType(&filter)
	if err != nil {
		return
	}
	reply = &pb.DictType{
		DictId:   int32(it.DictId),
		DictName: it.DictName,
		DictType: it.DictType,
		Status:   it.Status,
	}
	return
}

func (uc *AdminUsecase) InsertDictType(ctx context.Context, token string, req *pb.DictType) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	data := model.SysDictType{
		DictName: req.DictName,
		DictType: req.DictType,
		Status:   req.Status,
		Remark:   req.Remark,
	}
	data.CreateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.CreateDictType(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateDictType(ctx context.Context, req *pb.DictType) (err error) {
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
	data := model.SysDictType{
		DictId:   req.DictId,
		DictName: req.DictName,
		DictType: req.DictType,
		Status:   req.Status,
		Remark:   req.Remark,
	}
	data.UpdateBy = fmt.Sprint(permission.UserId)
	_, err = uc.d.UpdateDictType(&data, int(data.DictId))
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) DeleteDictType(ctx context.Context, req *pb.DeleteDictTypeRequest) (err error) {
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
	ids, err := strings.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteDictType(&data, ids)
	if err != nil {
		return
	}
	return
}
