package biz

import (
	"bytes"
	"context"
	"path"
	"text/template"

	pb "edu/api/sys/v1"
	"edu/pkg/tools"
	"edu/service/sys/internal/model"
)

// GenCode 生成代码工具
func (uc *AdminUsecase) GenCode(ctx context.Context, req *pb.GenCodeRequest) (err error) {
	table := model.SysTables{}
	table.TableId = (req.TableId)
	tab, err := uc.d.GetSysTables(&table)
	if err != nil {
		return
	}
	tab.Module = uc.cfg.Gen.Name
	t1, err := template.ParseFiles(uc.cfg.Gen.Template.Model)
	if err != nil {
		return
	}
	t2, err := template.ParseFiles(uc.cfg.Gen.Template.Dao)
	if err != nil {
		return
	}
	t3, err := template.ParseFiles(uc.cfg.Gen.Template.Js)
	if err != nil {
		return
	}
	t4, err := template.ParseFiles(uc.cfg.Gen.Template.Vue)
	if err != nil {
		return
	}
	tools.PathCreate(path.Join(uc.cfg.Gen.Backpath, "model"))
	tools.PathCreate(path.Join(uc.cfg.Gen.Backpath, "dao"))
	tools.PathCreate(path.Join(uc.cfg.Gen.Frontpath, "api", uc.cfg.Gen.Name))
	tools.PathCreate(path.Join(uc.cfg.Gen.Frontpath, "views", uc.cfg.Gen.Name, tab.PackageName))

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)

	tools.FileCreate(b1, path.Join(uc.cfg.Gen.Backpath, "model", tab.PackageName+".go"))
	tools.FileCreate(b2, path.Join(uc.cfg.Gen.Backpath, "dao", tab.PackageName+".go"))
	tools.FileCreate(b3, path.Join(uc.cfg.Gen.Frontpath, "api", uc.cfg.Gen.Name, tab.PackageName+".js"))
	tools.FileCreate(b4, path.Join(uc.cfg.Gen.Frontpath, "views", uc.cfg.Gen.Name, tab.PackageName, "index.vue"))

	return
}

