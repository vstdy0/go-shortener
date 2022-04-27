// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/grpc/url_service.proto

package urlService

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

// Validate checks the field values on ShortenURLReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ShortenURLReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ShortenURLReqMultiError, or
// nil if none found.
func (m *ShortenURLReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = ShortenURLReqValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := ShortenURLReqValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ShortenURLReqMultiError(errors)
	}

	return nil
}

// ShortenURLReqMultiError is an error wrapping multiple validation errors
// returned by ShortenURLReq.ValidateAll() if the designated constraints
// aren't met.
type ShortenURLReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLReqMultiError) AllErrors() []error { return m }

// ShortenURLReqValidationError is the validation error returned by
// ShortenURLReq.Validate if the designated constraints aren't met.
type ShortenURLReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLReqValidationError) ErrorName() string { return "ShortenURLReqValidationError" }

// Error satisfies the builtin error interface
func (e ShortenURLReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLReqValidationError{}

// Validate checks the field values on ShortenURLResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ShortenURLResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ShortenURLRespMultiError,
// or nil if none found.
func (m *ShortenURLResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return ShortenURLRespMultiError(errors)
	}

	return nil
}

// ShortenURLRespMultiError is an error wrapping multiple validation errors
// returned by ShortenURLResp.ValidateAll() if the designated constraints
// aren't met.
type ShortenURLRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLRespMultiError) AllErrors() []error { return m }

// ShortenURLRespValidationError is the validation error returned by
// ShortenURLResp.Validate if the designated constraints aren't met.
type ShortenURLRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLRespValidationError) ErrorName() string { return "ShortenURLRespValidationError" }

