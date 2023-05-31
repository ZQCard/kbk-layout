// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package example

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsRecordNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RECORD_NOT_FOUND.String() && e.Code == 404
}

func ErrorRecordNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_RECORD_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsRecordAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RECORD_ALREADY_EXISTS.String() && e.Code == 400
}

func ErrorRecordAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_RECORD_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BAD_REQUEST.String() && e.Code == 400
}

func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsSystemError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SYSTEM_ERROR.String() && e.Code == 500
}

func ErrorSystemError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SYSTEM_ERROR.String(), fmt.Sprintf(format, args...))
}
