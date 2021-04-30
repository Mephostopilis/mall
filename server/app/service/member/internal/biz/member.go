package biz

import (
	"context"

	pb "edu/api/member"
	uuidpb "edu/api/uuid"
	"edu/service/member/internal/dao"
	"edu/service/member/internal/model"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MemberUsecase struct {
	uuidc uuidpb.UUIDClient
	d     dao.Dao
	log   *log.Helper
}

func NewMemberUsecase(logger log.Logger, d dao.Dao, r *registry.Registry) (*MemberUsecase, error) {
	cc, err := uuidpb.NewUUID(context.Background(), grpc.WithDiscovery(r))
	if err != nil {
		return nil, err
	}
	return &MemberUsecase{
		uuidc: cc,
		d:     d,
		log:   log.NewHelper("usecase/greeter", logger),
	}, nil
}

func (uc *MemberUsecase) ListMemberPage(ctx context.Context, req *pb.ListMemberRequest) (page []*pb.Member, total int64, err error) {
	filter := model.Member{}
	list, total, err := uc.d.GetMemberPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Member, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Member{
			Id:                 uint64(it.Id),
			MemberLevelId:      it.MemberLevelId,
			SourceType:         it.SourceType,
			Integration:        it.Integration,
			Growth:             it.Growth,
			LuckeyCount:        it.LuckeyCount,
			HistoryIntegration: it.HistoryIntegration,
			CreatedAt:          timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *MemberUsecase) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) CreateMember(ctx context.Context, req *pb.Member) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) UpdateMember(ctx context.Context, req *pb.Member) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) ListMemberAssets(ctx context.Context, req *pb.ListMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) GetMemberAssets(ctx context.Context, req *pb.GetMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) CreateMemberAssets(ctx context.Context, req *pb.MemberAssets) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) UpdateMemberAssets(ctx context.Context, req *pb.MemberAssets) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
func (uc *MemberUsecase) DeleteMemberAssets(ctx context.Context, req *pb.DeleteMemberAssetsRequest) (*pb.ApiReply, error) {
	return &pb.ApiReply{}, nil
}
