package service

import (
	"context"

	pb "edu/api/tiku"
	"edu/pkg/meta"
	"edu/service/tiku/internal/biz"
	"edu/service/tiku/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.TikuUsecase
}

func NewAdminService(logger log.Logger, uc *biz.TikuUsecase) *AdminService {
	log := log.NewHelper("service", logger)
	s := &AdminService{
		uc:  uc,
		log: log,
	}
	return s
}

func (s *AdminService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	s.log.Infof("hello url %s", req.Name)
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	return
}

func (s *AdminService) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{}, nil
}

func (s *AdminService) GetChoiceList(ctx context.Context, req *pb.GetChoiceListRequest) (reply *pb.ApiReply, err error) {
	choices, cnt, err := s.uc.GetChoicesPage(ctx, &model.Choices{}, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(choices); i++ {
		it := choices[i]
		any, err1 := ptypes.MarshalAny(it)
		if err1 != nil {
			err = err1
			return
		}
		list = append(list, any)
	}
	reply = &pb.ApiReply{
		Data:  list,
		Count: int32(cnt),
	}
	return
}

func (s *AdminService) GetChoice(ctx context.Context, req *pb.GetChoiceRequest) (reply *pb.ApiReply, err error) {
	d, err := s.uc.GetChoice(ctx, req.Id)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(d)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Data:  list,
		Count: int32(1),
	}
	return
}

func (s *AdminService) InsertChoice(ctx context.Context, req *pb.Choice) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	err = s.uc.CreateChoices(ctx, token, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) UpdateChoice(ctx context.Context, req *pb.Choice) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	err = s.uc.UpdateChoice(ctx, token, req)
	if err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteChoice(ctx context.Context, req *pb.DeleteChoiceRequest) (reply *pb.ApiReply, err error) {
	return
}

func (s *AdminService) GetExerciseList(ctx context.Context, req *pb.GetExerciseListRequest) (reply *pb.ApiReply, err error) {
	exercises, cnt, err := s.uc.GetExercisePage(ctx, req, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	for i := 0; i < len(exercises); i++ {
		it := exercises[i]
		any, err1 := ptypes.MarshalAny(it)
		if err1 != nil {
			err = err1
			return
		}
		list = append(list, any)
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
		Count:  int32(cnt),
		Data:   list,
	}
	return
}

func (s *AdminService) GetExercise(ctx context.Context, req *pb.GetExerciseRequest) (reply *pb.ApiReply, err error) {
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) InsertExercise(ctx context.Context, req *pb.Exercise) (reply *pb.ApiReply, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	if err = s.uc.CreateExercise(ctx, token, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) UpdateExercise(ctx context.Context, req *pb.Exercise) (reply *pb.ApiReply, err error) {
	return
}

func (s *AdminService) DeleteExercise(ctx context.Context, req *pb.DeleteExerciseRequest) (reply *pb.ApiReply, err error) {
	return
}
