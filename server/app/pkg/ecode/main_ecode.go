package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func Unknown(format string, a ...interface{}) error {
	return InternalServer("Unknown", fmt.Sprintf(format, a...))
}

// main ecode interval is [0,990000]
func SignCheckErr(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10001", "签名错误"))
}

// 系统错误
func ReadTimeout(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10100", "读取超时"))
}

func ErrSignService(format string, a ...interface{}) error {
	return InternalServer("11131", fmt.Sprintf(format, a...))
}

func IsErrSignService(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "11131" && se.Code == 500
}

func ErrBadPublicKey(format string, a ...interface{}) error {
	return InternalServer("11131", fmt.Sprintf(format, a...))
}

func IsErrBadPublicKey(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "11131" && se.Code == 500
}

func ErrNoPrivKeyFile(format string, a ...interface{}) error {
	return InternalServer("11109", "") // ErrForbidden when HTTP status 403 is give
}
