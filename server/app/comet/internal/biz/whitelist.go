package biz

import (
	"os"

	"edu/comet/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

var whitelist *Whitelist

// Whitelist .
type Whitelist struct {
	log  *log.Helper
	list map[int64]struct{} // whitelist for debug
}

// InitWhitelist a whitelist struct.
func InitWhitelist(c *conf.Server_Whitelist, logger log.Logger) (err error) {
	var (
		mid int64
	)
	if _, err := os.OpenFile(c.WhiteLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err == nil {
		whitelist = new(Whitelist)
		whitelist.log = log.NewHelper("biz/whitelist", logger)
		whitelist.list = make(map[int64]struct{})
		for _, mid = range c.Whitelist {
			whitelist.list[mid] = struct{}{}
		}
	}
	return
}

// Contains whitelist contains a mid or not.
func (w *Whitelist) Contains(mid int64) (ok bool) {
	if mid > 0 {
		_, ok = w.list[mid]
	}
	return
}

// Printf calls l.Output to print to the logger.
func (w *Whitelist) Printf(format string, v ...interface{}) {
	w.log.Infof(format, v...)
}
