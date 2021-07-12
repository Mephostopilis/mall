package ecode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

func SaramaErrClosedClient(err error) error {
	se := InternalServer("SaramaErrClosedClient", err.Error())
	return se
}

func IsSaramaErrClosedClient(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "SaramaErrClosedClient" && se.Code == 500
}
