package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *AdminService) CreateSysjob(ctx context.Context, req *pb.Job) (reply *pb.ApiReply, err error) {
	err = s.uc.CreateSysjob(ctx, req)
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
	err = s.uc.UpdateSysjob(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) DeleteSysjob(ctx context.Context, req *pb.DeleteSysjobRequest) (reply *pb.ApiReply, err error) {
	err = s.uc.DeleteSysjob(ctx, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) GetSysjob(ctx context.Context, req *pb.GetSysjobRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.GetSysjob(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(result)
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
	result, count, err := s.uc.ListSysjobPage(ctx, req)
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
