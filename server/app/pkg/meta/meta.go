package meta

import (
	"context"
	"encoding/json"

	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/metadata"
)

func GetDataPermissions(ctx context.Context) (permission *DataPermission, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ecode.Unknown("不存在md")
		return
	}
	v := md.Get("x-md-global-dp")
	if len(v) <= 0 {
		err = ecode.Unknown("不存在dp")
		return
	}
	var dp DataPermission
	if err = json.Unmarshal([]byte(v), &dp); err != nil {
		return
	}
	permission = &dp
	return
}

func GetToken(ctx context.Context) (token string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ecode.Unknown("不存在md")
		return
	}
	v := md.Get("x-md-global-token")
	if len(v) <= 0 {
		err = ecode.Unknown("不存在token")
		return
	}
	token = v
	return
}

func GetUA(ctx context.Context) (ua string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ecode.Unknown("不存在md")
		return
	}
	v := md.Get("ua")
	if len(v) <= 0 {
		err = ecode.Unknown("meta")
		return
	}
	ua = v
	return
}

func GetRemoteAddr(ctx context.Context) (remoteAddr string, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ecode.Unknown("不存在md")
		return
	}
	v := md.Get("remote_addr")
	if len(v) <= 0 {
		err = ecode.Unknown("meta")
		return
	}
	remoteAddr = v
	return
}
