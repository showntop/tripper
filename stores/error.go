package stores

import (
	"errors"
)

type StoreError struct {
	error
}

func (se *StoreError) Error() string {
	return se.Error()
}

func NewStoreError(msg string) error {
	return &StoreError{errors.New(msg)}
}
