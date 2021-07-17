package biz

import (
	"os"
	"testing"
	"time"

	ssopb "edu/api/sso"
	. "edu/pkg/jwtauth"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	logger := log.NewStdLogger(os.Stderr)
	du, _ := time.ParseDuration("3600s")
	mw, err := New(&Jwt{Secret: "xx", Timeout: du}, logger, DataPermissionToClaimsFunc, ClaimsToDataPermissionFunc)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	dp := &ssopb.DataPermission{
		UserId:  11,
		RoleId:  1,
		RoleKey: "admin",
	}
	tokenstr, _, err := mw.TokenGenerator(dp)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	out, err := mw.ValidationToken(tokenstr)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	o := out.(*ssopb.DataPermission)
	assert.Equal(t, dp.UserId, o.UserId)

	tokenstr1, _, err := mw.RefreshToken(tokenstr)
	if err != nil {
		t.Errorf("err = %v", err)
		return
	}
	assert.NotEqual(t, tokenstr, tokenstr1)
}

func TestRefresh(t *testing.T) {

}
