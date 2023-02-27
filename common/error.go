package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrRecordNotFound = errors.New("record of not found")
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Key        string `json:"error_key"`
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewErrorResponse(statusCode int, root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
		Key:        key,
	}
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		http.StatusBadRequest,
		err,
		"something went wrong with DB",
		"ErrDB")
}

func ErrBadRequest(err error, msg string) *AppError {
	msg = strings.TrimSpace(msg)

	if msg == "" {
		msg = "your request is in bad format"
	}

	return NewErrorResponse(
		http.StatusInternalServerError,
		err,
		msg,
		"ErrBadRequest",
	)
}

func ErrInternalServer(err error, msg string) *AppError {
	msg = strings.TrimSpace(msg)

	if msg == "" {
		msg = "we encountered an error while processing your request"
	}

	return NewErrorResponse(
		http.StatusInternalServerError,
		err,
		msg,
		"ErrInternalServer",
	)
}

func ErrCannotCreateEntityResource(entity string, err error) *AppError {
	return NewErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("cannot create %s resource", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%sResource", entity),
	)
}

func ErrCannotGetEntityResource(entity string, err error) *AppError {
	return NewErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("cannot get %s resource", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%sResource", entity),
	)
}

func ErrCannotUpdateEntityResource(entity string, err error) *AppError {
	return NewErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("cannot get %s resource", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%sResource", entity),
	)
}

func ErrCannotDeleteEntityResource(entity string, err error) *AppError {
	return NewErrorResponse(
		http.StatusNotFound,
		err,
		fmt.Sprintf("cannot delete %s resource", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%sResource", entity),
	)
}
