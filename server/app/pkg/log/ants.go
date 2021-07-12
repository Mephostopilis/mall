package log

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"
)

type antsLogger struct {
	logger log.Logger
}

// NewZapLogger new a zap logger with options.
func NewAntsLogger(logger log.Logger) ants.Logger {
	return &antsLogger{
		logger: logger,
	}
}

// Print print the kv pairs log.
func (s *antsLogger) Printf(format string, args ...interface{}) {
	s.logger.Log(log.LevelInfo, fmt.Sprintf(format, args...))
}
