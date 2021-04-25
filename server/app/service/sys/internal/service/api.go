package service

import (
	"context"

	pb "edu/api/sys/v1"
	"edu/service/sys/internal/biz"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

type ApiService struct {
	pb.UnimplementedApiServer
	uc *biz.AdminUsecase
}

func NewApiService(uc *biz.AdminUsecase) *ApiService {
	return &ApiService{
		uc: uc,
	}
}

func (s *ApiService) GetDictDataListByDictType(ctx context.Context, req *pb.GetDictDataListByDictTypeRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.GetDictDataListByDictType(ctx, req)
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
		Count:   int32(len(list)),
		Data:    list,
	}
	return
}

func (s *ApiService) GetDictData(ctx context.Context, req *pb.GetDictDataRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *ApiService) GetDictTypeList(ctx context.Context, req *pb.GetDictTypeListRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (s *ApiService) GetDictType(ctx context.Context, req *pb.GetDictTypeRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
