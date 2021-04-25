package biz

import (
	"context"
	"fmt"

	ssopb "edu/api/sso"
	pb "edu/api/sys/v1"
	"edu/pkg/meta"
	"edu/service/sys/internal/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertBoolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func convertMenu(menu model.Menu) (pbs *pb.Menu) {
	pbs = &pb.Menu{
		MenuId:     int32(menu.MenuId),
		MenuName:   menu.MenuName,
		Title:      menu.Title,
		Icon:       menu.Icon,
		Path:       menu.Path,
		Paths:      menu.Paths,
		MenuType:   menu.MenuType,
		Action:     menu.Action,
		Permission: menu.Permission,
		ParentId:   int32(menu.ParentId),
		NoCache:    menu.NoCache,
		Breadcrumb: menu.Breadcrumb,
		Component:  menu.Component,
		Sort:       int32(menu.Sort),
		Visible:    convertBoolStr(menu.Visible),
		IsFrame:    convertBoolStr(menu.IsFrame),
		CreatedAt:  timestamppb.New(menu.CreatedAt),
	}
	return
}

func diguiMenu(menulist []model.Menu, menu model.Menu) *pb.Menu {
	d := convertMenu(menu)
	min := make([]*pb.Menu, 0)
	for j := 0; j < len(menulist); j++ {
		it := menulist[j]
		if menu.MenuId != it.ParentId {
			continue
		}
		if it.MenuType != "F" {
			ms := diguiMenu(menulist, it)
			min = append(min, ms)
		} else {
			mi := convertMenu(it)
			min = append(min, mi)
		}
	}
	d.Children = min
	return d
}

func diguiMenuLable(menulist []model.Menu, menu model.Menu) *pb.MenuLable {
	d := &pb.MenuLable{
		Id:    menu.MenuId,
		Label: menu.Title,
	}
	min := make([]*pb.MenuLable, 0)
	for j := 0; j < len(menulist); j++ {
		it := menulist[j]
		if menu.MenuId != it.ParentId {
			continue
		}
		if it.MenuType != "F" {
			ms := diguiMenuLable(menulist, it)
			min = append(min, ms)
		} else {
			mi := &pb.MenuLable{
				Id:    it.MenuId,
				Label: it.Title,
			}
			min = append(min, mi)
		}

	}
	d.Children = min
	return d
}

func (uc *AdminUsecase) ListMenu(ctx context.Context, token string, req *pb.ListMenuRequest) (list []*pb.Menu, err error) {
	// out, err := uc.mw.ValidationToken(token)
	// if err != nil {
	// 	return
	// }
	// dp := out.(*ssopb.DataPermission)

	// var pageSize = int(req.PageSize)
	// var pageIndex = int(req.PageIndex)

	var menu model.Menu
	menu.Visible = req.Visible
	menu.Title = req.Title
	result, err := uc.d.GetMenuPage(&menu)
	if err != nil {
		return
	}
	list = make([]*pb.Menu, 0)
	for i := 0; i < len(result); i++ {
		it := result[i]
		if it.ParentId != 0 {
			continue
		}
		menusInfo := diguiMenu(result, it)
		list = append(list, menusInfo)
	}
	return
}

func (uc *AdminUsecase) ListMenuByRole(ctx context.Context, e *model.Menu, rolename string) (list []*pb.Menu, err error) {
	token, err := meta.GetToken(ctx)
	if err != nil {
		return
	}
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	menulist, err := uc.d.GetMenusByRoleName(e, dp.RoleKey)
	list = make([]*pb.Menu, 0)
	for i := 0; i < len(menulist); i++ {
		it := menulist[i]
		if it.ParentId != 0 {
			continue
		}
		menusInfo := diguiMenu(menulist, it)
		list = append(list, menusInfo)
	}
	return
}

func (uc *AdminUsecase) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (mu *pb.Menu, err error) {
	var data model.Menu
	data.MenuId = req.Id
	it, err := uc.d.GetMenuByMenuId(&data)
	if err != nil {
		return
	}
	mu = convertMenu(it)
	return
}

func (uc *AdminUsecase) CreateMenu(ctx context.Context, tok string, req *pb.Menu) (err error) {
	out, err := uc.mw.ValidationToken(tok)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	data := model.Menu{
		MenuId:     req.MenuId,
		MenuName:   req.MenuName,
		Title:      req.Title,
		Icon:       req.Icon,
		Path:       req.Path,
		Paths:      req.Paths,
		MenuType:   req.MenuType,
		Action:     req.Action,
		Permission: req.Permission,
		ParentId:   req.ParentId,
		NoCache:    req.NoCache,
		Breadcrumb: req.Breadcrumb,
		Component:  req.Component,
		Sort:       req.Sort,
	}
	if req.Visible == "true" {
		data.Visible = true
	}
	if req.IsFrame == "true" {
		data.IsFrame = true
	}
	data.CreateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.CreateMenu(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) UpdateMenu(ctx context.Context, token string, req *pb.Menu) (err error) {
	out, err := uc.mw.ValidationToken(token)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	data := model.Menu{
		MenuId:     req.MenuId,
		MenuName:   req.MenuName,
		Title:      req.Title,
		Icon:       req.Icon,
		Path:       req.Path,
		Paths:      req.Paths,
		MenuType:   req.MenuType,
		Action:     req.Action,
		Permission: req.Permission,
		ParentId:   req.ParentId,
		NoCache:    req.NoCache,
		Breadcrumb: req.Breadcrumb,
		Component:  req.Component,
		Sort:       req.Sort,
		UpdateBy:   fmt.Sprint(dp.UserId),
	}
	if req.Visible == "1" {
		data.Visible = true
	}
	if req.IsFrame == "1" {
		data.IsFrame = true
	}
	_, err = uc.d.UpdateMenu(&data, int(data.MenuId))
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) ListMenuLable(e *model.Menu) (m []*pb.MenuLable, err error) {
	menulist, err := uc.d.GetMenus(e)
	m = make([]*pb.MenuLable, 0)
	for i := 0; i < len(menulist); i++ {
		it := menulist[i]
		if it.ParentId != 0 {
			continue
		}
		menusInfo := diguiMenuLable(menulist, it)
		m = append(m, menusInfo)
	}
	return
}

func (uc *AdminUsecase) DeleteMenu(ctx context.Context, req *pb.DeleteMenuRequest) (err error) {
	// _, err = uc.d.DeleteMenu(&model.Menu{}, int(req.Id))
	// if err != nil {
	// 	return
	// }
	return
}
