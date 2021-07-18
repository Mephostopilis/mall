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
	ErrCometFull = ecode.InternalServer("ErrCometFull", "comet proto chan full")
	// ErrRoomFull room chan full.
	ErrRoomFull = ecode.InternalServer("ErrRoomFull", "room proto chan full")

	// server
	ErrHandshake = ecode.InternalServer("ErrHandshake", "handshake failed")
	ErrOperation = ecode.InternalServer("ErrOperation", "request operation not valid")
	// ring
	ErrRingEmpty = ecode.InternalServer("ErrRingEmpty", "ring buffer empty")
	ErrRingFull  = ecode.InternalServer("ErrRingFull", "ring buffer full")
	// timer
	ErrTimerFull   = ecode.InternalServer("ErrTimerFull", "timer full")
	ErrTimerEmpty  = ecode.InternalServer("ErrTimerEmpty", "timer empty")
	ErrTimerNoItem = ecode.InternalServer("ErrTimerNoItem", "timer item not exist")
	// channel
	ErrPushMsgArg   = ecode.InternalServer("ErrPushMsgArg", "rpc pushmsg arg error")
	ErrPushMsgsArg  = ecode.InternalServer("ErrPushMsgsArg", "rpc pushmsgs arg error")
	ErrMPushMsgArg  = ecode.InternalServer("ErrMPushMsgArg", "rpc mpushmsg arg error")
	ErrMPushMsgsArg = ecode.InternalServer("ErrMPushMsgsArg", "rpc mpushmsgs arg error")
	// bucket
	ErrBroadCastArg     = ecode.InternalServer("ErrBroadCastArg", "rpc broadcast arg error")
	ErrBroadCastRoomArg = ecode.InternalServer("ErrBroadCastRoomArg", "rpc broadcast  room arg error")

	// room
	ErrRoomDroped = ecode.InternalServer("ErrRoomDroped", "room droped")
	// rpc
	ErrLogic = ecode.InternalServer("ErrLogic", "logic rpc is not available")
)
