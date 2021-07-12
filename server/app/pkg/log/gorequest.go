package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/parnurzeal/gorequest"
)

type gorequestLogger struct {
	prefix string
	logger log.Logger
}

// NewZapLogger new a zap logger with options.
func NewGorequsetLogger(logger log.Logger) gorequest.Logger {
	return &gorequestLogger{
		prefix: "gorequest",
		logger: logger,
	}
}

func (l *gorequestLogger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *gorequestLogger) Printf(format string, v ...interface{}) {
	l.logger.Log(log.LevelInfo, format, v)
}

func (l *gorequestLogger) Println(v ...interface{}) {
	l.logger.Log(log.LevelInfo, v...)
}
