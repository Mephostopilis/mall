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
