package biz

import (
	"context"

	pb "edu/api/pms"
	uuidpb "edu/api/uuid"
	"edu/pkg/strings"
	"edu/service/pms/internal/dao"
	"edu/service/pms/internal/model"

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

func (uc *GreeterUsecase) ListAlbumPage(ctx context.Context, req *pb.ListAlbumRequest) (page []*pb.Album, total int64, err error) {
	filter := model.PmsAlbum{}
	list, total, err := uc.d.GetPmsAlbumPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Album, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Album{
			Id:          it.Id,
			Name:        it.Name,
			CoverPic:    it.CoverPic,
			PicCount:    it.PicCount,
			Sort:        it.Sort,
			Description: it.Description,
			CreatedAt:   timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (reply *pb.Album, err error) {
	it, err := uc.d.GetPmsAlbum(ctx, &model.PmsAlbum{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.Album{
		Id:          it.Id,
		Name:        it.Name,
		CoverPic:    it.CoverPic,
		PicCount:    it.PicCount,
		Sort:        it.Sort,
		Description: it.Description,
		CreatedAt:   timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateAlbum(ctx context.Context, req *pb.Album) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsAlbum{
		Id:          id.Id,
		Name:        req.Name,
		CoverPic:    req.CoverPic,
		PicCount:    req.PicCount,
		Description: req.Description,
		Sort:        req.Sort,
	}
	_, err = uc.d.CreatePmsAlbum(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateAlbum(ctx context.Context, req *pb.Album) (err error) {
	d := model.PmsAlbum{
		Id:          req.Id,
		Name:        req.Name,
		CoverPic:    req.CoverPic,
		PicCount:    req.PicCount,
		Description: req.Description,
		Sort:        req.Sort,
	}
	_, err = uc.d.UpdatePmsAlbum(ctx, &d, (req.Id))
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteAlbum(ctx context.Context, req *pb.DeleteAlbumRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsAlbum(ctx, &model.PmsAlbum{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListBrandPage(ctx context.Context, req *pb.ListBrandRequest) (page []*pb.Brand, total int64, err error) {
	filter := model.PmsBrand{}
	list, total, err := uc.d.GetPmsBrandPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Brand, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Brand{
			Id:                  it.Id,
			Name:                it.Name,
			FirstLetter:         it.FirstLetter,
			Sort:                it.Sort,
			FactoryStatus:       it.FactoryStatus,
			ShowStatus:          it.ShowStatus,
			ProductCount:        it.ProductCount,
			ProductCommentCount: it.ProductCommentCount,
			Logo:                it.Logo,
			BigPic:              it.BigPic,
			BrandStory:          it.BrandStory,
			CreatedAt:           timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetBrand(ctx context.Context, req *pb.GetBrandRequest) (reply *pb.Brand, err error) {
	it, err := uc.d.GetPmsBrand(ctx, &model.PmsBrand{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.Brand{
		Id:                  it.Id,
		Name:                it.Name,
		FirstLetter:         it.FirstLetter,
		Sort:                it.Sort,
		FactoryStatus:       it.FactoryStatus,
		ShowStatus:          it.ShowStatus,
		ProductCount:        it.ProductCount,
		ProductCommentCount: it.ProductCommentCount,
		Logo:                it.Logo,
		BigPic:              it.BigPic,
		BrandStory:          it.BrandStory,
		CreatedAt:           timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateBrand(ctx context.Context, req *pb.Brand) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsBrand{
		Id:                  id.Id,
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	}
	_, err = uc.d.CreatePmsBrand(ctx, &d)
	if err != nil {
		return
	}
	return
}
func (uc *GreeterUsecase) UpdateBrand(ctx context.Context, req *pb.Brand) (err error) {
	d := model.PmsBrand{
		Id:                  req.Id,
		Name:                req.Name,
		FirstLetter:         req.FirstLetter,
		Sort:                req.Sort,
		FactoryStatus:       req.FactoryStatus,
		ShowStatus:          req.ShowStatus,
		ProductCount:        req.ProductCount,
		ProductCommentCount: req.ProductCommentCount,
		Logo:                req.Logo,
		BigPic:              req.BigPic,
		BrandStory:          req.BrandStory,
	}
	_, err = uc.d.UpdatePmsBrand(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteBrand(ctx context.Context, req *pb.DeleteBrandRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsBrand(ctx, &model.PmsBrand{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) genProductCategoryTree(ctx context.Context, p *model.PmsProductCategory, list []model.PmsProductCategory) (doc *pb.ProductCategory, err error) {
	doc = &pb.ProductCategory{
		Id:           p.Id,
		ParentId:     p.ParentId,
		Name:         p.Name,
		Level:        p.Level,
		ProductCount: p.ProductCount,
		ProductUnit:  p.ProductUnit,
		NavStatus:    p.NavStatus,
		ShowStatus:   p.ShowStatus,
		Sort:         p.Sort,
		Icon:         p.Icon,
		Keywords:     p.Keywords,
		Description:  p.Description,
		CreatedAt:    timestamppb.New(p.CreatedAt),
	}

	children := make([]*pb.ProductCategory, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		if it.ParentId == p.Id {
			x, err := uc.genProductCategoryTree(ctx, &it, list)
			if err != nil {
				continue
			}
			children = append(children, x)
		}
	}
	doc.Children = children
	return
}

func (uc *GreeterUsecase) ListProductCategoryTree(ctx context.Context, req *pb.ListProductCategoryTreeRequest) (docs []*pb.ProductCategory, err error) {
	filter := model.PmsProductCategory{}
	list, err := uc.d.GetPmsProductCategoryList(ctx, &filter)
	if err != nil {
		return
	}
	docs = make([]*pb.ProductCategory, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		if it.ParentId == 0 {
			x, err := uc.genProductCategoryTree(ctx, &it, list)
			if err != nil {
				continue
			}
			docs = append(docs, x)
		}
	}
	return
}

func (uc *GreeterUsecase) GetProductCategory(ctx context.Context, req *pb.GetProductCategoryRequest) (reply *pb.ProductCategory, err error) {
	list, err := uc.d.GetPmsProductCategoryList(ctx, &model.PmsProductCategory{})
	if err != nil {
		return
	}
	for i := 0; i < len(list); i++ {
		it := list[i]
		if it.Id == req.Id {
			x, err1 := uc.genProductCategoryTree(ctx, &it, list)
			if err1 != nil {
				err = err1
				return
			}
			reply = x
			return
		}
	}
	return
}

func (uc *GreeterUsecase) CreateProductCategory(ctx context.Context, req *pb.ProductCategory) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsProductCategory{
		Id:           id.Id,
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: req.ProductCount,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.NavStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	}
	_, err = uc.d.CreatePmsProductCategory(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateProductCategory(ctx context.Context, req *pb.ProductCategory) (err error) {
	d := model.PmsProductCategory{
		Id:           req.Id,
		ParentId:     req.ParentId,
		Name:         req.Name,
		Level:        req.Level,
		ProductCount: req.ProductCount,
		ProductUnit:  req.ProductUnit,
		NavStatus:    req.NavStatus,
		ShowStatus:   req.ShowStatus,
		Sort:         req.Sort,
		Icon:         req.Icon,
		Keywords:     req.Keywords,
		Description:  req.Description,
	}
	_, err = uc.d.UpdatePmsProductCategory(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteProductCategory(ctx context.Context, req *pb.DeleteProductCategoryRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsProductCategory(ctx, &model.PmsProductCategory{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListProductPage(ctx context.Context, req *pb.ListProductRequest) (page []*pb.Product, total int64, err error) {
	filter := model.PmsProduct{}
	list, total, err := uc.d.GetPmsProductPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.Product, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.Product{
			Id:                         it.Id,
			BrandId:                    it.BrandId,
			ProductCategoryId:          it.ProductCategoryId,
			FeightTemplateId:           it.FeightTemplateId,
			ProductAttributeCategoryId: it.ProductAttributeCategoryId,
			Name:                       it.Name,
			Pic:                        it.Pic,
			ProductSn:                  it.ProductSn,
			PublishStatus:              it.PublishStatus,
			NewStatus:                  it.NewStatus,
			RecommandStatus:            it.RecommandStatus,
			VerifyStatus:               it.VerifyStatus,
			Sort:                       it.Sort,
			Sale:                       it.Sale,
			Price:                      it.Price,
			PromotionPrice:             it.PromotionPrice,
			GiftGrowth:                 it.GiftGrowth,
			GiftPoint:                  it.GiftPoint,
			UsePointLimit:              it.UsePointLimit,
			SubTitle:                   it.SubTitle,
			Description:                it.Description,
			OriginalPrice:              it.OriginalPrice,
			Stock:                      it.Stock,
			LowStock:                   it.LowStock,
			Unit:                       it.Unit,
			Weight:                     it.Weight,
			PreviewStatus:              it.PreviewStatus,
			ServiceIds:                 it.ServiceIds,
			Keywords:                   it.Keywords,
			Note:                       it.Note,
			AlbumPics:                  it.AlbumPics,
			DetailTitle:                it.DetailTitle,
			DetailDesc:                 it.DetailDesc,
			DetailHtml:                 it.DetailHtml,
			PromotionStartTime:         timestamppb.New(it.PromotionStartTime),
			PromotionEndTime:           timestamppb.New(it.PromotionEndTime),
			PromotionPerLimit:          it.PromotionPerLimit,
			PromotionType:              it.PromotionType,
			BrandName:                  it.BrandName,
			ProductCategoryName:        it.ProductCategoryName,
			CreatedAt:                  timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetProduct(ctx context.Context, req *pb.GetProductRequest) (reply *pb.Product, err error) {
	it, err := uc.d.GetPmsProduct(ctx, &model.PmsProduct{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.Product{
		Id:                         it.Id,
		BrandId:                    it.BrandId,
		ProductCategoryId:          it.ProductCategoryId,
		FeightTemplateId:           it.FeightTemplateId,
		ProductAttributeCategoryId: it.ProductAttributeCategoryId,
		Name:                       it.Name,
		Pic:                        it.Pic,
		ProductSn:                  it.ProductSn,
		PublishStatus:              it.PublishStatus,
		NewStatus:                  it.NewStatus,
		RecommandStatus:            it.RecommandStatus,
		VerifyStatus:               it.VerifyStatus,
		Sort:                       it.Sort,
		Sale:                       it.Sale,
		Price:                      it.Price,
		PromotionPrice:             it.PromotionPrice,
		GiftGrowth:                 it.GiftGrowth,
		GiftPoint:                  it.GiftPoint,
		UsePointLimit:              it.UsePointLimit,
		SubTitle:                   it.SubTitle,
		Description:                it.Description,
		OriginalPrice:              it.OriginalPrice,
		Stock:                      it.Stock,
		LowStock:                   it.LowStock,
		Unit:                       it.Unit,
		Weight:                     it.Weight,
		PreviewStatus:              it.PreviewStatus,
		ServiceIds:                 it.ServiceIds,
		Keywords:                   it.Keywords,
		Note:                       it.Note,
		AlbumPics:                  it.AlbumPics,
		DetailTitle:                it.DetailTitle,
		DetailDesc:                 it.DetailDesc,
		DetailHtml:                 it.DetailHtml,
		PromotionStartTime:         timestamppb.New(it.PromotionStartTime),
		PromotionEndTime:           timestamppb.New(it.PromotionEndTime),
		PromotionPerLimit:          it.PromotionPerLimit,
		PromotionType:              it.PromotionType,
		BrandName:                  it.BrandName,
		ProductCategoryName:        it.ProductCategoryName,
		CreatedAt:                  timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateProduct(ctx context.Context, req *pb.Product) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsProduct{
		Id:                         id.Id,
		BrandId:                    req.BrandId,
		ProductCategoryId:          req.ProductCategoryId,
		FeightTemplateId:           req.FeightTemplateId,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		Pic:                        req.Pic,
		ProductSn:                  req.ProductSn,
		PublishStatus:              req.PublishStatus,
		NewStatus:                  req.NewStatus,
		RecommandStatus:            req.RecommandStatus,
		VerifyStatus:               req.VerifyStatus,
		Sort:                       req.Sort,
		Sale:                       req.Sale,
		Price:                      req.Price,
		PromotionPrice:             req.PromotionPrice,
		GiftGrowth:                 req.GiftGrowth,
		GiftPoint:                  req.GiftPoint,
		UsePointLimit:              req.UsePointLimit,
		SubTitle:                   req.SubTitle,
		Description:                req.Description,
		OriginalPrice:              req.OriginalPrice,
		Stock:                      req.Stock,
		LowStock:                   req.LowStock,
		Unit:                       req.Unit,
		Weight:                     req.Weight,
		PreviewStatus:              req.PreviewStatus,
		ServiceIds:                 req.ServiceIds,
		Keywords:                   req.Keywords,
		Note:                       req.Note,
		AlbumPics:                  req.AlbumPics,
		DetailTitle:                req.DetailTitle,
		DetailDesc:                 req.DetailDesc,
		DetailHtml:                 req.DetailHtml,
		PromotionStartTime:         req.PromotionStartTime.AsTime(),
		PromotionEndTime:           req.PromotionEndTime.AsTime(),
		PromotionPerLimit:          req.PromotionPerLimit,
		PromotionType:              req.PromotionType,
		BrandName:                  req.BrandName,
		ProductCategoryName:        req.ProductCategoryName,
	}
	_, err = uc.d.CreatePmsProduct(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateProduct(ctx context.Context, req *pb.Product) (err error) {
	d := model.PmsProduct{
		Id:                         req.Id,
		BrandId:                    req.BrandId,
		ProductCategoryId:          req.ProductCategoryId,
		FeightTemplateId:           req.FeightTemplateId,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		Pic:                        req.Pic,
		ProductSn:                  req.ProductSn,
		PublishStatus:              req.PublishStatus,
		NewStatus:                  req.NewStatus,
		RecommandStatus:            req.RecommandStatus,
		VerifyStatus:               req.VerifyStatus,
		Sort:                       req.Sort,
		Sale:                       req.Sale,
		Price:                      req.Price,
		PromotionPrice:             req.PromotionPrice,
		GiftGrowth:                 req.GiftGrowth,
		GiftPoint:                  req.GiftPoint,
		UsePointLimit:              req.UsePointLimit,
		SubTitle:                   req.SubTitle,
		Description:                req.Description,
		OriginalPrice:              req.OriginalPrice,
		Stock:                      req.Stock,
		LowStock:                   req.LowStock,
		Unit:                       req.Unit,
		Weight:                     req.Weight,
		PreviewStatus:              req.PreviewStatus,
		ServiceIds:                 req.ServiceIds,
		Keywords:                   req.Keywords,
		Note:                       req.Note,
		AlbumPics:                  req.AlbumPics,
		DetailTitle:                req.DetailTitle,
		DetailDesc:                 req.DetailDesc,
		DetailHtml:                 req.DetailHtml,
		PromotionStartTime:         req.PromotionStartTime.AsTime(),
		PromotionEndTime:           req.PromotionEndTime.AsTime(),
		PromotionPerLimit:          req.PromotionPerLimit,
		PromotionType:              req.PromotionType,
		BrandName:                  req.BrandName,
		ProductCategoryName:        req.ProductCategoryName,
	}
	_, err = uc.d.UpdatePmsProduct(ctx, &d, req.Id)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsProduct(ctx, &model.PmsProduct{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListProductAttributePage(ctx context.Context, req *pb.ListProductAttributeRequest) (page []*pb.ProductAttribute, total int64, err error) {
	filter := model.PmsProductAttribute{}
	list, total, err := uc.d.GetPmsProductAttributePage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.ProductAttribute, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.ProductAttribute{
			Id:                         it.Id,
			ProductAttributeCategoryId: it.ProductAttributeCategoryId,
			Name:                       it.Name,
			SelectType:                 it.SelectType,
			InputType:                  it.InputType,
			InputList:                  it.InputList,
			Sort:                       it.Sort,
			FilterType:                 it.FilterType,
			SearchType:                 it.SearchType,
			RelatedStatus:              it.RelatedStatus,
			HandAddStatus:              it.HandAddStatus,
			Type:                       it.Type,
			CreatedAt:                  timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetProductAttribute(ctx context.Context, req *pb.GetProductAttributeRequest) (reply *pb.ProductAttribute, err error) {
	it, err := uc.d.GetPmsProductAttribute(ctx, &model.PmsProductAttribute{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.ProductAttribute{
		Id:                         it.Id,
		ProductAttributeCategoryId: it.ProductAttributeCategoryId,
		Name:                       it.Name,
		SelectType:                 it.SelectType,
		InputType:                  it.InputType,
		InputList:                  it.InputList,
		Sort:                       it.Sort,
		FilterType:                 it.FilterType,
		SearchType:                 it.SearchType,
		RelatedStatus:              it.RelatedStatus,
		HandAddStatus:              it.HandAddStatus,
		Type:                       it.Type,
		CreatedAt:                  timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateProductAttribute(ctx context.Context, req *pb.ProductAttribute) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsProductAttribute{
		Id:                         id.Id,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		SelectType:                 req.SelectType,
		InputType:                  req.InputType,
		InputList:                  req.InputList,
		Sort:                       req.Sort,
		FilterType:                 req.FilterType,
		SearchType:                 req.SearchType,
		RelatedStatus:              req.RelatedStatus,
		HandAddStatus:              req.HandAddStatus,
		Type:                       req.Type,
	}
	_, err = uc.d.CreatePmsProductAttribute(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateProductAttribute(ctx context.Context, req *pb.ProductAttribute) (err error) {
	d := model.PmsProductAttribute{
		Id:                         req.Id,
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		SelectType:                 req.SelectType,
		InputType:                  req.InputType,
		InputList:                  req.InputList,
		Sort:                       req.Sort,
		FilterType:                 req.FilterType,
		SearchType:                 req.SearchType,
		RelatedStatus:              req.RelatedStatus,
		HandAddStatus:              req.HandAddStatus,
		Type:                       req.Type,
	}
	_, err = uc.d.UpdatePmsProductAttribute(ctx, &d, (req.Id))
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteProductAttribute(ctx context.Context, req *pb.DeleteProductAttributeRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsProductAttribute(ctx, &model.PmsProductAttribute{}, ids)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) ListProductAttributeCategoryPage(ctx context.Context, req *pb.ListProductAttributeCategoryRequest) (page []*pb.ProductAttributeCategory, total int64, err error) {
	filter := model.PmsProductAttributeCategory{}
	list, total, err := uc.d.GetPmsProductAttributeCategoryPage(ctx, &filter, int(req.PageSize), int(req.PageIndex))
	if err != nil {
		return
	}
	page = make([]*pb.ProductAttributeCategory, 0)
	for i := 0; i < len(list); i++ {
		it := list[i]
		page = append(page, &pb.ProductAttributeCategory{
			Id:             it.Id,
			Name:           it.Name,
			AttributeCount: it.AttributeCount,
			ParamCount:     it.ParamCount,
			CreatedAt:      timestamppb.New(it.CreatedAt),
		})
	}
	return
}

func (uc *GreeterUsecase) GetProductAttributeCategory(ctx context.Context, req *pb.GetProductAttributeCategoryRequest) (reply *pb.ProductAttributeCategory, err error) {
	it, err := uc.d.GetPmsProductAttributeCategory(ctx, &model.PmsProductAttributeCategory{Id: req.Id})
	if err != nil {
		return
	}
	reply = &pb.ProductAttributeCategory{
		Id:             it.Id,
		Name:           it.Name,
		AttributeCount: it.AttributeCount,
		ParamCount:     it.ParamCount,
		CreatedAt:      timestamppb.New(it.CreatedAt),
	}
	return
}

func (uc *GreeterUsecase) CreateProductAttributeCategory(ctx context.Context, req *pb.ProductAttributeCategory) (err error) {
	id, err := uc.uuidc.GenID(ctx, &uuidpb.GenIDReq{})
	if err != nil {
		return
	}
	d := model.PmsProductAttributeCategory{
		Id:             id.Id,
		Name:           req.Name,
		AttributeCount: req.AttributeCount,
		ParamCount:     req.ParamCount,
	}
	_, err = uc.d.CreatePmsProductAttributeCategory(ctx, &d)
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) UpdateProductAttributeCategory(ctx context.Context, req *pb.ProductAttributeCategory) (err error) {
	d := model.PmsProductAttributeCategory{
		Id:             req.Id,
		Name:           req.Name,
		AttributeCount: req.AttributeCount,
		ParamCount:     req.ParamCount,
	}
	_, err = uc.d.UpdatePmsProductAttributeCategory(ctx, &d, (req.Id))
	if err != nil {
		return
	}
	return
}

func (uc *GreeterUsecase) DeleteProductAttributeCategory(ctx context.Context, req *pb.DeleteProductAttributeCategoryRequest) (err error) {
	ids, err := strings.SplitUint64s(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeletePmsProductAttributeCategory(ctx, &model.PmsProductAttributeCategory{}, ids)
	if err != nil {
		return
	}
	return
}
