package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger  *zap.Logger
	slogger *zap.SugaredLogger
}

// NewZapLogger new a zap logger with options.
func NewZapLogger() log.Logger {
	logger, err := zap.NewDevelopment(zap.AddCallerSkip(3))
	if err != nil {
		panic(err)
	}
	slogger := logger.Sugar()
	return &zapLogger{
		logger:  logger,
		slogger: slogger,
	}
}

// Print print the kv pairs log.
func (s *zapLogger) Print(kvpair ...interface{}) {
	if len(kvpair) == 0 {
		return
	}
	if len(kvpair)%2 != 0 {
		kvpair = append(kvpair, "")
	}
	var level log.Level
	var message string
	fields := make([]interface{}, 0)
	for i := 0; i < len(kvpair); i += 2 {
		key := kvpair[i].(string)
		if key == "level" {
			level = kvpair[i+1].(log.Level)
		} else if key == "message" {
			message = kvpair[i+1].(string)
		} else {
			field := zap.Reflect(key, kvpair[i+1])
			fields = append(fields, field)
		}
		// fmt.Fprintf(buf, "%s=%v ", kvpair[i], kvpair[i+1])
	}
	if level == log.LevelDebug {
		s.slogger.Debugw(message, fields...)
	} else if level == log.LevelInfo {
		s.slogger.Infow(message, fields...)
	} else if level == log.LevelWarn {
		s.slogger.Warnw(message, fields...)
	} else if level == log.LevelError {
		s.slogger.Errorw(message, fields...)
	}
}
