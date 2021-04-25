package biz

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) GetDictTypeList(ctx context.Context, token string, req *pb.GetDictTypeListRequest) (reply *pb.ApiReply, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysDictType
	data.DictId = int(req.DictId)
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

func (uc *AdminUsecase) InsertDictType(ctx context.Context, token string, req *pb.DictType) (reply *pb.ApiReply, err error) {
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
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}
