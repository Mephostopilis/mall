package ecode

import (
	"encoding/json"

	"github.com/go-kratos/kratos/v2/errors"
)

func JsonSyntaxError(err *json.SyntaxError) error {
	return InternalServer("JsonSyntaxError", err.Error())
}

func IsJsonSyntaxError(err error) bool {
	se := errors.FromError(err)
	return se.Reason == "JsonSyntaxError" && se.Code == 500
}
