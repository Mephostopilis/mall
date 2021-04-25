package dao

import (
	"context"
)

var d *dao
var ctx = context.Background()

// func TestMain(m *testing.M) {
// 	flag.Set("conf", "../../test")
// 	flag.Set("f", "../../test/docker-compose.yaml")
// 	flag.Parse()
// 	disableLich := os.Getenv("DISABLE_LICH") != ""
// 	disableLich = true
// 	if !disableLich {
// 		if err := lich.Setup(); err != nil {
// 			panic(err)
// 		}
// 	}
// 	var err error
// 	if err = paladin.Init(); err != nil {
// 		panic(err)
// 	}
// 	var cf func()
// 	if d, cf, err = newTestDao(); err != nil {
// 		panic(err)
// 	}
// 	ret := m.Run()
// 	cf()
// 	if !disableLich {
// 		_ = lich.Teardown()
// 	}
// 	os.Exit(ret)
// }

// func TestGetVipAll(t *testing.T) {
// 	list, err := d.GetVipAll(ctx)
// 	if err != nil {
// 		t.Errorf("err = %v", err)
// 	}
// 	for _, v := range list {
// 		t.Logf("vip = %v", v)
// 	}
// }

// func TestGetUserAll(t *testing.T) {
// 	list, err := d.GetUserAll(ctx)
// 	if err != nil {
// 		t.Errorf("err = %v", err)
// 	}
// 	for _, v := range list {
// 		t.Logf("u = %v", v)
// 	}
// }
