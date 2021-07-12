package log

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	ctx     context.Context
	cf      context.CancelFunc
	wg      sync.WaitGroup
	hook    *lumberjack.Logger
	logger  *zap.Logger
	slogger *zap.SugaredLogger
}

func getFilePath(dir, prefix string) string {
	// now := time.Now()
	return filepath.Join(dir, prefix) + ".log"
}

// NewZapLogger new a zap logger with options.
func NewZapLogger(path, prefix string, stdout bool, level zapcore.Level) *ZapLogger {

	hook := lumberjack.Logger{
		Filename:   getFilePath(path, prefix), // 日志文件路径
		MaxSize:    128,                       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                        // 日志文件最多保存多少个备份
		MaxAge:     7,                         // 文件最多保存多少天
		Compress:   false,                     // 是否压缩
	}

	// encoderConfig := zapcore.EncoderConfig{
	// 	TimeKey:        "time",
	// 	LevelKey:       "level",
	// 	NameKey:        "logger",
	// 	CallerKey:      "linenum",
	// 	MessageKey:     "msg",
	// 	StacktraceKey:  "stacktrace",
	// 	LineEnding:     zapcore.DefaultLineEnding,
	// 	EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
	// 	EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
	// 	EncodeDuration: zapcore.SecondsDurationEncoder, //
	// 	EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	// 	EncodeName:     zapcore.FullNameEncoder,
	// }

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	var writer zapcore.WriteSyncer
	if stdout {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)) // 打印到控制台和文件
	} else {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), // 编码器配置
		writer,
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCallerSkip(3)

	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	logger := zap.New(core, caller, development)

	// logger := zap.NewDevelopment(caller)
	slogger := logger.Sugar()

	ctx, cf := context.WithCancel(context.Background())
	zl := &ZapLogger{
		ctx:     ctx,
		cf:      cf,
		hook:    &hook,
		logger:  logger,
		slogger: slogger,
	}
	// zl.wg.Add(1)
	// go zl.rotate()
	return zl
}

func (s *ZapLogger) rotate() {
	defer s.wg.Done()
	for {
		select {
		case <-s.ctx.Done():
			s.hook.Rotate()
			return
		default:
			s.hook.Rotate()
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *ZapLogger) Close() {
	s.cf()
	s.wg.Wait()
	s.hook.Close()
}

// Print print the kv pairs log.
func (s *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	var (
		msg           string
		keysAndValues []interface{}
	)
	for i := 0; i < len(keyvals); i += 2 {
		key := keyvals[i]
		if key == "msg" {
			msg = keyvals[i+1].(string)
		} else {
			keysAndValues = append(keysAndValues, keyvals[i], keyvals[i+1])
		}
	}
	if level == log.LevelDebug {
		s.slogger.Debugw(msg, keysAndValues...)
	} else if level == log.LevelInfo {
		s.slogger.Infow(msg, keysAndValues...)
	} else if level == log.LevelWarn {
		s.slogger.Warnw(msg, keysAndValues...)
	} else if level == log.LevelError {
		s.slogger.Errorw(msg, keysAndValues...)
	}
	return nil
}
