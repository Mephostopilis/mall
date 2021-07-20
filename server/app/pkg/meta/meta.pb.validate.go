// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/meta/meta.proto

package meta

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on DataPermission with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DataPermission) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DataScope

	// no validation rules for UserId

	// no validation rules for RoleId

	// no validation rules for RoleKey

	return nil
}

// DataPermissionValidationError is the validation error returned by
// DataPermission.Validate if the designated constraints aren't met.
type DataPermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataPermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataPermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataPermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataPermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataPermissionValidationError) ErrorName() string { return "DataPermissionValidationError" }

// Error satisfies the builtin error interface
func (e DataPermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataPermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataPermissionValidationError{}