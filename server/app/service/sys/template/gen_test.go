package template

import (
	"os"
	"testing"
	"text/template"

	"edu/service/sys/internal/conf"
	"edu/service/sys/internal/dao"
	"edu/service/sys/internal/model"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

func TestGoModelTemplate(t *testing.T) {
	c := config.New(config.WithSource(
		file.NewSource("../configs/admin.yaml"),
	))
	if err := c.Load(); err != nil {
		panic(err)
	}
	logger := log.NewStdLogger(os.Stdout)
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	d, err := dao.New(bc.Data, logger)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	defer d.Close()
	t1, err := template.ParseFiles("model.go.template")
	if err != nil {
		t.Error(err)
	}
	table := model.SysTables{}
	tab, err := d.GetSysTables(&table)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("models/" + table.PackageName + ".go")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	_ = t1.Execute(file, tab)
	t.Log("")
}

func TestGoApiTemplate(t *testing.T) {
	c := config.New(config.WithSource(
		file.NewSource("../configs/admin.yaml"),
	))
	if err := c.Load(); err != nil {
		panic(err)
	}

	logger := log.NewStdLogger(os.Stdout)
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	d, err := dao.New(bc.Data, logger)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	defer d.Close()
	t1, err := template.ParseFiles("api.go.template")
	if err != nil {
		t.Error(err)
	}
	table := model.SysTables{}
	tab, _ := d.GetSysTables(&table)
	file, err := os.Create("apis/" + table.PackageName + ".go")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	_ = t1.Execute(file, tab)
	t.Log("")
}
