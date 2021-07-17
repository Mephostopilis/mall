package v1

import (
	"fmt"

	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/errors"
)

// admin jwt
// ErrMissingSecretKey indicates Secret key is required
func ErrMissingSecretKey(format string, a ...interface{}) error {
	return ecode.InternalServer("11101", "secret key is required")
}

func ErrUpdateSysUser(format string, a ...interface{}) error {
	return ecode.InternalServer("11103", "") // ErrForbidden when HTTP status 403 is give
}

func AdminTableEmpty(format string, a ...interface{}) error {
	return ecode.InternalServer("11105", "") // ErrForbidden when HTTP status 403 is give
}

func AdminNotSupportedJob(format string, a ...interface{}) error {
	return ecode.InternalServer("11106", "") // ErrForbidden when HTTP status 403 is give
}

func ErrGetSysUser(format string, a ...interface{}) error {
	return ecode.InternalServer("11107", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUsername(format string, a ...interface{}) error {
	return ecode.InternalServer("11108", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidPrivKey(format string, a ...interface{}) error {
	return ecode.InternalServer("11110", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidSigningAlgorithm(format string, a ...interface{}) error {
	return ecode.InternalServer("11111", "") // ErrForbidden when HTTP status 403 is give
}

func ErrExpiredToken(format string, a ...interface{}) error {
	return ecode.InternalServer("11112", "") // ErrForbidden when HTTP status 403 is give
}

func ErrMissingExpField(format string, a ...interface{}) error {
	return ecode.InternalServer("11113", "") // ErrForbidden when HTTP status 403 is give
}

func ErrWrongFormatOfExp(format string, a ...interface{}) error {
	return ecode.InternalServer("11114", "") // ErrForbidden when HTTP status 403 is give
}

func ErrMissingLoginValues(format string, a ...interface{}) error {
	return ecode.InternalServer("11115", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidVerificationode(format string, a ...interface{}) error {
	return ecode.InternalServer("11116", "无效的验证码") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedResponseType(format string, a ...interface{}) error {
	return ecode.InternalServer("11117", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrUnauthorizedClient(format string, a ...interface{}) error {
	return ecode.InternalServer("11118", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidScope(format string, a ...interface{}) error {
	return ecode.InternalServer("11119", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUserId(format string, a ...interface{}) error {
	return ecode.InternalServer("11120", fmt.Sprintf(format, a...)) // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidRequest(format string, a ...interface{}) error {
	return ecode.InternalServer("11121", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedGrantType(format string, a ...interface{}) error {
	return ecode.InternalServer("11122", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUsername(format string, a ...interface{}) error {
	return ecode.InternalServer("11123", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidUid(format string, a ...interface{}) error {
	return ecode.InternalServer("11124", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAuthorizeCode(format string, a ...interface{}) error {
	return ecode.InternalServer("11125", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidClient(format string, a ...interface{}) error {
	return ecode.InternalServer("11126", "") // ErrForbidden when HTTP status 403 is give
}

func ErrUnsupportedErrType(format string, a ...interface{}) error {
	return ecode.InternalServer("11127", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidGrant(format string, a ...interface{}) error {
	return ecode.InternalServer("11128", "") // ErrForbidden when HTTP status 403 is give
}

func ErrInvalidAccessToken(format string, a ...interface{}) error {
	return ecode.InternalServer("11129", "") // ErrForbidden when HTTP status 403 is give
}

func UserNotExist(format string, a ...interface{}) error {
	return ecode.InternalServer("11130", "") // ErrForbidden when HTTP status 403 is give
}

func ErrOutOfRange(format string, a ...interface{}) error {
	return ecode.InternalServer("ErrOutOfRange", fmt.Sprintf(format, a...))
}

func IsErrOutOfRange(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrOutOfRange" && se.Code == 500
}
