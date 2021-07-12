package ecode

import (
	pkgerr "github.com/pkg/errors"
)

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string) error {
	return pkgerr.WithStack(Newf(403, reason, message))
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string) error {
	return pkgerr.WithStack(Newf(500, reason, message))
}
