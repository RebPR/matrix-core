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
	return e.Reason == CreationErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsCreateArticleDraftFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CreationErrorReason_CREATE_ARTICLE_DRAFT_FAILED.String() && e.Code == 500
}

func ErrorCreateArticleDraftFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CreationErrorReason_CREATE_ARTICLE_DRAFT_FAILED.String(), fmt.Sprintf(format, args...))
}
