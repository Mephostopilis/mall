package log

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type cronLogger struct {
	logger log.Logger
}

// NewZapLogger new a zap logger with options.
func NewCronLogger(logger log.Logger) cron.Logger {
	return &cronLogger{
		logger: logger,
	}
}

// Print print the kv pairs log.
func (s *cronLogger) Printf(format string, args ...interface{}) {
	s.logger.Log(log.LevelInfo, fmt.Sprintf(format, args...))
}

func (s *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	s.logger.Log(log.LevelInfo, "msg", msg, keysAndValues)
}

// Error logs an error condition.
func (s *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	s.logger.Log(log.LevelError, "err", err, "msg", msg, keysAndValues)
}
