package service

import (
	"context"

	pb "edu/api/cms"
	"edu/service/cms/internal/biz"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc *biz.GreeterUsecase
}

func NewAdminService(uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{
		uc: uc,
	}
}

func (s *AdminService) ListHelp(ctx context.Context, req *pb.ListHelpRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListHelpPage(ctx, req)
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
		Count:   int32(count),
		Data:    list,
	}
	return
}
func (s *AdminService) GetHelp(ctx context.Context, req *pb.GetHelpRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetHelp(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
		Data:    list,
	}
	return
}

func (s *AdminService) CreateHelp(ctx context.Context, req *pb.Help) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateHelp(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateHelp(ctx context.Context, req *pb.Help) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateHelp(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) DeleteHelp(ctx context.Context, req *pb.DeleteHelpRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteHelp(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) ListHelpCategory(ctx context.Context, req *pb.ListHelpCategoryRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListHelpCategoryPage(ctx, req)
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
		Count:   int32(count),
		Data:    list,
	}
	return
}

func (s *AdminService) GetHelpCategory(ctx context.Context, req *pb.GetHelpCategoryRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetHelpCategory(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
		Data:    list,
	}
	return
}

func (s *AdminService) CreateHelpCategory(ctx context.Context, req *pb.HelpCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateHelpCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateHelpCategory(ctx context.Context, req *pb.HelpCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateHelpCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) DeleteHelpCategory(ctx context.Context, req *pb.DeleteHelpCategoryRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteHelpCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) ListSubject(ctx context.Context, req *pb.ListSubjectRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListSubjectPage(ctx, req)
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
		Count:   int32(count),
		Data:    list,
	}
	return &pb.ApiReply{}, nil
}

func (s *AdminService) GetSubject(ctx context.Context, req *pb.GetSubjectRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetSubject(ctx, req)
	if err != nil {
		return
	}
	list := make([]*anypb.Any, 0)
	any, err1 := ptypes.MarshalAny(it)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
		Data:    list,
	}
	return
}

func (s *AdminService) CreateSubject(ctx context.Context, req *pb.Subject) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateSubject(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateSubject(ctx context.Context, req *pb.Subject) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateSubject(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteSubject(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}
