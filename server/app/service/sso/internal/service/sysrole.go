package service

import (
	"context"
	"fmt"

	pb "edu/api/sso"
	"edu/pkg/meta"
	"edu/service/sso/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色/Role
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/rolelist [get]
// @Security Bearer
func (s *AdminService) ListRole(c context.Context, req *pb.ListRoleRequest) (reply *pb.ApiReply, err error) {

	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysRole
	data.RoleKey = req.RoleKey
	data.RoleName = req.RoleName
	data.Status = req.Status
	result, total, err := s.adao.GetSysRolePage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.Role{
			RoleId:    int64(it.RoleId),
			RoleName:  it.RoleName,
			Status:    it.Status,
			RoleKey:   it.RoleKey,
			RoleSort:  int32(it.RoleSort),
			Flag:      it.Flag,
			Remark:    it.Remark,
			Admin:     it.Admin,
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

// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param roleId path string false "roleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/role [get]
// @Security Bearer
func (s *AdminService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (reply *pb.ApiReply, err error) {
	role := model.SysRole{
		RoleId: int(req.RoleId),
	}
	it, err := s.adao.GetSysRole(&role)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.Role{
		RoleId:    int64(it.RoleId),
		RoleName:  it.RoleName,
		Status:    it.Status,
		RoleKey:   it.RoleKey,
		RoleSort:  int32(it.RoleSort),
		Flag:      it.Flag,
		Remark:    it.Remark,
		Admin:     it.Admin,
		CreatedAt: timestamppb.New(it.CreatedAt),
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

// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body models.SysRole true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/role [post]
func (s *AdminService) CreateRole(ctx context.Context, req *pb.Role) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	dp, err := s.jwtuc.ValidationMidToken(ctx, token)
	if err != nil {
		return
	}
	data := model.SysRole{
		RoleName: req.RoleName,
		Status:   req.Status,
		RoleKey:  req.RoleKey,
		RoleSort: int(req.RoleSort),
		Flag:     req.Flag,
		Remark:   req.Remark,
		Admin:    req.Admin,
		CreateBy: fmt.Sprint(dp.UserId),
	}

	// TODO:
	// 分布式事务使用
	id, err := s.adao.InsertSysRole(&data)
	if err != nil {
		return
	}
	data.RoleId = id

	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body pb.Role true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /admin/v1/role [put]
func (s *AdminService) UpdateRole(ctx context.Context, req *pb.Role) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	dp, err := s.jwtuc.ValidationMidToken(ctx, token)
	if err != nil {
		return
	}
	data := model.SysRole{
		RoleName: req.RoleName,
		Status:   req.Status,
		RoleKey:  req.RoleKey,
		RoleSort: int(req.RoleSort),
		Flag:     req.Flag,
		Remark:   req.Remark,
		Admin:    req.Admin,
		UpdateBy: fmt.Sprint(dp.UserId),
	}
	_, err = s.adao.UpdateSysRole(&data, data.RoleId)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色/Role
// @Param roleId path int true "roleId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/role/{roleId} [delete]
func (s *AdminService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (reply *pb.ApiReply, err error) {
	var role model.SysRole
	_, err = s.adao.BatchDeleteSysRole(&role, req.Ids)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}
