package gcasbin

import (
	"net/http"

	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func mySubjectFn(w http.ResponseWriter, r *http.Request) string {
	return ""
}

// Config is the identify config model.
type Config struct {
	Driver string
	Source string
}

// New 默認配置
func New(cfg *Config, logger log.Logger) (c *CasbinMiddleware, err error) {
	db, err := gorm.Open(mysql.Open(cfg.Source), &gorm.Config{})
	if err != nil {
		return
	}
	// db.DB().SetMaxIdleConns(cfg.Idle)
	// db.DB().SetMaxOpenConns(cfg.Active)
	apter, err := gormAdapter.NewAdapterByDB(db)
	if err != nil {
		return
	}
	return newCasbinMiddleware(logger, text, apter, mySubjectFn)
}
