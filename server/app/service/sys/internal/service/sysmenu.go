package service

import (
	"context"
	"fmt"

	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/service/sys/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/menulist [get]
// @Security Bearer
func (s *AdminService) ListMenu(ctx context.Context, req *pb.ListMenuRequest) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	result, err := s.uc.ListMenu(ctx, token, req)
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

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/menu [get]
// @Security Bearer
func (s *AdminService) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetMenu(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
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

// func (s *AdminService) GetMenuTreeRoleselect(ctx context.Context, req *pb.GetMenuTreeRoleselectRequest) (reply *pb.ApiReply, err error) {
// dp, err := meta.GetDataPermissions(ctx)
// if err != nil {
// 	return
// }
// var Menu model.Menu
// result, err := s.adao.GetMenuLable(&Menu)
// if err != nil {
// 	return
// }
// var SysRole model.SysRole
// SysRole.RoleId = int(req.RoleId)
// menuIds, err := s.adao.GetRoleMeunId(&SysRole)
// if err != nil {
// 	return
// }

// menuids
// checkedKeys := make([]int32, 0)
// for _, id := range menuIds {
// 	checkedKeys = append(checkedKeys, int32(id))
// }
// menus := s.ConvertMenuLable(ctx, result)

// list := make([]*anypb.Any, 0)
// d := &pb.GetMenuTreeRoleselectReply{
// 	Menus:       menus,
// 	CheckedKeys: checkedKeys,
// }
// any, err1 := ptypes.MarshalAny(d)
// if err1 != nil {
// 	err = err1
// 	return
// }
// list = append(list, any)
// 	reply = &pb.ApiReply{
// 		Code:    0,
// 		Message: "OK",
// 		Count:   int32(1),
// 		Data:    list,
// 	}
// 	return
// }

// @Summary 获取菜单树
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/menuTreeselect [get]
// @Security Bearer
func (s *AdminService) GetMenuTreeelect(ctx context.Context, req *pb.GetMenuTreeelectRequest) (reply *pb.ApiReply, err error) {
	var data model.Menu
	result, err := s.uc.ListMenuLable(&data)
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

// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param menuName formData string true "menuName"
// @Param Path formData string false "Path"
// @Param Action formData string true "Action"
// @Param Permission formData string true "Permission"
// @Param ParentId formData string true "ParentId"
// @Param IsDel formData string true "IsDel"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/menu [post]
// @Security Bearer
func (s *AdminService) CreateMenu(ctx context.Context, req *pb.Menu) (reply *pb.ApiReply, err error) {
	tok, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.CreateMenu(ctx, tok, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param data body pb.Menu true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /admin/v1/menu/{id} [put]
// @Security Bearer
func (s *AdminService) UpdateMenu(ctx context.Context, req *pb.Menu) (reply *pb.ApiReply, err error) {
	tok, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	err = s.uc.UpdateMenu(ctx, tok, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/menu/{id} [delete]
func (s *AdminService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (reply *pb.ApiReply, err error) {
	// tok, err := meta.GetToken(ctx)
	// if err != nil {
	// 	return
	// }
	// var it model.Menu
	// it.UpdateBy = fmt.Sprint(dp.UserId)
	err = s.uc.DeleteMenu(ctx, req)
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

// @Summary 根据角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/menurole [get]
// @Security Bearer
func (s *AdminService) GetMenuRole(ctx context.Context, req *pb.GetMenuRoleRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.ListMenuByRole(ctx, &model.Menu{}, "")
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

// @Summary 获取角色对应的菜单id数组
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/menuids/{id} [get]
// @Security Bearer
func (s *AdminService) GetMenuIDS(ctx context.Context, req *pb.GetMenuIDSRequest) (reply *pb.ApiReply, err error) {
	dp, err := meta.GetDataPermissions(ctx)
	if err != nil {
		return
	}
	var data model.SysRoleMenu
	data.RoleName = dp.RoleKey
	data.UpdateBy = fmt.Sprint(dp.UserId)
	result, err := s.adao.GetRoleMenuIDSFromSysMenu(&data)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.MenuPath{
			Path: it.Path,
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
