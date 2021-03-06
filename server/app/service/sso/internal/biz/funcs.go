package biz

import "edu/pkg/meta"

func DataPermissionToClaimsFunc(data interface{}) MapClaims {
	if v, ok := data.(*meta.DataPermission); ok {
		return MapClaims{
			IdentityKey: v.UserId,
			NiceKey:     v.UserId,
			RoleIdKey:   v.RoleId,
			RoleKey:     v.RoleKey,
			RoleNameKey: v.RoleKey,
		}
	}
	return MapClaims{}
}

func ClaimsToDataPermissionFunc(claims MapClaims) (interface{}, error) {
	return &meta.DataPermission{
		UserId:  uint64(claims["identity"].(float64)),
		RoleId:  uint64(claims["roleid"].(float64)),
		RoleKey: claims["rolekey"].(string),
	}, nil
}
