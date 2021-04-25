package service

import (
	"context"

	pb "edu/api/sso"

	"edu/service/sso/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /admin/v1/sysUserList [get]
// @Security Bearer
func (s *AdminService) ListSysUser(ctx context.Context, req *pb.ListSysUserRequest) (reply *pb.ApiReply, err error) {

	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysUser
	data.Username = req.Username
	data.Status = req.Status
	data.Phone = req.Phone
	data.DeptId = int(req.DeptId)
	data.PostId = int(req.PostId)

	result, count, err := s.adao.GetSysUserPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.SysUser{
			UserId:    int64(it.UserId),
			NickName:  it.NickName,
			Phone:     it.Phone,
			RoleId:    int32(it.RoleId),
			Avatar:    it.Avatar,
			Sex:       it.Sex,
			Email:     it.Email,
			DeptId:    int32(it.DeptId),
			PostId:    int32(it.PostId),
			Remark:    it.Remark,
			Status:    it.Status,
			CreatedAt: timestamppb.New(it.CreatedAt),
			Username:  it.Username,
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

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/sysUser/{userId} [get]
// @Security Bearer
func (s *AdminService) GetSysUser(ctx context.Context, req *pb.GetSysUserRequest) (reply *pb.ApiReply, err error) {
	var SysUser model.SysUser
	SysUser.UserId = uint64(req.UserId)
	it, err := s.adao.GetSysUserInfo(&SysUser)
	if err != nil {
		return
	}

	list := make([]*anypb.Any, 0)
	d := &pb.GetSysUserReply{
		Data: &pb.SysUser{
			UserId:   int64(it.UserId),
			NickName: it.NickName,
			Phone:    it.Phone,
			RoleId:   int32(it.RoleId),
			Avatar:   it.Avatar,
			Sex:      it.Sex,
			Email:    it.Email,
			DeptId:   int32(it.DeptId),
			PostId:   int32(it.PostId),
			Remark:   it.Remark,
			Status:   it.Status,
		},
		PostIds: make([]int32, 0),
		RoleIds: make([]int32, 0),
		Roles:   make([]*pb.Role, 0),
		Posts:   make([]*pb.Post, 0),
	}
	d.PostIds = append(d.PostIds, int32(it.PostId))
	d.RoleIds = append(d.RoleIds, int32(it.RoleId))

	var SysRole model.SysRole
	roles, err := s.adao.GetSysRoleList(&SysRole)
	if err != nil {
		return
	}
	for _, role := range roles {
		d.Roles = append(d.Roles, &pb.Role{
			RoleId:    int64(role.RoleId),
			RoleName:  role.RoleName,
			Status:    role.Status,
			RoleKey:   role.RoleKey,
			RoleSort:  int32(role.RoleSort),
			Flag:      role.Flag,
			Remark:    role.Remark,
			Admin:     role.Admin,
			CreatedAt: timestamppb.New(role.CreatedAt),
		})
	}

	// var post models.SysPost
	// posts, err := s.adao.GetPostList(&post)
	// if err != nil {
	// 	return
	// }
	// for _, post := range posts {
	// 	d.Posts = append(d.Posts, &pb.Post{
	// 		PostId:   int32(post.PostId),
	// 		PostName: post.PostName,
	// 		PostCode: post.PostCode,
	// 		Sort:     int32(post.Sort),
	// 		Status:   post.Status,
	// 		Remark:   post.Remark,
	// 	})
	// }

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

// @Summary 获取用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/v1/sysUser [get]
// @Security Bearer
func (s *AdminService) GetSysUserInit(ctx context.Context, req *pb.GetSysUserInitRequest) (reply *pb.ApiReply, err error) {

	list := make([]*anypb.Any, 0)
	d := &pb.GetSysUserInitReply{
		Roles: make([]*pb.Role, 0),
		Posts: make([]*pb.Post, 0),
	}
	var SysRole model.SysRole
	roles, err := s.adao.GetSysRoleList(&SysRole)
	if err != nil {
		return
	}
	for _, role := range roles {
		d.Roles = append(d.Roles, &pb.Role{
			RoleId:    int64(role.RoleId),
			RoleName:  role.RoleName,
			Status:    role.Status,
			RoleKey:   role.RoleKey,
			RoleSort:  int32(role.RoleSort),
			Flag:      role.Flag,
			Remark:    role.Remark,
			Admin:     role.Admin,
			CreatedAt: timestamppb.New(role.CreatedAt),
		})
	}

	// posts, err := s.adao.GetPostList(&models.SysPost{})
	// if err != nil {
	// 	return
	// }
	// for _, post := range posts {
	// 	d.Posts = append(d.Posts, &pb.Post{
	// 		PostId:   int32(post.PostId),
	// 		PostName: post.PostName,
	// 		PostCode: post.PostCode,
	// 		Sort:     int32(post.Sort),
	// 		Status:   post.Status,
	// 		Remark:   post.Remark,
	// 	})
	// }

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

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body pb.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /admin/v1/sysUser [post]

func (s *AdminService) CreateSysUser(ctx context.Context, req *pb.SysUser) (reply *pb.ApiReply, err error) {
	err = s.uc.CreateSysUser(ctx, req)
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

// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body pb.SysUser true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /admin/v1/sysuser/{userId} [put]
func (s *AdminService) UpdateSysUser(ctx context.Context, req *pb.SysUser) (reply *pb.ApiReply, err error) {
	// var data models.SysUser
	// data.UpdateBy = jwt.GetUserIdStr(ctx)
	// result, err := s.adao.UpdateSysUser(&data, data.UserId)
	// if err != nil {
	// 	return
	// }
	// list := make([]*anypb.Any, 1)
	// d := &pb.SysUser{}
	// any, err1 := ptypes.MarshalAny(d)
	// if err1 != nil {
	// 	err = err1
	// 	return
	// }
	// list = append(list, any)
	// reply = &pb.ApiReply{
	// 	Code:    0,
	// 	Message: "OK",
	// 	Count:   int32(1),
	// 	Data:    list,
	// }
	return
}

// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /admin/v1/sysuser/{userId} [delete]
func (s *AdminService) DeleteSysUser(ctx context.Context, req *pb.DeleteSysUserRequest) (reply *pb.ApiReply, err error) {
	// var data models.SysUser
	// data.UpdateBy = jwt.GetUserIdStr(c)
	// IDS := jwt.IdsStrToIdsIntGroup("userId", c)
	// result, err := s.adao.BatchDeleteSysUser(&data, IDS)
	// if err != nil {
	// 	return
	// }
	return
}
