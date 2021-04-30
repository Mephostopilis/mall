package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) ListOperLogPage(ctx context.Context, req *pb.ListOperLogRequest) (page []*pb.OperLog, total int64, err error) {
	var data model.SysOperLog
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	data.OperName = req.OperName
	data.Status = req.Status
	data.OperIp = req.OperIp
	result, total, err := uc.d.GetSysOperLogPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	list := make([]*pb.OperLog, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.OperLog{
			OperId:        it.OperId,
			Title:         it.Title,
			BusinessType:  it.BusinessType,
			BusinessTypes: it.BusinessTypes,
			Method:        it.Method,
			RequestMethod: it.RequestMethod,
			OperatorType:  it.OperatorType,
			OperName:      it.OperName,
			DeptName:      it.DeptName,
			OperUrl:       it.OperUrl,
			OperIp:        it.OperIp,
			OperLocation:  it.OperLocation,
			OperParam:     it.OperParam,
			Status:        it.Status,
			OperTime:      timestamppb.New(it.OperTime),
			JsonResult:    it.JsonResult,
		}
		list = append(list, d)
	}
	return
}

func (uc *AdminUsecase) GetOperLog(ctx context.Context, req *pb.GetOperLogRequest) (reply *pb.OperLog, err error) {
	var filter model.SysOperLog
	filter.OperId = req.OperId
	it, err := uc.d.GetSysOperLog(&filter)
	if err != nil {
		return
	}
	reply = &pb.OperLog{
		OperId:        it.OperId,
		Title:         it.Title,
		BusinessType:  it.BusinessType,
		BusinessTypes: it.BusinessTypes,
		Method:        it.Method,
		RequestMethod: it.RequestMethod,
		OperatorType:  it.OperatorType,
		OperName:      it.OperName,
		DeptName:      it.DeptName,
		OperUrl:       it.OperUrl,
		OperIp:        it.OperIp,
		OperLocation:  it.OperLocation,
		OperParam:     it.OperParam,
		Status:        it.Status,
		OperTime:      timestamppb.New(it.OperTime),
		JsonResult:    it.JsonResult,
	}
	return
}

func (uc *AdminUsecase) CreateOperLog(ctx context.Context, req *pb.OperLog) (err error) {
	data := model.SysOperLog{
		Title:         req.Title,
		BusinessType:  req.BusinessType,
		BusinessTypes: req.BusinessTypes,
		Method:        req.Method,
		RequestMethod: req.RequestMethod,
		OperatorType:  req.OperatorType,
		OperName:      req.OperName,
		DeptName:      req.DeptName,
		OperUrl:       req.OperUrl,
		OperIp:        req.OperIp,
		OperLocation:  req.OperLocation,
		OperParam:     req.OperParam,
		Status:        req.Status,
		OperTime:      req.OperTime.AsTime(),
		JsonResult:    req.JsonResult,
		// CreateBy:      fmt.Sprint(dp.UserId),
	}
	_, err = uc.d.CreateSysOperLog(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateOperLog(context.Context, *pb.OperLog) (err error) {
	return
}

func (uc *AdminUsecase) DeleteOperLog(ctx context.Context, req *pb.DeleteOperLogRequest) (err error) {
	// var data models.SysOperLog
	// data.UpdateBy = jwt.GetUserIdStr(ctx)
	// IDS := jwt.IdsStrToIdsIntGroup("operId", c)
	// result, err := s.adao.BatchDeleteSysOperLog(&data, IDS)
	// if err != nil {
	// 	return
	// }
	// reply = &pb.ApiReply{
	// 	Code:    0,
	// 	Message: "OK",
	// }
	return
}
