package service

import (
	"context"

	pb "edu/api/member"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	pb.UnimplementedApiServer
	log *log.Helper
}

func NewApiService(logger log.Logger) *ApiService {
	return &ApiService{
		log: log.NewHelper(log.With(logger, "module", "service/api")),
	}
}

func (s *ApiService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{}, nil
}
func (s *ApiService) ListMember(ctx context.Context, req *pb.ListMemberRequest) (*pb.ListMemberReply, error) {
	return &pb.ListMemberReply{}, nil
}
func (s *ApiService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.GetMemberReply, error) {
	return &pb.GetMemberReply{}, nil
}
func (s *ApiService) CreateMember(ctx context.Context, req *pb.Member) (*pb.CreateMemberReply, error) {
	return &pb.CreateMemberReply{}, nil
}
