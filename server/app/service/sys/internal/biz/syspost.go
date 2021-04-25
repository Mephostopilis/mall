package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) ListPost(c context.Context, token string, req *pb.ListPostRequest) (reply *pb.ApiReply, err error) {
	// out, err := uc.mw.ValidationToken(token)
	// if err != nil {
	// 	return
	// }
	// dp := out.(*ssopb.DataPermission)

	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var data model.SysPost
	data.PostId = int(req.PostId)
	data.PostCode = req.PostCode
	data.PostName = req.PostName
	data.Status = req.Status
	result, count, err := uc.d.GetPostPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.Post{
			PostId:    int32(it.PostId),
			PostName:  it.PostName,
			PostCode:  it.PostCode,
			Sort:      int32(it.Sort),
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
		Count:   int32(count),
		Data:    list,
	}
	return
}
