package service

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *AdminService) CreateSysjob(ctx context.Context, req *pb.Job) (reply *pb.ApiReply, err error) {
	// dp, err := meta.GetDataPermissions(ctx)
	// if err != nil {
	// 	return
	// }
	data := model.SysJob{
		JobName:        req.JobName,
		JobGroup:       req.JobGroup,
		JobType:        int(req.JobType),
		CronExpression: req.CronExpression,
		InvokeTarget:   req.InvokeTarget,
		MisfirePolicy:  int(req.MisfirePolicy),
		Concurrent:     int(req.Concurrent),
		Status:         int(req.Status),
		EntryId:        int(req.EntryId),
	}
	// data.CreateBy = fmt.Sprint(dp.UserId)
	_, err = s.adao.CreateSysJob(&data)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateSysjob(ctx context.Context, req *pb.Job) (reply *pb.ApiReply, err error) {
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

	var data model.SysJob
	data.UpdateBy = fmt.Sprint(permission.UserId)
	result, err := s.adao.UpdateSysJob(&data, data.JobId)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.Job{
		JobId: int32(result.JobId),
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
		Data:    list,
	}
	return
}

func (s *AdminService) DeleteSysjob(ctx context.Context, req *pb.DeleteSysjobRequest) (reply *pb.ApiReply, err error) {
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
	var data model.SysJob
	data.UpdateBy = fmt.Sprint(permission.UserId)

	// IDS := jwt.IdsStrToIdsIntGroup("jobId", ctx)
	IDS := make([]int, 0)
	_, err = s.adao.BatchDeleteSysJob(&data, IDS)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.Job{}
	any, err1 := ptypes.MarshalAny(d)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Data:    list,
	}
	return
}

func (s *AdminService) GetSysjob(ctx context.Context, req *pb.GetSysjobRequest) (reply *pb.ApiReply, err error) {
	var data model.SysJob
	data.JobId = int(req.JobId)
	result, err := s.adao.GetSysJob(&data)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	d := &pb.Job{
		JobId: int32(result.JobId),
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
		Data:    list,
	}
	return
}

func (s *AdminService) ListSysjob(ctx context.Context, req *pb.ListSysjobRequest) (reply *pb.ApiReply, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysJob
	data.JobName = req.JobName
	data.JobGroup = req.JobGroup
	data.CronExpression = req.CronExpression
	data.InvokeTarget = req.InvokeTarget
	data.Status = int(req.Status)
	result, count, err := s.adao.GetSysJobPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.Job{
			JobId:          int32(it.JobId),
			JobName:        it.JobName,
			JobGroup:       it.JobGroup,
			JobType:        int32(it.JobType),
			CronExpression: it.CronExpression,
			InvokeTarget:   it.InvokeTarget,
			MisfirePolicy:  int32(it.MisfirePolicy),
			Concurrent:     int32(it.Concurrent),
			Status:         int32(it.Status),
			EntryId:        int32(it.EntryId),
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

func (s *AdminService) RemoveJob(ctx context.Context, req *pb.RemoveJobRequest) (reply *pb.ApiReply, err error) {
	var data model.SysJob
	data.JobId = int(req.JobId)
	// result, err := s.adao.GetSysJob(&data)
	// if err != nil {
	// 	return
	// }

	// cn := s.adao.RemoveJob(ctx, result.EntryId)
	// select {
	// case res := <-cn:
	// 	if res {
	// 		_, _ = s.adao.RemoveSysJobEntryID(&result, result.EntryId)
	// 		return
	// 	}
	// case <-time.After(time.Second * 1):
	// 	return
	// }
	return
}

func (s *AdminService) StartJob(ctx context.Context, req *pb.StartJobRequest) (reply *pb.ApiReply, err error) {
	var data model.SysJob
	data.JobId = int(req.JobId)
	_, err = s.adao.GetSysJob(&data)
	if err != nil {
		return
	}
	// if result.JobType == 1 {
	// 	var j = &model.HttpJob{}
	// 	j.InvokeTarget = result.InvokeTarget
	// 	j.CronExpression = result.CronExpression
	// 	j.JobId = result.JobId
	// 	j.Name = result.JobName
	// 	data.EntryId, err = svc.AddJob(j)
	// 	if err != nil {
	// 		c.JSON(nil, err)
	// 		return
	// 	}
	// } else {
	// 	var j = &model.ExecJob{}
	// 	j.InvokeTarget = result.InvokeTarget
	// 	j.CronExpression = result.CronExpression
	// 	j.JobId = result.JobId
	// 	j.Name = result.JobName
	// 	data.EntryId, err = svc.AddJob(j)
	// 	if err != nil {
	// 		c.JSON(nil, err)
	// 		return
	// 	}
	// }

	// result, err = s.adao.UpdateSysJob(&data, data.JobId)
	// if err != nil {
	// 	return
	// }
	return
}
