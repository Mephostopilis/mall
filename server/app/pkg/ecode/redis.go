package ecode

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
)

func RedisNil() error {
	return InternalServer("RedisNil", redis.Nil.Error())
}

func IsRedisNil(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "RedisNil" && se.Code == 500
}
