package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
	"edu/service/sys/internal/tools"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *AdminUsecase) GetDBTableList(ctx context.Context, req *pb.GetDBTableListRequest) (page []*pb.DBTable, total int64, err error) {
	d, err := tools.New(uc.cfg, uc.logger)
	if err != nil {
		return
	}
	defer d.Close()
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var filter model.DBTables
	filter.TableName = req.TableName
	result, total, err := d.GetDBTablesPage(uc.cfg.Gen.Dbname, &filter, pageSize, pageIndex)
	if err != nil {
		return
	}
	uc.log.Infof("GetDBTablesPage: %v", total)
	page = make([]*pb.DBTable, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		d := &pb.DBTable{
			TableName:      it.TableName,
			Engine:         it.Engine,
			TableRows:      it.TableRows,
			TableCollation: it.TableCollation,
			TableComment:   it.TableComment,
			CreateTime:     timestamppb.New(it.CreateTime),
			UpdateTime:     timestamppb.New(it.UpdateTime),
		}
		page = append(page, d)
	}
	return
}