// Error satisfies the builtin error interface
func (e ShortenURLRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLRespValidationError{}

// Validate checks the field values on ShortenURLsBatchReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ShortenURLsBatchReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLsBatchReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ShortenURLsBatchReqMultiError, or nil if none found.
func (m *ShortenURLsBatchReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLsBatchReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRequest() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ShortenURLsBatchReqValidationError{
						field:  fmt.Sprintf("Request[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ShortenURLsBatchReqValidationError{
						field:  fmt.Sprintf("Request[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ShortenURLsBatchReqValidationError{
					field:  fmt.Sprintf("Request[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ShortenURLsBatchReqMultiError(errors)
	}

	return nil
}

// ShortenURLsBatchReqMultiError is an error wrapping multiple validation
// errors returned by ShortenURLsBatchReq.ValidateAll() if the designated
// constraints aren't met.
type ShortenURLsBatchReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLsBatchReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLsBatchReqMultiError) AllErrors() []error { return m }

// ShortenURLsBatchReqValidationError is the validation error returned by
// ShortenURLsBatchReq.Validate if the designated constraints aren't met.
type ShortenURLsBatchReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLsBatchReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLsBatchReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLsBatchReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLsBatchReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLsBatchReqValidationError) ErrorName() string {
	return "ShortenURLsBatchReqValidationError"
}

// Error satisfies the builtin error interface
func (e ShortenURLsBatchReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLsBatchReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLsBatchReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLsBatchReqValidationError{}

// Validate checks the field values on ShortenURLsBatchResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ShortenURLsBatchResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLsBatchResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ShortenURLsBatchRespMultiError, or nil if none found.
func (m *ShortenURLsBatchResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLsBatchResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResponse() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ShortenURLsBatchRespValidationError{
						field:  fmt.Sprintf("Response[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ShortenURLsBatchRespValidationError{
						field:  fmt.Sprintf("Response[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ShortenURLsBatchRespValidationError{
					field:  fmt.Sprintf("Response[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ShortenURLsBatchRespMultiError(errors)
	}

	return nil
}

// ShortenURLsBatchRespMultiError is an error wrapping multiple validation
// errors returned by ShortenURLsBatchResp.ValidateAll() if the designated
// constraints aren't met.
type ShortenURLsBatchRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLsBatchRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLsBatchRespMultiError) AllErrors() []error { return m }

// ShortenURLsBatchRespValidationError is the validation error returned by
// ShortenURLsBatchResp.Validate if the designated constraints aren't met.
type ShortenURLsBatchRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLsBatchRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLsBatchRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLsBatchRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLsBatchRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLsBatchRespValidationError) ErrorName() string {
	return "ShortenURLsBatchRespValidationError"
}

// Error satisfies the builtin error interface
func (e ShortenURLsBatchRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLsBatchResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLsBatchRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLsBatchRespValidationError{}

// Validate checks the field values on GetOrigURLReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetOrigURLReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrigURLReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetOrigURLReqMultiError, or
// nil if none found.
func (m *GetOrigURLReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrigURLReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetOrigURLReqMultiError(errors)
	}

	return nil
}

// GetOrigURLReqMultiError is an error wrapping multiple validation errors
// returned by GetOrigURLReq.ValidateAll() if the designated constraints
// aren't met.
type GetOrigURLReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrigURLReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrigURLReqMultiError) AllErrors() []error { return m }

// GetOrigURLReqValidationError is the validation error returned by
// GetOrigURLReq.Validate if the designated constraints aren't met.
type GetOrigURLReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrigURLReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrigURLReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrigURLReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrigURLReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrigURLReqValidationError) ErrorName() string { return "GetOrigURLReqValidationError" }

// Error satisfies the builtin error interface
func (e GetOrigURLReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrigURLReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrigURLReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrigURLReqValidationError{}

// Validate checks the field values on GetUsersURLsResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUsersURLsResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUsersURLsResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUsersURLsRespMultiError, or nil if none found.
func (m *GetUsersURLsResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUsersURLsResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResponse() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUsersURLsRespValidationError{
						field:  fmt.Sprintf("Response[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUsersURLsRespValidationError{
						field:  fmt.Sprintf("Response[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersURLsRespValidationError{
					field:  fmt.Sprintf("Response[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetUsersURLsRespMultiError(errors)
	}

	return nil
}

// GetUsersURLsRespMultiError is an error wrapping multiple validation errors
// returned by GetUsersURLsResp.ValidateAll() if the designated constraints
// aren't met.
type GetUsersURLsRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUsersURLsRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUsersURLsRespMultiError) AllErrors() []error { return m }

// GetUsersURLsRespValidationError is the validation error returned by
// GetUsersURLsResp.Validate if the designated constraints aren't met.
type GetUsersURLsRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersURLsRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersURLsRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersURLsRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersURLsRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersURLsRespValidationError) ErrorName() string { return "GetUsersURLsRespValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersURLsRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersURLsResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersURLsRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersURLsRespValidationError{}

// Validate checks the field values on DelUserURLsReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DelUserURLsReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DelUserURLsReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DelUserURLsReqMultiError,
// or nil if none found.
func (m *DelUserURLsReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DelUserURLsReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DelUserURLsReqMultiError(errors)
	}

	return nil
}

// DelUserURLsReqMultiError is an error wrapping multiple validation errors
// returned by DelUserURLsReq.ValidateAll() if the designated constraints
// aren't met.
type DelUserURLsReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DelUserURLsReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DelUserURLsReqMultiError) AllErrors() []error { return m }

// DelUserURLsReqValidationError is the validation error returned by
// DelUserURLsReq.Validate if the designated constraints aren't met.
type DelUserURLsReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DelUserURLsReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DelUserURLsReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DelUserURLsReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DelUserURLsReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DelUserURLsReqValidationError) ErrorName() string { return "DelUserURLsReqValidationError" }

// Error satisfies the builtin error interface
func (e DelUserURLsReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelUserURLsReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DelUserURLsReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DelUserURLsReqValidationError{}

// Validate checks the field values on ShortenURLsBatchReq_UrlUnit with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ShortenURLsBatchReq_UrlUnit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLsBatchReq_UrlUnit with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ShortenURLsBatchReq_UrlUnitMultiError, or nil if none found.
func (m *ShortenURLsBatchReq_UrlUnit) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLsBatchReq_UrlUnit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CorrelationId

	if uri, err := url.Parse(m.GetOriginalUrl()); err != nil {
		err = ShortenURLsBatchReq_UrlUnitValidationError{
			field:  "OriginalUrl",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := ShortenURLsBatchReq_UrlUnitValidationError{
			field:  "OriginalUrl",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ShortenURLsBatchReq_UrlUnitMultiError(errors)
	}

	return nil
}

// ShortenURLsBatchReq_UrlUnitMultiError is an error wrapping multiple
// validation errors returned by ShortenURLsBatchReq_UrlUnit.ValidateAll() if
// the designated constraints aren't met.
type ShortenURLsBatchReq_UrlUnitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLsBatchReq_UrlUnitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLsBatchReq_UrlUnitMultiError) AllErrors() []error { return m }

// ShortenURLsBatchReq_UrlUnitValidationError is the validation error returned
// by ShortenURLsBatchReq_UrlUnit.Validate if the designated constraints
// aren't met.
type ShortenURLsBatchReq_UrlUnitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLsBatchReq_UrlUnitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLsBatchReq_UrlUnitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLsBatchReq_UrlUnitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLsBatchReq_UrlUnitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLsBatchReq_UrlUnitValidationError) ErrorName() string {
	return "ShortenURLsBatchReq_UrlUnitValidationError"
}

// Error satisfies the builtin error interface
func (e ShortenURLsBatchReq_UrlUnitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLsBatchReq_UrlUnit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLsBatchReq_UrlUnitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLsBatchReq_UrlUnitValidationError{}

// Validate checks the field values on ShortenURLsBatchResp_UrlUnit with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ShortenURLsBatchResp_UrlUnit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShortenURLsBatchResp_UrlUnit with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ShortenURLsBatchResp_UrlUnitMultiError, or nil if none found.
func (m *ShortenURLsBatchResp_UrlUnit) ValidateAll() error {
	return m.validate(true)
}

func (m *ShortenURLsBatchResp_UrlUnit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CorrelationId

	// no validation rules for ShortUrl

	if len(errors) > 0 {
		return ShortenURLsBatchResp_UrlUnitMultiError(errors)
	}

	return nil
}

// ShortenURLsBatchResp_UrlUnitMultiError is an error wrapping multiple
// validation errors returned by ShortenURLsBatchResp_UrlUnit.ValidateAll() if
// the designated constraints aren't met.
type ShortenURLsBatchResp_UrlUnitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShortenURLsBatchResp_UrlUnitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShortenURLsBatchResp_UrlUnitMultiError) AllErrors() []error { return m }

// ShortenURLsBatchResp_UrlUnitValidationError is the validation error returned
// by ShortenURLsBatchResp_UrlUnit.Validate if the designated constraints
// aren't met.
type ShortenURLsBatchResp_UrlUnitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShortenURLsBatchResp_UrlUnitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShortenURLsBatchResp_UrlUnitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShortenURLsBatchResp_UrlUnitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShortenURLsBatchResp_UrlUnitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShortenURLsBatchResp_UrlUnitValidationError) ErrorName() string {
	return "ShortenURLsBatchResp_UrlUnitValidationError"
}

// Error satisfies the builtin error interface
func (e ShortenURLsBatchResp_UrlUnitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShortenURLsBatchResp_UrlUnit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShortenURLsBatchResp_UrlUnitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShortenURLsBatchResp_UrlUnitValidationError{}

// Validate checks the field values on GetUsersURLsResp_UrlUnit with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUsersURLsResp_UrlUnit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUsersURLsResp_UrlUnit with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUsersURLsResp_UrlUnitMultiError, or nil if none found.
func (m *GetUsersURLsResp_UrlUnit) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUsersURLsResp_UrlUnit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShortUrl

	// no validation rules for OriginalUrl

	if len(errors) > 0 {
		return GetUsersURLsResp_UrlUnitMultiError(errors)
	}

	return nil
}

// GetUsersURLsResp_UrlUnitMultiError is an error wrapping multiple validation
// errors returned by GetUsersURLsResp_UrlUnit.ValidateAll() if the designated
// constraints aren't met.
type GetUsersURLsResp_UrlUnitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUsersURLsResp_UrlUnitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUsersURLsResp_UrlUnitMultiError) AllErrors() []error { return m }

// GetUsersURLsResp_UrlUnitValidationError is the validation error returned by
// GetUsersURLsResp_UrlUnit.Validate if the designated constraints aren't met.
type GetUsersURLsResp_UrlUnitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersURLsResp_UrlUnitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersURLsResp_UrlUnitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersURLsResp_UrlUnitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersURLsResp_UrlUnitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersURLsResp_UrlUnitValidationError) ErrorName() string {
	return "GetUsersURLsResp_UrlUnitValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersURLsResp_UrlUnitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersURLsResp_UrlUnit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersURLsResp_UrlUnitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersURLsResp_UrlUnitValidationError{}