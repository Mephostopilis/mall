package biz

import (
	"context"
	"fmt"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/pkg/strings"
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
	data.PostId = (req.PostId)
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

func (uc *AdminUsecase) GetPost(ctx context.Context, req *pb.GetPostRequest) (reply *pb.Post, err error) {
	var filter model.SysPost
	filter.PostId = (req.PostId)
	it, err := uc.d.GetPost(&filter)
	if err != nil {
		return
	}
	reply = &pb.Post{
		PostId:   int32(it.PostId),
		PostName: it.PostName,
		PostCode: it.PostCode,
		Sort:     int32(it.Sort),
		Status:   it.Status,
		Remark:   it.Remark,
	}
	return
}

func (uc *AdminUsecase) CreatePost(c context.Context, req *pb.Post) (err error) {
	data := model.SysPost{
		PostName: req.PostName,
		PostCode: req.PostCode,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
	}
	// data.CreateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.CreatePost(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdatePost(ctx context.Context, req *pb.Post) (err error) {
	// dp, err := meta.GetDataPermissions(ctx)
	data := model.SysPost{
		PostId:   req.PostId,
		PostName: req.PostName,
		PostCode: req.PostCode,
		Sort:     req.Sort,
		Status:   req.Status,
		Remark:   req.Remark,
	}
	// data.UpdateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.UpdatePost(&data, int(data.PostId))
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (err error) {
	dp, err := meta.GetDataPermissions(ctx)
	if err != nil {
		return
	}
	var data model.SysPost
	data.UpdateBy = fmt.Sprint(dp.UserId)
	ids, err := strings.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePost(&data, ids)
	if err != nil {
		return
	}
	return
}
