package ecode

import (
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// FromError 格式字符串
func WrapError(err error) error {
	if err == nil {
		return nil
	}
	if se := new(errors.Error); errors.As(err, &se) {
		return se
	}
	if errors.Is(err, redis.Nil) {
		return RedisNil()
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return GormErrRecordNotFound(err)
	}
	var jsonSyntaxError *json.SyntaxError
	if errors.As(err, &jsonSyntaxError) {
		return JsonSyntaxError(err.(*json.SyntaxError))
	}
	return Unknown("", "", err.Error())
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *errors.Error {
	// const size = 64 << 10
	// buf := make([]byte, size)
	// buf = buf[:runtime.Stack(buf, false)]
	// pl := fmt.Sprintf("%s\n", buf)
	se := errors.New(code, reason, fmt.Sprintf(format, a...))
	// md := map[string]string{
	// 	"stack": pl,
	// }
	// se = se.WithMetadata(md)
	return se
}
