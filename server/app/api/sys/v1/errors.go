package v1

import (
	"fmt"

	"edu/pkg/ecode"

	"github.com/go-kratos/kratos/v2/errors"
)

func ErrExistDictType(format string, a ...interface{}) error {
	return ecode.InternalServer("12500", fmt.Sprintf(format, a...))
}

func IsErrExistDictType(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12500" {
			return true
		}
	}
	return false
}

func ErrDontFixConfigName(format string, a ...interface{}) error {
	return ecode.InternalServer("12501", fmt.Sprintf(format, a...))
}

func IsErrDontFixConfigName(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12501" {
			return true
		}
	}
	return false
}

func ErrDontFixConfigKey(format string, a ...interface{}) error {
	return ecode.InternalServer("12502", fmt.Sprintf(format, a...))
}

func IsErrDontFixConfigKey(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12202" {
			return true
		}
	}
	return false
}

func ErrExistConfig(format string, a ...interface{}) error {
	return ecode.InternalServer("12503", fmt.Sprintf(format, a...))
}

func IsErrExistConfig(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12203" {
			return true
		}
	}
	return false
}

func ErrDontFixDictDataLable(format string, a ...interface{}) error {
	return ecode.InternalServer("12504", fmt.Sprintf(format, a...))
}

func IsErrDontFixDictDataLable(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12504" {
			return true
		}
	}
	return false
}

func ErrDontFixDictDataValue(format string, a ...interface{}) error {
	return ecode.InternalServer("12505", fmt.Sprintf(format, a...))
}

func IsErrDontFixDictDataValue(err error) bool {
	se := errors.FromError(err)
	if se.Code == 500 {
		if se.Reason == "12205" {
			return true
		}
	}
	return false
}

func ErrOutOfRange(format string, a ...interface{}) error {
	return ecode.InternalServer("ErrOutOfRange", fmt.Sprintf(format, a...))
}

func IsErrOutOfRange(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "ErrOutOfRange" && se.Code == 500
}

func UnexpectedErr(format string, a ...interface{}) error {
	return ecode.InternalServer("UnexpectedErr", fmt.Sprintf(format, a...))
}

func IsUnexpectedErr(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "UnexpectedErr" && se.Code == 500
}
