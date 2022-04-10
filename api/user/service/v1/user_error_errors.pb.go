// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, UserErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsLoginFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_LOGIN_FAILED.String() && e.Code == 401
}

func ErrorLoginFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_LOGIN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUsernameConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USERNAME_CONFLICT.String() && e.Code == 409
}

func ErrorUsernameConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, UserErrorReason_USERNAME_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsPhoneConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_PHONE_CONFLICT.String() && e.Code == 409
}

func ErrorPhoneConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, UserErrorReason_PHONE_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsEmailConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_EMAIL_CONFLICT.String() && e.Code == 409
}

func ErrorEmailConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, UserErrorReason_EMAIL_CONFLICT.String(), fmt.Sprintf(format, args...))
}
