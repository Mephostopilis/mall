package biz

import (
	"context"

	pb "edu/api/sys/v1"

	"edu/service/sys/internal/model"
	"edu/service/sys/internal/tools"
)

func (uc *AdminUsecase) GetDBColumnList(ctx context.Context, req *pb.GetDBColumnListRequest) (page []*pb.DBColumn, total int64, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var filter model.DBColumns
	filter.TableName = req.TableName
	d, err := tools.New(uc.cfg, uc.logger)
	if err != nil {
		return
	}
	defer d.Close()
	list, total, err := d.GetDBColumnsPage(uc.cfg.Gen.Dbname, &filter, pageSize, pageIndex)
	if err != nil {
		return
	}
	page = make([]*pb.DBColumn, 0)
	for i := 0; i < len(list); i++ {
		// it := list[i]
		page = append(page, &pb.DBColumn{})
	}
	return
}
