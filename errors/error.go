package errors

import (
	"encoding/json"
	"errors"
)

type SunError interface {
	GetMsg() string
	ToJson() string
}

type baseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	error
}

func (be baseError) GetMsg() string {
	return be.error.Error()
}

func (be baseError) ToJson() string {
	be.Message = be.error.Error()
	b, err := json.Marshal(be)
	if err != nil {
		return ""
	} else {
		return string(b)
	}

}

////////////////////////////////////////////
type ModelError struct {
	baseError
}

type ValidateError struct {
	baseError
}

type DBError struct {
	baseError
}

type ServerError struct {
	baseError
}

type FileError struct {
	baseError
}

type AuthError struct {
	baseError
}

///////////////////////////////////////////////////////////////////////
func NewModelError(msg string) SunError {
	err := errors.New(msg)
	return ModelError{baseError{Error: "ModelError", error: err}}
}

func NewValidateError(msg string) SunError {
	err := errors.New(msg)
	return ValidateError{baseError{Error: "ValidateError", error: err}}
}

func NewServerError(msg string) SunError {
	err := errors.New(msg)
	return ServerError{baseError{Error: "ServerError", error: err}}
}

func NewDBError(msg string) SunError {
	err := errors.New(msg)
	return DBError{baseError{Error: "DBError", error: err}}
}

func NewFileError(msg string) SunError {
	err := errors.New(msg)
	return FileError{baseError{Error: "FileError", error: err}}
}

func NewAuthError(msg string) SunError {
	err := errors.New(msg)
	return AuthError{baseError{Error: "AuthError", error: err}}
}
