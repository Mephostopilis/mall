package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary 岗位列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/post [get]
// @Security Bearer
func (s *AdminService) ListPost(c context.Context, req *pb.ListPostRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(c)
	if err != nil {
		return
	}
	return s.uc.ListPost(c, token, req)
}

// @Summary 获取岗位信息
// @Description 获取JSON
// @Tags 岗位
// @Param postId path int true "postId"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/post/{postId} [get]
// @Security Bearer
func (s *AdminService) GetPost(ctx context.Context, req *pb.GetPostRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.GetPost(ctx, req)
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

// @Summary 添加岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body pb.Post true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/post [post]
// @Security Bearer
func (s *AdminService) CreatePost(c context.Context, req *pb.Post) (reply *pb.ApiReply, err error) {
	err = s.uc.CreatePost(c, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 修改岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body pb.Post true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/post/ [put]
// @Security Bearer
func (s *AdminService) UpdatePost(ctx context.Context, req *pb.Post) (reply *pb.ApiReply, err error) {
	err = s.uc.UpdatePost(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 删除岗位
// @Description 删除数据
// @Tags 岗位
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 500 {string} string	"{"code": 500, "message": "删除失败"}"
// @Router /admin/v1/post/{postId} [delete]
func (s *AdminService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (reply *pb.ApiReply, err error) {
	err = s.uc.DeletePost(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