func (uc *AdminUsecase) Preview(ctx context.Context, req *pb.PreviewRequest) (reply *pb.PreviewReply, err error) {
	t1, err := template.ParseFiles(uc.cfg.Gen.Template.Model)
	if err != nil {
		return
	}
	t2, err := template.ParseFiles(uc.cfg.Gen.Template.Dao)
	if err != nil {
		return
	}
	t3, err := template.ParseFiles(uc.cfg.Gen.Template.Js)
	if err != nil {
		return
	}
	t4, err := template.ParseFiles(uc.cfg.Gen.Template.Vue)
	if err != nil {
		return
	}

	table := model.SysTables{}
	table.TableId = (req.TableId)
	tab, err := uc.d.GetSysTables(&table)
	if err != nil {
		return
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	// var b5 bytes.Buffer
	// err = t5.Execute(&b5, tab)

	reply = &pb.PreviewReply{
		Model: b1.String(),
		Dao:   b2.String(),
		Js:    b3.String(),
		Vue:   b4.String(),
	}
	return
}

func (uc *AdminUsecase) GenMenuAndApi(ctx context.Context, req *pb.GenMenuAndApiRequest) (reply *pb.GenMenuAndApiReply, err error) {

	table := model.SysTables{}
	timeNow := tools.GetCurrentTime()
	table.TableId = (req.TableId)
	tab, _ := uc.d.GetSysTables(&table)
	Mmenu := model.Menu{}
	Mmenu.MenuName = tab.TBName + "管理"
	Mmenu.Title = tab.TableComment
	Mmenu.Icon = "pass"
	Mmenu.Path = "/" + tab.TBName
	Mmenu.MenuType = "M"
	Mmenu.Action = "无"
	Mmenu.ParentId = 0
	Mmenu.NoCache = false
	Mmenu.Component = "Layout"
	Mmenu.Sort = 0
	Mmenu.Visible = true
	Mmenu.IsFrame = false
	Mmenu.CreateBy = "1"
	Mmenu.UpdateBy = "1"
	Mmenu.CreatedAt = timeNow
	Mmenu.UpdatedAt = timeNow
	Mmenu.MenuId, err = uc.d.CreateMenu(&Mmenu)

	Cmenu := model.Menu{}
	Cmenu.MenuName = tab.TBName + "管理"
	Cmenu.Title = tab.TableComment
	Cmenu.Icon = "pass"
	Cmenu.Path = tab.TBName
	Cmenu.MenuType = "C"
	Cmenu.Action = "无"
	Cmenu.Permission = tab.PackageName + ":" + tab.ModuleName + ":list"
	Cmenu.ParentId = Mmenu.MenuId
	Cmenu.NoCache = false
	Cmenu.Component = "/" + tab.ModuleName + "/index"
	Cmenu.Sort = 0
	Cmenu.Visible = true
	Cmenu.IsFrame = false
	Cmenu.CreateBy = "1"
	Cmenu.UpdateBy = "1"
	Cmenu.CreatedAt = timeNow
	Cmenu.UpdatedAt = timeNow
	Cmenu.MenuId, err = uc.d.CreateMenu(&Cmenu)

	MList := model.Menu{}
	MList.MenuName = tab.TBName
	MList.Title = "分页获取" + tab.TableComment
	MList.Icon = "pass"
	MList.Path = tab.TBName
	MList.MenuType = "F"
	MList.Action = "无"
	MList.Permission = tab.PackageName + ":" + tab.ModuleName + ":query"
	MList.ParentId = Cmenu.MenuId
	MList.NoCache = false
	MList.Sort = 0
	MList.Visible = true
	MList.IsFrame = false
	MList.CreateBy = "1"
	MList.UpdateBy = "1"
	MList.CreatedAt = timeNow
	MList.UpdatedAt = timeNow
	MList.MenuId, err = uc.d.CreateMenu(&MList)

	MCreate := model.Menu{}
	MCreate.MenuName = tab.TBName
	MCreate.Title = "创建" + tab.TableComment
	MCreate.Icon = "pass"
	MCreate.Path = tab.TBName
	MCreate.MenuType = "F"
	MCreate.Action = "无"
	MCreate.Permission = tab.PackageName + ":" + tab.ModuleName + ":add"
	MCreate.ParentId = Cmenu.MenuId
	MCreate.NoCache = false
	MCreate.Sort = 0
	MCreate.Visible = true
	MCreate.IsFrame = false
	MCreate.CreateBy = "1"
	MCreate.UpdateBy = "1"
	MCreate.CreatedAt = timeNow
	MCreate.UpdatedAt = timeNow
	MCreate.MenuId, err = uc.d.CreateMenu(&MCreate)

	MUpdate := model.Menu{}
	MUpdate.MenuName = tab.TBName
	MUpdate.Title = "修改" + tab.TableComment
	MUpdate.Icon = "pass"
	MUpdate.Path = tab.TBName
	MUpdate.MenuType = "F"
	MUpdate.Action = "无"
	MUpdate.Permission = tab.PackageName + ":" + tab.ModuleName + ":edit"
	MUpdate.ParentId = Cmenu.MenuId
	MUpdate.NoCache = false
	MUpdate.Sort = 0
	MUpdate.Visible = true
	MUpdate.IsFrame = false
	MUpdate.CreateBy = "1"
	MUpdate.UpdateBy = "1"
	MUpdate.CreatedAt = timeNow
	MUpdate.UpdatedAt = timeNow
	MUpdate.MenuId, err = uc.d.CreateMenu(&MUpdate)

	MDelete := model.Menu{}
	MDelete.MenuName = tab.TBName
	MDelete.Title = "删除" + tab.TableComment
	MDelete.Icon = "pass"
	MDelete.Path = tab.TBName
	MDelete.MenuType = "F"
	MDelete.Action = "无"
	MDelete.Permission = tab.PackageName + ":" + tab.ModuleName + ":remove"
	MDelete.ParentId = Cmenu.MenuId
	MDelete.NoCache = false
	MDelete.Sort = 0
	MDelete.Visible = true
	MDelete.IsFrame = false
	MDelete.CreateBy = "1"
	MDelete.UpdateBy = "1"
	MDelete.CreatedAt = timeNow
	MDelete.UpdatedAt = timeNow
	MDelete.MenuId, err = uc.d.CreateMenu(&MDelete)

	var InterfaceId = int32(63)
	Amenu := model.Menu{}
	Amenu.MenuName = tab.TBName
	Amenu.Title = tab.TableComment
	Amenu.Icon = "bug"
	Amenu.Path = tab.TBName
	Amenu.MenuType = "M"
	Amenu.Action = "无"
	Amenu.ParentId = InterfaceId
	Amenu.NoCache = false
	Amenu.Sort = 0
	Amenu.Visible = true
	Amenu.IsFrame = false
	Amenu.CreateBy = "1"
	Amenu.UpdateBy = "1"
	Amenu.CreatedAt = timeNow
	Amenu.UpdatedAt = timeNow
	Amenu.MenuId, err = uc.d.CreateMenu(&Amenu)

	AList := model.Menu{}
	AList.MenuName = tab.TBName
	AList.Title = "分页获取" + tab.TableComment
	AList.Icon = "bug"
	AList.Path = "/api/v1/" + tab.ModuleName + "List"
	AList.MenuType = "A"
	AList.Action = "GET"
	AList.ParentId = Amenu.MenuId
	AList.NoCache = false
	AList.Sort = 0
	AList.Visible = true
	AList.IsFrame = false
	AList.CreateBy = "1"
	AList.UpdateBy = "1"
	AList.CreatedAt = timeNow
	AList.UpdatedAt = timeNow
	AList.MenuId, err = uc.d.CreateMenu(&AList)

	AGet := model.Menu{}
	AGet.MenuName = tab.TBName
	AGet.Title = "根据id获取" + tab.TableComment
	AGet.Icon = "bug"
	AGet.Path = "/api/v1/" + tab.ModuleName + "/:id"
	AGet.MenuType = "A"
	AGet.Action = "GET"
	AGet.ParentId = Amenu.MenuId
	AGet.NoCache = false
	AGet.Sort = 0
	AGet.Visible = true
	AGet.IsFrame = false
	AGet.CreateBy = "1"
	AGet.UpdateBy = "1"
	AGet.CreatedAt = timeNow
	AGet.UpdatedAt = timeNow
	AGet.MenuId, err = uc.d.CreateMenu(&AGet)

	ACreate := model.Menu{}
	ACreate.MenuName = tab.TBName
	ACreate.Title = "创建" + tab.TableComment
	ACreate.Icon = "bug"
	ACreate.Path = "/api/v1/" + tab.ModuleName
	ACreate.MenuType = "A"
	ACreate.Action = "POST"
	ACreate.ParentId = Amenu.MenuId
	ACreate.NoCache = false
	ACreate.Sort = 0
	ACreate.Visible = true
	ACreate.IsFrame = false
	ACreate.CreateBy = "1"
	ACreate.UpdateBy = "1"
	ACreate.CreatedAt = timeNow
	ACreate.UpdatedAt = timeNow
	ACreate.MenuId, err = uc.d.CreateMenu(&ACreate)

	AUpdate := model.Menu{}
	AUpdate.MenuName = tab.TBName
	AUpdate.Title = "修改" + tab.TableComment
	AUpdate.Icon = "bug"
	AUpdate.Path = "/api/v1/" + tab.ModuleName
	AUpdate.MenuType = "A"
	AUpdate.Action = "PUT"
	AUpdate.ParentId = Amenu.MenuId
	AUpdate.NoCache = false
	AUpdate.Sort = 0
	AUpdate.Visible = true
	AUpdate.IsFrame = false
	AUpdate.CreateBy = "1"
	AUpdate.UpdateBy = "1"
	AUpdate.CreatedAt = timeNow
	AUpdate.UpdatedAt = timeNow
	AUpdate.MenuId, err = uc.d.CreateMenu(&AUpdate)

	ADelete := model.Menu{}
	ADelete.MenuName = tab.TBName
	ADelete.Title = "删除" + tab.TableComment
	ADelete.Icon = "bug"
	ADelete.Path = "/api/v1/" + tab.ModuleName + "/:id"
	ADelete.MenuType = "A"
	ADelete.Action = "DELETE"
	ADelete.ParentId = Amenu.MenuId
	ADelete.NoCache = false
	ADelete.Sort = 0
	ADelete.Visible = true
	ADelete.IsFrame = false
	ADelete.CreateBy = "1"
	ADelete.UpdateBy = "1"
	ADelete.CreatedAt = timeNow
	ADelete.UpdatedAt = timeNow
	ADelete.MenuId, err = uc.d.CreateMenu(&ADelete)
	if err != nil {
		return
	}
	return
}
