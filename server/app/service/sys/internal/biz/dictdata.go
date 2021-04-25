package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"

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
			DictCode:  int64(it.DictCode),
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
			DictCode:  int64(it.DictCode),
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
