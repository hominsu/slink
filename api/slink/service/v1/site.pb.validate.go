// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: slink/service/v1/site.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on PingReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PingReplyMultiError, or nil
// if none found.
func (m *PingReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PingReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	if len(errors) > 0 {
		return PingReplyMultiError(errors)
	}

	return nil
}

// PingReplyMultiError is an error wrapping multiple validation errors returned
// by PingReply.ValidateAll() if the designated constraints aren't met.
type PingReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingReplyMultiError) AllErrors() []error { return m }

// PingReplyValidationError is the validation error returned by
// PingReply.Validate if the designated constraints aren't met.
type PingReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingReplyValidationError) ErrorName() string { return "PingReplyValidationError" }

// Error satisfies the builtin error interface
func (e PingReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingReplyValidationError{}

// Validate checks the field values on CountReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CountReplyMultiError, or
// nil if none found.
func (m *CountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return CountReplyMultiError(errors)
	}

	return nil
}

// CountReplyMultiError is an error wrapping multiple validation errors
// returned by CountReply.ValidateAll() if the designated constraints aren't met.
type CountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountReplyMultiError) AllErrors() []error { return m }

// CountReplyValidationError is the validation error returned by
// CountReply.Validate if the designated constraints aren't met.
type CountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountReplyValidationError) ErrorName() string { return "CountReplyValidationError" }

// Error satisfies the builtin error interface
func (e CountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountReplyValidationError{}