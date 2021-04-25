package ecode

// import (
// 	"strings"

// 	pkgerr "github.com/pkg/errors"
// )

// // FormatErr 格式字符串
// func FormatErr(err error) error {
// 	e := pkgerr.Cause(err)
// 	switch e {
// 	default:
// 		es := e.Error()
// 		switch {
// 		case strings.HasPrefix(es, "read"):
// 			return ReadTimeout
// 		case strings.HasPrefix(es, "dial"):
// 			return DialTimeout
// 		case strings.HasPrefix(es, "write"):
// 			return WriteTimeout
// 		case strings.Contains(es, "EOF"):
// 			return EOF
// 		case strings.Contains(es, "reset"):
// 			return Reset
// 		case strings.Contains(es, "broken"):
// 			return BrokenPipe
// 		case strings.HasPrefix(es, "rpc error:"):
// 			return RpcConnectionClosed
// 		case strings.Contains(es, "hashedPassword is not the hash of the given password"):
// 			return ErrPassword
// 		default:
// 			return UnexpectedErr
// 		}
// 	}
// }

// // FormatMgoErr 格式字符串
// func FormatMgoErr(err error) error {
// 	e := pkgerr.Cause(err)
// 	switch e {
// 	default:
// 		es := e.Error()
// 		switch {
// 		case strings.HasPrefix(es, "read"):
// 			return ReadTimeout
// 		case strings.HasPrefix(es, "dial"):
// 			return DialTimeout
// 		case strings.HasPrefix(es, "write"):
// 			return WriteTimeout
// 		case strings.Contains(es, "EOF"):
// 			return EOF
// 		case strings.Contains(es, "reset"):
// 			return Reset
// 		case strings.Contains(es, "broken"):
// 			return BrokenPipe
// 		case strings.HasPrefix(es, "rpc error:"):
// 			return RpcConnectionClosed
// 		case strings.Contains(es, "hashedPassword is not the hash of the given password"):
// 			return ErrPassword
// 		default:
// 			return UnexpectedErr
// 		}
// 	}
// }
