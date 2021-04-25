package service

import (
	"context"

	pb "edu/api/sys/v1"

	"edu/service/sys/internal/model"
	"edu/service/sys/internal/tools"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / Tools
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /admin/tools/v1/db/tables/page [get]
func (s *AdminService) GetDBTableList(ctx context.Context, req *pb.GetDBTableListRequest) (reply *pb.ApiReply, err error) {
	var data model.DBTables
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	data.TableName = req.TableName
	d, err := tools.New(s.cfg, s.logger)
	if err != nil {
		return
	}
	defer d.Close()

	result, count, err := d.GetDBTablesPage(s.cfg.Gen.Dbname, &data, pageSize, pageIndex)
	if err != nil {
		return
	}
	s.log.Infof("GetDBTablesPage: %v", count)
	list := make([]*anypb.Any, 0)
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
