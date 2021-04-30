package service

import (
	"context"

	pb "edu/api/sys/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *AdminService) Preview(ctx context.Context, req *pb.PreviewRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.Preview(ctx, req)
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

func (s *AdminService) GenCode(ctx context.Context, req *pb.GenCodeRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.GenCode(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictId path int true "dictId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/genMenu/{dictId} [delete]
func (s *AdminService) GenMenuAndApi(ctx context.Context, req *pb.GenMenuAndApiRequest) (reply *pb.ApiReply, err error) {

	d, err := s.uc.GenMenuAndApi(ctx, req)
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
