package apis

import (
	"encoding/json"
	"io"

	// "github.com/showntop/tripper/errors"
)

var (
	params map[string]interface{}
)

// var (
// 	ErrParseRequestParams = errors.New("Bad jsonapi struct tag format")
// )

func ParseParams(data io.Reader) error {
	params = make(map[string]interface{})
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&params)
	if err != nil {
		return err //errors.NewServerError(err.Error())
	}
	return nil
}
