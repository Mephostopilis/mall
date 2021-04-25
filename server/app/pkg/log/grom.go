package log

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"
)

type gormLogger struct {
	log      *log.Helper
	LogLevel logger.LogLevel
}

func NewGormLogger(logger log.Logger) logger.Interface {
	return &gormLogger{
		log: log.NewHelper("gorm", logger),
	}
}

// LogMode log mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.log.Infof(msg, data)
	}
}

// Warn print warn messages
func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.log.Warnf(msg, data)
	}
}

// Error print error messages
func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.log.Errorf(msg, data)
	}
}

// Trace print sql message
func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		// elapsed := time.Since(begin)
		switch {
		// case err != nil && l.LogLevel >= logger.Error:
		// 	sql, rows := fc()
		// 	if rows == -1 {
		// 		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		// 	} else {
		// 		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		// 	}
		// case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
		// 	sql, rows := fc()
		// 	slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		// 	if rows == -1 {
		// 		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		// 	} else {
		// 		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		// 	}
		// case l.LogLevel == Info:
		// 	sql, rows := fc()
		// 	if rows == -1 {
		// 		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		// 	} else {
		// 		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		// 	}
		}
	}
}
