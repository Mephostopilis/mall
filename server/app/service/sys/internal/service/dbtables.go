package service

import (
	"context"

	pb "edu/api/sys/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
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
	page, total, err := s.uc.GetDBTableList(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(page); i++ {
		it := page[i]
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
		Count:   int32(total),
		Data:    list,
	}
	return
}
