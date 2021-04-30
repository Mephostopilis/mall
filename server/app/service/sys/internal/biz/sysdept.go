package biz

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) GetDept(ctx context.Context, deptID int32) (reply *pb.Dept, err error) {
	it, err := uc.d.GetSysDept(&model.SysDept{DeptId: int(deptID)})
	if err != nil {
		return
	}
	reply = &pb.Dept{
		DeptId:    int64(it.DeptId),
		ParentId:  int64(it.ParentId),
		DeptPath:  it.DeptPath,
		DeptName:  it.DeptName,
		Sort:      int32(it.Sort),
		Leader:    it.Leader,
		Phone:     it.Phone,
		Email:     it.Email,
		Status:    it.Status,
		CreatedAt: timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *AdminUsecase) CreateDept(ctx context.Context, token string, req *pb.Dept) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	data := model.SysDept{
		DeptId:   int(req.DeptId),
		ParentId: int(req.ParentId),
		DeptPath: req.DeptPath,
		DeptName: req.DeptName,
		Sort:     int(req.Sort),
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	}
	data.CreateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.CreateSysDept(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateDept(ctx context.Context, req *pb.Dept) (err error) {
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
	data := model.SysDept{
		DeptId:   int(req.DeptId),
		ParentId: int(req.ParentId),
		DeptPath: req.DeptPath,
		DeptName: req.DeptName,
		Sort:     int(req.Sort),
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	}
	data.UpdateBy = fmt.Sprint(permission.UserId)
	_, err = uc.d.UpdateSysDept(&data, data.DeptId)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) ListDeptTree(ctx context.Context) (list []*pb.Dept, err error) {
	depts, err := uc.d.GetSysDeptList(&model.SysDept{})
	if err != nil {
		return
	}
	list = make([]*pb.Dept, 0)
	for i := 0; i < len(depts); i++ {
		it := depts[i]
		if it.ParentId != 0 {
			continue
		}
		info := diguiDept(depts, it)
		list = append(list, info)
	}
	return
}

func diguiDept(deptlist []model.SysDept, top model.SysDept) *pb.Dept {
	d := &pb.Dept{
		DeptId:    int64(top.DeptId),
		ParentId:  int64(top.ParentId),
		DeptPath:  top.DeptPath,
		DeptName:  top.DeptName,
		Sort:      int32(top.Sort),
		Leader:    top.Leader,
		Phone:     top.Phone,
		Email:     top.Email,
		Status:    top.Status,
		CreatedAt: timestamppb.New(top.CreatedAt),
	}
	min := make([]*pb.Dept, 0)
	for j := 0; j < len(deptlist); j++ {
		it := deptlist[j]
		if it.ParentId != top.DeptId {
			continue
		}
		ms := diguiDept(deptlist, it)
		min = append(min, ms)
	}
	d.Children = min
	return d
}

func (uc *AdminUsecase) ListDeptLableTree(e *model.SysDept) (list []*pb.DeptLable, err error) {
	deptlist, err := uc.d.GetSysDeptList(e)

	list = make([]*pb.DeptLable, 0)
	for i := 0; i < len(deptlist); i++ {
		it := deptlist[i]
		if it.ParentId != 0 {
			continue
		}
		deptsInfo := diguiDeptLable(deptlist, &it)
		list = append(list, deptsInfo)
	}
	return
}

func diguiDeptLable(deptlist []model.SysDept, dept *model.SysDept) *pb.DeptLable {
	d := &pb.DeptLable{
		Id:    int32(dept.DeptId),
		Label: dept.DeptName,
	}
	min := make([]*pb.DeptLable, 0)
	for j := 0; j < len(deptlist); j++ {
		it := deptlist[j]
		if it.DeptId != dept.DeptId {
			continue
		}
		ms := diguiDeptLable(deptlist, &it)
		min = append(min, ms)
	}
	d.Children = min
	return d
}
