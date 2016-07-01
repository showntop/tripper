package apis

import (
	"encoding/json"
	"io"
	"reflect"
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
	switch {
	case err == io.EOF: //empty body
		return nil
	case err != nil:
		return err
	}
	return nil
}

func Require(fieldName string, fieldType reflect.Kind, args ...interface{}) error {
	if _, ok := params[fieldName]; !ok {
		return NewParamsError("params" + fieldName + "is missing")
	}
	if reflect.TypeOf(params[fieldName]).Kind() != fieldType {
		return NewParamsError("params" + fieldName + "type is wrong")
	}
	return nil
	// ruleTag = args[0]
	// ruleFun = args[1]
}

func Optional(fieldName string, fieldType reflect.Kind, args ...interface{}) error {
	if _, ok := params[fieldName]; !ok {
		params[fieldName] = ""
		return nil
	}
	if reflect.TypeOf(params[fieldName]).Kind() != fieldType {
		return NewParamsError("params" + fieldName + "type is wrong")
	}
	return nil
	// ruleTag = args[0]
	// ruleFun = args[1]
}
