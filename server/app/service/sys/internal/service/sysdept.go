package service

import (
	"context"

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

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/deptList [get]
// @Security Bearer
func (s *AdminService) ListDept(ctx context.Context, req *pb.ListDeptRequest) (reply *pb.ApiReply, err error) {
	var Dept model.SysDept
	Dept.DeptName = req.DeptName
	Dept.Status = req.Status
	Dept.DeptId = int(req.DeptId)
	result, err := s.uc.ListDeptTree(ctx)
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

func (s *AdminService) GetDeptTree(ctx context.Context, req *pb.GetDeptTreeRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.ListDeptTree(ctx)
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

// @Summary 部门列表数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func (s *AdminService) GetDept(ctx context.Context, req *pb.GetDeptRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetDept(ctx, req.DeptId)
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

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body pb.Dept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/dept [post]
// @Security Bearer
func (s *AdminService) CreateDept(ctx context.Context, req *pb.Dept) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.CreateDept(ctx, token, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
	}
	return
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body pb.Dept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/dept [put]
// @Security Bearer
func (s *AdminService) UpdateDept(ctx context.Context, req *pb.Dept) (reply *pb.ApiReply, err error) {
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
	if err = s.uc.UpdateDept(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/dept/{id} [delete]
func (s *AdminService) DeleteDept(ctx context.Context, req *pb.DeleteDeptRequest) (reply *pb.ApiReply, err error) {
	// var data models.SysDept
	// _, err = s.adao.DeleteSysDept(&data, int(req.Id))
	// if err != nil {
	// 	return
	// }
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) GetDeptTreeRoleselect(ctx context.Context, req *pb.GetDeptTreeRoleselectRequest) (reply *pb.ApiReply, err error) {
	var Dept model.SysDept
	depts, err := s.uc.ListDeptLableTree(&Dept)
	if err != nil {
		return
	}

	// TODO: 修改角色与部门之间设计
	// var SysRole models.SysRole
	// SysRole.RoleId = int(req.RoleId)
	// menuIds, err := s.adao.GetRoleDeptId(&SysRole)
	// if err != nil {
	// 	return
	// }
	// checkedKeys := make([]int32, 0)
	// for _, k := range menuIds {
	// 	checkedKeys = append(checkedKeys, int32(k))
	// }

	list := make([]*anypb.Any, 0)
	d := &pb.GetDeptTreeRoleselectReply{
		Depts: depts,
		// CheckedKeys: checkedKeys,
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
