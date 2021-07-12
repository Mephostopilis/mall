package ecode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func GormErrRecordNotFound(err error) error {
	return InternalServer("GormErrRecordNotFound", err.Error())
}

func IsGormErrRecordNotFound(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "GormErrRecordNotFound" && se.Code == 500
}
