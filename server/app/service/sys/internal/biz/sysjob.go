package biz

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/pkg/strings"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func (uc *AdminUsecase) CreateSysjob(ctx context.Context, req *pb.Job) (err error) {
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
	_, err = uc.d.CreateSysJob(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateSysjob(ctx context.Context, req *pb.Job) (err error) {
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
	data.UpdateBy = fmt.Sprint(permission.UserId)
	_, err = uc.d.UpdateSysJob(&data, data.JobId)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) DeleteSysjob(ctx context.Context, req *pb.DeleteSysjobRequest) (err error) {
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
	ids, err := strings.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteSysJob(&data, ids)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) GetSysjob(ctx context.Context, req *pb.GetSysjobRequest) (reply *pb.Job, err error) {
	var data model.SysJob
	data.JobId = int(req.JobId)
	result, err := uc.d.GetSysJob(&data)
	if err != nil {
		return
	}
	reply = &pb.Job{
		JobId: int32(result.JobId),
	}
	return
}

func (uc *AdminUsecase) ListSysjobPage(ctx context.Context, req *pb.ListSysjobRequest) (page []*pb.Job, total int64, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)

	var data model.SysJob
	data.JobName = req.JobName
	data.JobGroup = req.JobGroup
	data.CronExpression = req.CronExpression
	data.InvokeTarget = req.InvokeTarget
	data.Status = int(req.Status)
	result, total, err := uc.d.GetSysJobPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	page = make([]*pb.Job, 0)
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
		page = append(page, d)
	}
	return
}

func (uc *AdminUsecase) RemoveJob(ctx context.Context, req *pb.RemoveJobRequest) (reply *pb.ApiReply, err error) {
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

func (uc *AdminUsecase) StartJob(ctx context.Context, req *pb.StartJobRequest) (reply *pb.ApiReply, err error) {
	// var data model.SysJob
	// data.JobId = int(req.JobId)
	// _, err = s.adao.GetSysJob(&data)
	// if err != nil {
	// 	return
	// }
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
