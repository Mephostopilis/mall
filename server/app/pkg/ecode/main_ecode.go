package ecode

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

func Unknown(domain, reason, format string, a ...interface{}) error {
	return InternalServer("Unknown", fmt.Sprintf(format, a...))
}

// main ecode interval is [0,990000]
func SignCheckErr(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10001", "签名错误"))
}

func RequestErr(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10002", "应用不存在错误"))
}

func ServerErr(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10003", "token接口失效"))
}

func KeyNotFound(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10004", "接口参数非法"))
}

// 系统错误
func ReadTimeout(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("10100", "读取超时"))
}

// admin jwt
// ErrMissingSecretKey indicates Secret key is required
func ErrMissingSecretKey(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11101", "secret key is required"))
}

func ErrForbidden(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11102", "禁止")) // ErrForbidden when HTTP status 403 is give
}

func ErrUpdateSysUser(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11103", "")) // ErrForbidden when HTTP status 403 is give
}

func AdminTableEmpty(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11105", "")) // ErrForbidden when HTTP status 403 is give
}

func AdminNotSupportedJob(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11106", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrGetSysUser(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11107", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrUsername(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11108", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrNoPrivKeyFile(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11109", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidPrivKey(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11110", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidSigningAlgorithm(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11111", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrExpiredToken(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11112", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrMissingExpField(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11113", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrWrongFormatOfExp(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11114", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrMissingLoginValues(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11115", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidVerificationode(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11116", "无效的验证码")) // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedResponseType(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11117", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrUnauthorizedClient(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11118", fmt.Sprintf(format, a...))) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidScope(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11119", fmt.Sprintf(format, a...))) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUserId(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11120", fmt.Sprintf(format, a...))) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidRequest(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11121", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedGrantType(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11122", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUsername(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11123", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUid(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11124", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAuthorizeCode(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11125", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidClient(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11126", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedErrType(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11127", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidGrant(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11128", "")) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAccessToken(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11129", "")) // ErrForbidden when HTTP status 403 is give
}

func UserNotExist(domain, reason, format string, a ...interface{}) error {
	return (InternalServer("11130", "")) // ErrForbidden when HTTP status 403 is give
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
