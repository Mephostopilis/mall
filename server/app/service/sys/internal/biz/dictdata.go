package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/pkg/ecode"
	"edu/pkg/strings"
	"edu/service/sys/internal/model"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) GetDictDataPage(ctx context.Context, req *pb.GetDictDataListRequest) (list []*pb.DictData, total int64, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var data model.SysDictData
	data.DictType = req.DictType
	data.DictLabel = req.DictLabel
	data.Status = req.Status
	result, total, err := uc.d.GetDictDataPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}

	list = make([]*pb.DictData, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DictData{
			DictCode:  int32(it.DictCode),
			DictSort:  int32(it.DictSort),
			DictLabel: it.DictLabel,
			DictValue: it.DictValue,
			DictType:  it.DictType,
			CssClass:  it.CssClass,
			ListClass: it.ListClass,
			IsDefault: it.IsDefault,
			Status:    it.Status,
			Default:   it.Default,
			Remark:    it.Remark,
			CreatedAt: timestamppb.New(it.CreatedAt),
		}
		list = append(list, d)
	}
	return
}

func (uc *AdminUsecase) GetDictDataListByDictType(ctx context.Context, req *pb.GetDictDataListByDictTypeRequest) (list []*pb.DictData, err error) {
	var dictData model.SysDictData
	dictData.DictType = req.DictType
	result, err := uc.d.GetDictData(&dictData)
	if err != nil {
		return
	}
	list = make([]*pb.DictData, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DictData{
			DictCode:  int32(it.DictCode),
			DictSort:  int32(it.DictSort),
			DictLabel: it.DictLabel,
			DictValue: it.DictValue,
			DictType:  it.DictType,
			CssClass:  it.CssClass,
			ListClass: it.ListClass,
			IsDefault: it.IsDefault,
			Status:    it.Status,
			Default:   it.Default,
			Remark:    it.Remark,
			CreatedAt: timestamppb.New(it.CreatedAt),
		}
		list = append(list, d)
	}
	return
}

func (uc *AdminUsecase) GetDictData(ctx context.Context, req *pb.GetDictDataRequest) (reply *pb.DictData, err error) {
	var filter model.SysDictData
	filter.DictLabel = req.DictLabel
	filter.DictCode = req.DictCode
	it, err := uc.d.GetDictDataByCode(&filter)
	if err != nil {
		return
	}
	reply = &pb.DictData{
		DictCode:  it.DictCode,
		DictSort:  int32(it.DictSort),
		DictLabel: it.DictLabel,
		DictValue: it.DictValue,
		DictType:  it.DictType,
		CssClass:  it.CssClass,
		ListClass: it.ListClass,
		IsDefault: it.IsDefault,
		Status:    it.Status,
		Default:   it.Default,
	}
	return
}

func (uc *AdminUsecase) InsertDictData(ctx context.Context, req *pb.DictData) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = ecode.Unknown("meta", "error")
		return
	}
	v := md.Get("UserID")
	if len(v) < 0 {
		err = ecode.Unknown("meta", "error")
		return
	}
	data := model.SysDictData{
		DictSort:  req.DictSort,
		DictLabel: req.DictLabel,
		DictValue: req.DictValue,
		DictType:  req.DictType,
		CssClass:  req.CssClass,
		ListClass: req.ListClass,
		IsDefault: req.IsDefault,
		Status:    req.Status,
		Default:   req.Default,
		Remark:    req.Remark,
	}
	data.CreateBy = v[0]
	_, err = uc.d.CreateDictData(&data)
	if err != nil {
		return
	}
	return
}

// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body pb.DictData true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [put]
// @Security Bearer
func (uc *AdminUsecase) UpdateDictData(ctx context.Context, req *pb.DictData) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = ecode.Unknown("meta", "error")
		return
	}
	v := md.Get("UserID")
	if len(v) < 0 {
		err = ecode.Unknown("meta", "error")
		return
	}
	data := model.SysDictData{
		DictCode:  req.DictCode,
		DictSort:  req.DictSort,
		DictLabel: req.DictLabel,
		DictValue: req.DictValue,
		DictType:  req.DictType,
		CssClass:  req.CssClass,
		ListClass: req.ListClass,
		IsDefault: req.IsDefault,
		Status:    req.Status,
		Default:   req.Default,
		Remark:    req.Remark,
	}
	data.UpdateBy = v[0]
	_, err = uc.d.UpdateDictData(&data, int(data.DictCode))
	if err != nil {
		return
	}
	return
}

// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode path int true "dictCode"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/data/{dictCode} [delete]
func (uc *AdminUsecase) DeleteDictData(ctx context.Context, req *pb.DeleteDictDataRequest) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = ecode.Unknown("meta", "error")
		return
	}
	v := md.Get("UserID")
	if len(v) < 0 {
		err = ecode.Unknown("meta", "error")
		return
	}
	var data model.SysDictData
	data.UpdateBy = v[0]
	IDS, err := strings.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteDictData(&data, IDS)
	if err != nil {
		return
	}
	return
}
