package biz

import (
	"context"

	ssopb "edu/api/sso"
	pb "edu/api/tiku"
	uuidpb "edu/api/uuid"
	"edu/pkg/jwtauth"
	"edu/service/tiku/internal/conf"
	"edu/service/tiku/internal/dao"
	"edu/service/tiku/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TikuUsecase struct {
	d     dao.Dao
	uuidc uuidpb.UUIDClient
	mw    *jwtauth.GinJWTMiddleware
	log   *log.Helper
}

func NewTikuUsecase(c *conf.App, logger log.Logger, d dao.Dao) (*TikuUsecase, error) {
	log := log.NewHelper("usecase/tiku", logger)
	uuidc, err := uuidpb.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	mw, err := jwtauth.New(
		&jwtauth.Jwt{Secret: c.Jwt.Secret, Timeout: c.Jwt.Timeout.AsDuration()},
		logger,
		jwtauth.DataPermissionToClaimsFunc,
		jwtauth.ClaimsToDataPermissionFunc)
	if err != nil {
		return nil, err
	}
	return &TikuUsecase{
		d:     d,
		uuidc: uuidc,
		mw:    mw,
		log:   log}, nil
}

func (uc *TikuUsecase) GetChoicesPage(ctx context.Context, e *model.Choices, pageSize int, pageIndex int) (docs []*pb.Choice, total int64, err error) {
	list, total, err := uc.d.GetChoicesPage(ctx, e, pageSize, pageIndex)
	if err != nil {
		return
	}
	docs = make([]*pb.Choice, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		d := &pb.Choice{
			Id:        it.Id,
			Ty:        it.Ty,
			Level:     it.Level,
			Pics:      it.Pics,
			Title:     it.Title,
			Answer:    it.Answer,
			CreatedAt: timestamppb.New(it.CreatedAt),
		}
		docs = append(docs, d)
	}
	return
}

func (uc *TikuUsecase) GetChoice(ctx context.Context, id uint64) (reply *pb.Choice, err error) {
	d, err := uc.d.GetChoices(ctx, &model.Choices{Id: id})
	if err != nil {
		return
	}
	reply = &pb.Choice{
		Id:        d.Id,
		Ty:        d.Ty,
		Level:     d.Level,
		Pics:      d.Pics,
		Title:     d.Title,
		Answer:    d.Answer,
		CreatedAt: timestamppb.New(d.CreatedAt),
	}
	return
}

func (uc *TikuUsecase) CreateChoices(ctx context.Context, token string, req *pb.Choice) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	uc.log.Debugf("create id = %d", id.Id)
	e := &model.Choices{
		Id:       id.Id,
		Level:    req.Level,
		Pics:     req.Pics,
		Title:    req.Title,
		Ty:       req.Ty,
		Answer:   req.Answer,
		CreateBy: dp.UserId,
	}
	_, err = uc.d.CreateChoices(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (uc *TikuUsecase) UpdateChoice(ctx context.Context, token string, req *pb.Choice) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	ud := &model.Choices{
		Id:       req.Id,
		Ty:       req.Ty,
		Level:    req.Level,
		Title:    req.Title,
		Pics:     req.Pics,
		Answer:   req.Answer,
		UpdateBy: dp.UserId,
	}
	_, err = uc.d.UpdateChoices(ctx, ud, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *TikuUsecase) GetUserChoice(ctx context.Context, id uint64) (*model.Choices, *model.UserChoices, error) {
	choices, err := uc.d.GetChoices(ctx, &model.Choices{Id: id})
	if err != nil {
		return nil, nil, err
	}
	userChoices, err := uc.d.GetUserChoices(ctx, &model.UserChoices{Id: id})
	if err != nil {
		return nil, nil, err
	}
	return &choices, &userChoices, nil
}

func (uc *TikuUsecase) GetExercisePage(ctx context.Context, req *pb.GetExerciseListRequest, pageSize int, pageIndex int) (docs []*pb.Exercise, total int64, err error) {
	filter := &model.Exercise{
		Ty:    req.Ty,
		Level: req.Level,
	}
	list, total, err := uc.d.GetExercisePage(ctx, filter, pageSize, pageIndex)
	if err != nil {
		return
	}
	docs = make([]*pb.Exercise, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		d := &pb.Exercise{
			Id:        it.Id,
			Ty:        it.Ty,
			Level:     it.Level,
			Title:     it.Title,
			Pics:      it.Pics,
			Answer:    it.Answer,
			CreatedAt: timestamppb.New(it.CreatedAt),
		}
		docs = append(docs, d)
	}
	return
}

func (uc *TikuUsecase) CreateExercise(ctx context.Context, token string, req *pb.Exercise) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return err
	}
	e := &model.Exercise{
		Id:       id.Id,
		Ty:       req.Ty,
		Level:    req.Level,
		Title:    req.Title,
		Pics:     req.Pics,
		Answer:   req.Answer,
		CreateBy: dp.UserId,
	}
	_, err = uc.d.CreateExercise(ctx, e)
	if err != nil {
		return err
	}
	return
}
