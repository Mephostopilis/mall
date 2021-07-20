package ecode

import (
	pkgerr "github.com/pkg/errors"
)

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(reason, message string) error {
	return pkgerr.WithStack(Newf(400, reason, message))
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(reason, message string) error {
	return pkgerr.WithStack(Newf(401, reason, message))
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string) error {
	return pkgerr.WithStack(Newf(403, reason, message))
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(reason, message string) error {
	return pkgerr.WithStack(Newf(404, reason, message))
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(reason, message string) error {
	return pkgerr.WithStack(Newf(409, reason, message))
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string) error {
	return pkgerr.WithStack(Newf(500, reason, message))
}
