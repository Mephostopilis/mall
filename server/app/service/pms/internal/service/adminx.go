package service

import (
	"context"

	pb "edu/api/pms"
	"edu/service/pms/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.GreeterUsecase
}

func NewAdminService(logger log.Logger, uc *biz.GreeterUsecase) *AdminService {
	return &AdminService{
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
		uc:  uc,
	}
}

func (s *AdminService) ListAlbum(ctx context.Context, req *pb.ListAlbumRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListAlbumPage(ctx, req)
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

func (s *AdminService) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetAlbum(ctx, req)
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

func (s *AdminService) CreateAlbum(ctx context.Context, req *pb.Album) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateAlbum(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateAlbum(ctx context.Context, req *pb.Album) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateAlbum(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteAlbum(ctx context.Context, req *pb.DeleteAlbumRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteAlbum(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) ListBrand(ctx context.Context, req *pb.ListBrandRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListBrandPage(ctx, req)
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

func (s *AdminService) GetBrand(ctx context.Context, req *pb.GetBrandRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetBrand(ctx, req)
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

func (s *AdminService) CreateBrand(ctx context.Context, req *pb.Brand) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateBrand(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateBrand(ctx context.Context, req *pb.Brand) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateBrand(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteBrand(ctx context.Context, req *pb.DeleteBrandRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteBrand(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) ListProductCategory(context.Context, *pb.ListProductCategoryRequest) (reply *pb.ApiReply, err error) {
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) ListProductCategoryTree(ctx context.Context, req *pb.ListProductCategoryTreeRequest) (reply *pb.ApiReply, err error) {
	result, err := s.uc.ListProductCategoryTree(ctx, req)
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
		Data:    list,
	}
	return
}

func (s *AdminService) GetProductCategory(ctx context.Context, req *pb.GetProductCategoryRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetProductCategory(ctx, req)
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

func (s *AdminService) CreateProductCategory(ctx context.Context, req *pb.ProductCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateProductCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateProductCategory(ctx context.Context, req *pb.ProductCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateProductCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) DeleteProductCategory(ctx context.Context, req *pb.DeleteProductCategoryRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteProductCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListProductPage(ctx, req)
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

func (s *AdminService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetProduct(ctx, req)
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

func (s *AdminService) CreateProduct(ctx context.Context, req *pb.Product) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateProduct(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateProduct(ctx context.Context, req *pb.Product) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateProduct(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteProduct(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) ListProductAttribute(ctx context.Context, req *pb.ListProductAttributeRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListProductAttributePage(ctx, req)
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

func (s *AdminService) GetProductAttribute(ctx context.Context, req *pb.GetProductAttributeRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetProductAttribute(ctx, req)
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

func (s *AdminService) CreateProductAttribute(ctx context.Context, req *pb.ProductAttribute) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateProductAttribute(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateProductAttribute(ctx context.Context, req *pb.ProductAttribute) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateProductAttribute(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteProductAttribute(ctx context.Context, req *pb.DeleteProductAttributeRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteProductAttribute(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) ListProductAttributeCategory(ctx context.Context, req *pb.ListProductAttributeCategoryRequest) (reply *pb.ApiReply, err error) {
	result, count, err := s.uc.ListProductAttributeCategoryPage(ctx, req)
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

func (s *AdminService) GetProductAttributeCategory(ctx context.Context, req *pb.GetProductAttributeCategoryRequest) (reply *pb.ApiReply, err error) {
	it, err := s.uc.GetProductAttributeCategory(ctx, req)
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

func (s *AdminService) CreateProductAttributeCategory(ctx context.Context, req *pb.ProductAttributeCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.CreateProductAttributeCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
	}
	return
}

func (s *AdminService) UpdateProductAttributeCategory(ctx context.Context, req *pb.ProductAttributeCategory) (reply *pb.ApiReply, err error) {
	if err = s.uc.UpdateProductAttributeCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}

func (s *AdminService) DeleteProductAttributeCategory(ctx context.Context, req *pb.DeleteProductAttributeCategoryRequest) (reply *pb.ApiReply, err error) {
	if err = s.uc.DeleteProductAttributeCategory(ctx, req); err != nil {
		return
	}
	reply = &pb.ApiReply{
		Code:   0,
		Reason: "OK",
	}
	return
}
