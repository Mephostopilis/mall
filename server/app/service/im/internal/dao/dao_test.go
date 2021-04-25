package dao

import (
	"flag"
	"os"
	"testing"
)

var (
	d *dao
)

func TestMain(m *testing.M) {
	if err := flag.Set("conf", "../../../cmd/logic/logic-example.toml"); err != nil {
		panic(err)
	}
	flag.Parse()
	// if err := conf.Init(); err != nil {
	// 	panic(err)
	// }
	// d, _, err := newDao(conf.Conf)
	// if err != nil {
	// 	os.Exit(-1)
	// }
	// if err := d.Ping(context.TODO()); err != nil {
	// 	os.Exit(-1)
	// }
	// d.Close()
	// if err := d.Close(); err != nil {
	// 	os.Exit(-1)
	// }
	// if err := d.Ping(context.TODO()); err == nil {
	// 	os.Exit(-1)
	// }
	// d, _, err = newDao(conf.Conf)
	// if err != nil {
	// 	os.Exit(-1)
	// }
	os.Exit(m.Run())
}
