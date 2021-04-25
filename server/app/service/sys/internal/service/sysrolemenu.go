package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/service/sys/internal/model"
)

func (s *AdminService) UpdateRoleMenu(ctx context.Context, req *pb.Menu) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

func (s *AdminService) ListRoleMenu(ctx context.Context, req *pb.ListRoleMenuRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}

// @Summary RoleMenu列表数据
// @Description 获取JSON
// @Tags 角色菜单
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/rolemenu [get]
// @Security Bearer
func (s *AdminService) GetRoleMenu(ctx context.Context, req *pb.GetRoleMenuRequest) (reply *pb.ApiReply, err error) {
	// var Rm models.RoleMenu
	// result, err := s.adao.GetRoleMenu(&Rm)
	// if err != nil {
	// 	return
	// }

	// list := make([]*anypb.Any, 1)
	// d := &pb.GetDeptTreeRoleselectReply{
	// 	Depts:       depts,
	// 	CheckedKeys: checkedKeys,
	// }
	// any, err1 := ptypes.MarshalAny(d)
	// if err1 != nil {
	// 	err = err1
	// 	return
	// }
	// list = append(list, any)
	// reply = &pb.ApiReply{
	// 	Code:    0,
	// 	Message: "OK",
	// 	Data:    list,
	// }
	return
}

type RoleMenuPost struct {
	RoleId   string
	RoleMenu []model.SysRoleMenu
}

func (s *AdminService) CreateRoleMenu(ctx context.Context, req *pb.Menu) (reply *pb.ApiReply, err error) {
	return &pb.ApiReply{}, nil
}

// @Summary 删除用户菜单数据
// @Description 删除数据
// @Tags 角色菜单
// @Param id path string true "id"
// @Param menu_id query string false "menu_id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/rolemenu/{id} [delete]
func (s *AdminService) DeleteRoleMenu(ctx context.Context, req *pb.DeleteRoleMenuRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	return s.uc.DeleteRoleMenu(ctx, token, req)
}
