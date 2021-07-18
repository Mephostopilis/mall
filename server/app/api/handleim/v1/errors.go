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

var (
	// ErrComet commet error.
	ErrComet = ecode.InternalServer("ErrComet", "comet rpc is not available")
	// ErrCometFull comet chan full.
	ErrCometFull = ecode.InternalServer( "ErrCometFull","comet proto chan full")
	// ErrRoomFull room chan full.
	ErrRoomFull = ecode.InternalServer("ErrRoomFull","room proto chan full")
)
