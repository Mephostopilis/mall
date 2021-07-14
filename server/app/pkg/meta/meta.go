package meta

import (
	"context"

	ssopb "edu/api/sso"
	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"google.golang.org/protobuf/proto"
)

func GetDataPermissions(ctx context.Context) (permission ssopb.DataPermission, err error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		err = ecode.Unknown("meta", "error")
		return
	}
	v := md.Get("permision")
	if len(v) < 0 {
		err = ecode.Unknown("meta", "error")
		return
	}
	if err = proto.Unmarshal([]byte(v[0]), &permission); err != nil {
		return
	}
	return
}

func GetToken(ctx context.Context) (token string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("token")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	token = v[0]
	return
}

func GetUA(ctx context.Context) (ua string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("ua")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	ua = v[0]
	return
}

func GetRemoteAddr(ctx context.Context) (remoteAddr string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = errors.Unknown("meta", "error")
		return
	}
	v := md.Get("remote_addr")
	if len(v) < 0 {
		err = errors.Unknown("meta", "error")
		return
	}
	remoteAddr = v[0]
	return
}

//IdsStrToIdsIntGroup 获取URL中批量id并解析
// func IdsStrToIdsIntGroup(key string, ctx context.Context) ([]int, error) {
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		err = errors.Unknown("meta", "error")
// 		return
// 	}
// 	v := md.Get("permision")
// 	if len(v) < 0 {
// 		err = errors.Unknown("meta", "error")
// 		return
// 	}
// 	var permission ssopb.DataPermission
// 	if err = proto.Unmarshal([]byte(v[0]), &permission); err != nil {
// 		return
// 	}
// 	p := c.Value(key).(string)
// 	return idsStrToIdsIntGroup(p)
// }

// func idsStrToIdsIntGroup(keys string) []int {
// 	IDS := make([]int, 0)
// 	ids := strings.Split(keys, ",")
// 	for i := 0; i < len(ids); i++ {
// 		ID, _ := tools.StringToInt(ids[i])
// 		IDS = append(IDS, ID)
// 	}
// 	return IDS
// }
