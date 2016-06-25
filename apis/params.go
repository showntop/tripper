package apis

import (
	"encoding/json"
	"io"

	"github.com/showntop/tripper/errors"
)

var (
	params map[string]string
)

func ParseParams(data io.Reader) errors.SunError {
	params = make(map[string]string)
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&params)
	if err != nil {
		return errors.NewServerError("parse req json error")
	}
	//validate allowed params key and value and type
	return nil
}
