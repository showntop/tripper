package apis

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	// "github.com/showntop/tripper/errors"
)

var (
	params map[string]interface{}
)

// var (
// 	ErrParseRequestParams = errors.New("Bad jsonapi struct tag format")
// )

func ParseJson(data io.Reader) error {
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

func ParseMultiPart2(req *http.Request) error {
	params = make(map[string]interface{})
	err := req.ParseMultipartForm(0)
	if err != nil {
		return err
	}
	log.Printf("%v", req.MultipartForm)
	for key, value := range req.MultipartForm.Value {
		params[key] = value
	}
	for key, file := range req.MultipartForm.File {
		params[key] = file
	}
	return nil
}

func ParseMultiPart(data *multipart.Reader) error {
	for {
		p, err := data.NextPart()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		contentType := p.Header.Get("Content-Type")
		switch {
		case strings.Contains(contentType, "application/json"):
			//ParseJson()
		}
		slurp, err := ioutil.ReadAll(p)
		if err != nil {
			log.Fatal(err)
			return err
		}
		_ = slurp
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
