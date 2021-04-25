package service

import (
	"context"

	pb "edu/api/tiku"
	"edu/service/tiku/internal/biz"
	"edu/service/tiku/internal/model"

	"github.com/go-kratos/kratos/v2/log"
)

type ApiService struct {
	pb.UnimplementedApiServer
	log *log.Helper
	uc  *biz.TikuUsecase
}

func NewApiService(logger log.Logger, uc *biz.TikuUsecase) *ApiService {
	log := log.NewHelper("service", logger)
	s := &ApiService{
		uc:  uc,
		log: log,
	}
	return s
}

func (s *ApiService) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	s.log.Infof("hello url %s", req.Name)
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	return
}

func (s *ApiService) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{}, nil
}

func (s *ApiService) GetChoiceList(ctx context.Context, req *pb.GetChoiceListRequest) (reply *pb.GetChoiceListReply, err error) {
	choices, cnt, err := s.uc.GetChoicesPage(ctx, &model.Choices{}, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	list := make([]*pb.Choice, 0)
	for i := 0; i < len(choices); i++ {
		it := choices[i]
		d := &pb.Choice{
			Id: it.Id,
		}
		list = append(list, d)
	}
	reply = &pb.GetChoiceListReply{
		List:  list,
		Total: int32(cnt),
	}
	return
}

func (s *ApiService) GetChoice(ctx context.Context, req *pb.GetChoiceRequest) (*pb.Choice, error) {
	return &pb.Choice{}, nil
}

func (s *ApiService) AnswerChoice(ctx context.Context, req *pb.AnswerChoiceRequest) (*pb.AnswerChoiceReply, error) {
	return &pb.AnswerChoiceReply{}, nil
}

func (s *ApiService) GetExerciseList(ctx context.Context, req *pb.GetExerciseListRequest) (reply *pb.GetExerciseListReply, err error) {
	exercises, cnt, err := s.uc.GetExercisePage(ctx, req, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	reply = &pb.GetExerciseListReply{
		Total: int32(cnt),
		List:  exercises,
	}
	return
}

func (s *ApiService) GetExercise(ctx context.Context, req *pb.GetExerciseRequest) (*pb.Exercise, error) {
	return &pb.Exercise{}, nil
}
func (s *ApiService) AnswerExercise(ctx context.Context, req *pb.AnswerExerciseRequest) (*pb.AnswerExerciseReply, error) {
	return &pb.AnswerExerciseReply{}, nil
}
