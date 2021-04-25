package biz

import (
	"flag"
	"os"
	"testing"
)

var (
	lg *Logic
)

func TestMain(m *testing.M) {
	if err := flag.Set("conf", "../../cmd/logic/logic-example.toml"); err != nil {
		panic(err)
	}
	flag.Parse()
	// if err := conf.Init(); err != nil {
	// 	panic(err)
	// }
	// lg = New()
	// if err := lg.Ping(context.TODO()); err != nil {
	// 	panic(err)
	// }
	os.Exit(m.Run())
}
