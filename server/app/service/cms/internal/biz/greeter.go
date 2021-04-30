package biz

import (
	"context"

	pb "edu/api/cms"
	uuidpb "edu/api/uuid"
	"edu/pkg/strings"
	"edu/service/cms/internal/dao"
	"edu/service/cms/internal/model"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GreeterUsecase struct {
	uuidc uuidpb.UUIDClient
	d     dao.Dao
	log   *log.Helper
}

func NewGreeterUsecase(logger log.Logger, d dao.Dao, r *registry.Registry) (*GreeterUsecase, error) {
	cc, err := uuidpb.NewUUID(context.Background(), grpc.WithDiscovery(r))
	if err != nil {
		return nil, err
	}
	return &GreeterUsecase{
		uuidc: cc,
		d:     d,
		log:   log.NewHelper("usecase/greeter", logger),
	}, nil
}

func (uc *GreeterUsecase) ListHelpCategoryPage(ctx context.Context, req *pb.ListHelpCategoryRequest) (page []*pb.HelpCategory, total int64, err error) {
	filter := model.CmsHelpCategory{}
	list, total, err := uc.d.GetCmsHelpCategoryPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.HelpCategory, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.HelpCategory{
			Id:         uint64(it.Id),
			Name:       it.Name,
			Icon:       it.Icon,
			HelpCount:  (it.HelpCount),
			ShowStatus: it.ShowStatus,
			Sort:       it.Sort,
			CreatedAt:  timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetHelpCategory(ctx context.Context, req *pb.GetHelpCategoryRequest) (reply *pb.HelpCategory, err error) {
	d, err := uc.d.GetCmsHelpCategory(ctx, &model.CmsHelpCategory{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.HelpCategory{
		Id:         d.Id,
		Name:       d.Name,
		Icon:       d.Icon,
		HelpCount:  d.HelpCount,
		ShowStatus: d.ShowStatus,
		Sort:       d.Sort,
		CreatedAt:  timestamppb.New(d.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateHelpCategory(ctx context.Context, req *pb.HelpCategory) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.CmsHelpCategory{
		Id:         id.Id,
		Name:       req.Name,
		Icon:       req.Icon,
		HelpCount:  req.HelpCount,
		ShowStatus: req.ShowStatus,
		Sort:       req.Sort,
	}
	_, err = uc.d.CreateCmsHelpCategory(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateHelpCategory(ctx context.Context, req *pb.HelpCategory) (err error) {
	d := model.CmsHelpCategory{
		Id:         req.Id,
		Name:       req.Name,
		Icon:       req.Icon,
		HelpCount:  req.HelpCount,
		ShowStatus: req.ShowStatus,
		Sort:       req.Sort,
	}
	_, err = uc.d.UpdateCmsHelpCategory(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteHelpCategory(ctx context.Context, req *pb.DeleteHelpCategoryRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteCmsHelpCategory(ctx, &model.CmsHelpCategory{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListHelpPage(ctx context.Context, req *pb.ListHelpRequest) (page []*pb.Help, total int64, err error) {
	filter := model.CmsHelp{}
	list, total, err := uc.d.GetCmsHelpPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Help, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Help{
			Id:         it.Id,
			CategoryId: it.CategoryId,
			Content:    it.Content,
			Icon:       it.Icon,
			ReadCount:  it.ReadCount,
			ShowStatus: it.ShowStatus,
			Title:      it.Title,
			CreatedAt:  timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetHelp(ctx context.Context, req *pb.GetHelpRequest) (reply *pb.Help, err error) {
	filter := model.CmsHelp{
		Id: uint64(req.HelpId),
	}
	it, err := uc.d.GetCmsHelp(ctx, &filter)
	if err != nil {
		return
	}
	reply = &pb.Help{
		Id:         it.Id,
		CategoryId: it.CategoryId,
		Content:    it.Content,
		Icon:       it.Icon,
		ReadCount:  it.ReadCount,
		ShowStatus: it.ShowStatus,
		Title:      it.Title,
		CreatedAt:  timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateHelp(ctx context.Context, req *pb.Help) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.CmsHelp{
		Id:         id.Id,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Content:    req.Content,
		Icon:       req.Icon,
		ReadCount:  req.ReadCount,
		ShowStatus: req.ShowStatus,
	}
	_, err = uc.d.CreateCmsHelp(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateHelp(ctx context.Context, req *pb.Help) (err error) {
	d := model.CmsHelp{
		Id:         req.Id,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Content:    req.Content,
		Icon:       req.Icon,
		ReadCount:  req.ReadCount,
		ShowStatus: req.ShowStatus,
	}
	_, err = uc.d.UpdateCmsHelp(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteHelp(ctx context.Context, req *pb.DeleteHelpRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteCmsHelp(ctx, &model.CmsHelp{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListSubjectPage(ctx context.Context, req *pb.ListSubjectRequest) (page []*pb.Subject, total int64, err error) {
	filter := model.CmsSubject{}
	list, total, err := uc.d.GetCmsSubjectPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Subject, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Subject{
			Id:           it.Id,
			CategoryId:   it.CategoryId,
			Title:        it.Title,
			Pic:          it.Pic,
			ProductCount: it.ProductCount,
			CreatedAt:    timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetSubject(ctx context.Context, req *pb.GetSubjectRequest) (reply *pb.Subject, err error) {
	filter := model.CmsSubject{
		Id: uint64(req.Id),
	}
	it, err := uc.d.GetCmsSubject(ctx, &filter)
	if err != nil {
		return
	}
	reply = &pb.Subject{
		Id:           it.Id,
		CategoryId:   it.CategoryId,
		Title:        it.Title,
		Pic:          it.Pic,
		ProductCount: it.ProductCount,
		CreatedAt:    timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateSubject(ctx context.Context, req *pb.Subject) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.CmsSubject{
		Id:           id.Id,
		CategoryId:   req.CategoryId,
		Title:        req.Title,
		Pic:          req.Pic,
		ProductCount: req.ProductCount,
	}
	_, err = uc.d.CreateCmsSubject(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateSubject(ctx context.Context, req *pb.Subject) (err error) {
	d := model.CmsSubject{
		Id:           req.Id,
		CategoryId:   req.CategoryId,
		Title:        req.Title,
		Pic:          req.Pic,
		ProductCount: req.ProductCount,
	}
	_, err = uc.d.UpdateCmsSubject(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteSubject(ctx context.Context, req *pb.DeleteSubjectRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteCmsSubject(ctx, &model.CmsSubject{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListSubjectCategoryPage(ctx context.Context, req *pb.ListSubjectCategoryRequest) (page []*pb.SubjectCategory, total int64, err error) {
	filter := model.CmsSubjectCategory{}
	list, total, err := uc.d.GetCmsSubjectCategoryPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.SubjectCategory, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.SubjectCategory{
			Id:           it.Id,
			Name:         it.Name,
			Icon:         it.Icon,
			SubjectCount: it.SubjectCount,
			ShowStatus:   it.ShowStatus,
			Sort:         it.Sort,
			CreatedAt:    timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetSubjectCategory(ctx context.Context, req *pb.GetSubjectCategoryRequest) (reply *pb.SubjectCategory, err error) {
	filter := model.CmsSubjectCategory{
		Id: uint64(req.Id),
	}
	it, err := uc.d.GetCmsSubjectCategory(ctx, &filter)
	if err != nil {
		return
	}
	reply = &pb.SubjectCategory{
		Id:           it.Id,
		Name:         it.Name,
		Icon:         it.Icon,
		SubjectCount: it.SubjectCount,
		ShowStatus:   it.ShowStatus,
		Sort:         it.Sort,
		CreatedAt:    timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateSubjectCategory(ctx context.Context, req *pb.SubjectCategory) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.CmsSubjectCategory{
		Id:           id.Id,
		Name:         req.Name,
		Icon:         req.Icon,
		SubjectCount: req.SubjectCount,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
	}
	_, err = uc.d.CreateCmsSubjectCategory(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateSubjectCategory(ctx context.Context, req *pb.SubjectCategory) (err error) {
	d := model.CmsSubjectCategory{
		Id:           req.Id,
		Name:         req.Name,
		Icon:         req.Icon,
		SubjectCount: req.SubjectCount,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
	}
	_, err = uc.d.UpdateCmsSubjectCategory(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteSubjectCategory(ctx context.Context, req *pb.DeleteSubjectCategoryRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteCmsSubjectCategory(ctx, &model.CmsSubjectCategory{}, ids)
	if err != nil {
		return
	}
	return
}
