package biz

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/model"
)

func (uc *AdminUsecase) ListResource(ctx context.Context, req *pb.ListResourceRequest) (docs []*pb.SysResource, total int64, err error) {
	var pageSize = int(req.PageSize)
	var pageIndex = int(req.PageIndex)
	var data model.SysResource
	list, total, err := uc.d.GetSysResourcePage(ctx, &data, pageSize, pageIndex)
	if err != nil {
		return
	}
	docs = make([]*pb.SysResource, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		d := &pb.SysResource{
			Id: it.Id,
			V1: it.V1,
			V2: it.V2,
			V3: it.V3,
			V4: it.V4,
			V5: it.V5,
		}
		docs = append(docs, d)
	}
	return
}
