package biz

import (
	"context"
	"fmt"
	"time"

	pb "edu/api/sso/v1"
	uuidpb "edu/api/uuid"
	"edu/pkg/ecode"
	"edu/pkg/fastdfs"
	"edu/pkg/tools"
	"edu/service/sso/internal/dao"
	"edu/service/sso/internal/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GreeterUsecase struct {
	repo    dao.Dao
	uuidc   uuidpb.UUIDClient
	log     *log.Helper
	manager *manage.Manager
	oc      *server.Config
}

func NewGreeterUsecase(repo dao.Dao, logger log.Logger) *GreeterUsecase {
	ctx := context.Background()
	uuidc, err := uuidpb.NewUUID(ctx)
	if err != nil {
		panic(err)
	}
	manager := manage.NewManager()
	//设置授权码过期时间
	manager.SetAuthorizeCodeExp(time.Minute * 10)
	// 设置令牌
	mcfg := &manage.Config{
		// 访问令牌过期时间（默认为2小时）
		AccessTokenExp: time.Hour * 2,
		// 更新令牌过期时间（默认为72小时）
		RefreshTokenExp: time.Hour * 24 * 3,
		// 是否生成更新令牌（默认为true）
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(mcfg)

	manager.MapAuthorizeGenerate(generates.NewAuthorizeGenerate())
	// manager.MapAccessGenerate(generates.NewAccessGenerate())
	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("bing", []byte("00000000"), jwt.SigningMethodHS512))
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapClientStorage(repo)
	return &GreeterUsecase{
		repo:    repo,
		uuidc:   uuidc,
		manager: manager,
		log:     log.NewHelper(log.With(logger, "module", "usecase/greeter")),
		oc:      server.NewConfig(),
	}
}

func (uc *GreeterUsecase) GenID(ctx context.Context) (id uint64, err error) {
	reply, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	return reply.Id, nil
}

func (uc *GreeterUsecase) RefreshToken(ctx context.Context, req *pb.RefreshTokenReq) (reply *pb.RefreshTokenResp, err error) {
	ti, err := uc.manager.LoadRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	uc.log.Infof("ti = %v", ti)
	reply = &pb.RefreshTokenResp{
		AccessToken:  ti.GetAccess(),
		ExpiresIn:    fmt.Sprintf("%f", ti.GetAccessExpiresIn().Seconds()),
		Scope:        ti.GetScope(),
		RefreshToken: ti.GetRefresh(),
	}
	return
}

func (uc *GreeterUsecase) GetInfo(ctx context.Context, dp *pb.DataPermission, req *pb.GetInfoRequest) (reply *pb.GetInfoReply, err error) {
	var roles = make([]string, 1)
	roles[0] = dp.RoleKey

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	sysuser := model.SysUser{}
	sysuser.UserId = dp.UserId
	user, err := uc.repo.GetSysUserInfo(&sysuser)
	if err != nil {
		return
	}

	d := &pb.GetInfoReply{
		Roles: roles,
	}
	if dp.RoleKey == "admin" || dp.RoleKey == "系统管理员" {
		d.Permissions = permissions
		d.Buttons = buttons
	} else {
		// RoleMenu := model.RoleMenu{}
		// RoleMenu.RoleId = dp.RoleKey
		// list, err1 := s.adao.GetRoleMenuPermisFromSysMenu(&RoleMenu)
		// if err1 != nil {
		// 	err = err1
		// 	return
		// }
		// d.Permissions = list
		// d.Buttons = list
	}
	d.Introduction = " am a super administrator"
	d.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if user.Avatar != "" {
		d.Avatar = user.Avatar
	}
	d.UserName = user.NickName
	d.UserId = int64(user.UserId)
	d.DeptId = int64(user.DeptId)
	d.Name = user.NickName
	reply = d
	return
}

