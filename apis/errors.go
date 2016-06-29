package apis

import (
	"errors"
)

type HttpError struct {
	error
}

func (he *HttpError) name() {
	he.Error()
}

type ParamsError struct {
	HttpError
}

type ValuesError struct {
	HttpError
}

type AuthError struct {
	HttpError
}

func NewParamsError(msg string) error {
	return &ParamsError{HttpError{errors.New(msg)}}
}