// 获取用户个人信息
func (uc *GreeterUsecase) GetSysUserProfile(ctx context.Context, dp *pb.DataPermission, req *pb.GetSysUserProfileRequest) (reply *pb.ApiReply, err error) {
	var sysUser model.SysUser
	sysUser.UserId = dp.UserId
	result, err := uc.repo.GetSysUserInfo(&sysUser)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.GetSysUserProfileReply{}
	d.Result = &pb.SysUser{
		UserId:    int64(result.UserId),
		NickName:  result.NickName,
		Phone:     result.Phone,
		RoleId:    int32(result.RoleId),
		Avatar:    result.Avatar,
		Sex:       result.Sex,
		Email:     result.Email,
		DeptId:    int32(result.DeptId),
		PostId:    int32(result.PostId),
		Remark:    result.Remark,
		Status:    result.Status,
		CreatedAt: timestamppb.New(result.CreatedAt),
	}

	//获取角色列表
	roles, err := uc.repo.GetSysRoleList(&model.SysRole{})
	if err != nil {
		return
	}
	d.RoleIds = make([]int32, 0)
	d.RoleIds = append(d.RoleIds, int32(result.RoleId))
	d.Roles = make([]*pb.Role, 0)
	for i := 0; i < len(roles); i++ {
		role := roles[i]
		d.Roles = append(d.Roles, &pb.Role{
			RoleId:   int64(role.RoleId),
			RoleName: role.RoleName,
			Status:   role.Status,
			RoleKey:  role.RoleKey,
			RoleSort: int32(role.RoleSort),
			Flag:     role.Flag,
			Remark:   role.Remark,
			// Admin    ：role.ad
			// MenuIds  []int32 `protobuf:"varint,9,rep,packed,name=MenuIds,proto3" json:"MenuIds,omitempty"`
		})
	}

	// TODO: 获取职位
	//获取职位列表
	// posts, err := uc.sysc.ListPost(ctx, &syspb.ListPostRequest{})
	// posts, err := uc.repo.GetPostList(&models.SysPost{})
	// if err != nil {
	// 	return
	// }
	// d.PostIds = make([]int32, 0)
	// d.PostIds = append(d.PostIds, int32(result.PostId))
	// d.Posts = make([]*pb.Post, 0)
	// for i := 0; i < len(posts); i++ {
	// 	d.Posts = append(d.Posts, &pb.Post{})
	// }

	//获取部门列表
	// dept, err := uc.sysc.GetDept(ctx, &syspb.GetDeptRequest{DeptId: int32(result.DeptId)})
	// if err != nil {
	// 	return
	// }
	// d.Dept = &pb.Dept{
	// 	DeptId:   int64(dept.DeptId),
	// 	ParentId: int64(dept.ParentId),
	// 	DeptPath: dept.DeptPath,
	// 	DeptName: dept.DeptName,
	// 	Sort:     int32(dept.Sort),
	// 	Leader:   dept.Leader,
	// 	Phone:    dept.Phone,
	// 	Email:    dept.Email,
	// 	Status:   dept.Status,
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

func (uc *GreeterUsecase) InsetSysUserAvatar(ctx context.Context, req *pb.InsetSysUserAvatarRequest) (reply *pb.ApiReply, err error) {
	res, err := fastdfs.Upload("127.0.0.1", string(req.Avatar))
	if err != nil {
		return
	}
	uc.log.Debugf("res = %v", res)
	// sysuser := models.SysUser{}
	// sysuser.UserId = jwt.GetUserId(ctx)
	// sysuser.Avatar = "/" + filPath
	// sysuser.UpdateBy = jwt.GetUserIdStr(ctx)
	// result, err := s.dao.UpdateSysUser(&sysuser, sysuser.UserId)
	// if err != nil {
	// 	return
	// }
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (uc *GreeterUsecase) SetSysUserPwd(ctx context.Context, uid uint64, req *pb.SysUserUpdatePwdRequest) (err error) {
	user, err := uc.repo.GetSysUserInfo(&model.SysUser{SysUserId: model.SysUserId{UserId: uid}})
	if err != nil {
		return
	}
	_, err = tools.CompareHashAndPassword(user.Password, req.OldPassword)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	_, err = uc.repo.UpdateSysUser(&model.SysUser{LoginM: model.LoginM{PassWord: model.PassWord{Password: req.NewPassword}}}, uid)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) GetMemberInfo(ctx context.Context, dp *pb.DataPermission, req *pb.GetInfoRequest) (reply *pb.ApiReply, err error) {
	// var roles = make([]string, 0)
	// roles[0] = jwtauth.GetRoleName(ctx)

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	// RoleMenu := model.SysRoleMenu{}
	// RoleMenu.RoleId = int32(jwtauth.GetRoleId(ctx))

	// sysuser := model.SysUser{}
	// sysuser.UserId = jwtauth.GetUserId(ctx)
	// user, err := uc.repo.GetSysUser(&sysuser)
	// if err != nil {
	// 	return
	// }

	// list := make([]*anypb.Any, 0)
	// d := &pb.GetInfoReply{
	// 	Roles: roles,
	// }
	// if dp.RoleKey == "admin" || dp.RoleKey == "系统管理员" {
	// 	d.Permissions = permissions
	// 	d.Buttons = buttons
	// } else {
	// list, err1 := uc.sysc.GetRoleMenuPermisFromSysMenu(&RoleMenu)
	// if err1 != nil {
	// 	err = err1
	// 	return
	// }
	// d.Permissions = list
	// d.Buttons = list
	// }
	// d.Introduction = "am a super administrator"
	// d.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	// if user.Avatar != "" {
	// 	d.Avatar = user.Avatar
	// }
	// d.UserName = user.NickName
	// d.UserId = int64(user.UserId)
	// d.DeptId = int64(user.DeptId)
	// d.Name = user.NickName

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
