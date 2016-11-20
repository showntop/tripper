package controllers

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type application struct {
}

type HttpError struct {
	Code    int
	Message string
}

func (h *HttpError) Error() string {
	return h.Message
}

var (
	BadRequestErr       = &HttpError{400, "json format error"}
	IncorrectAccountErr = &HttpError{402, "用户名或者密码错误"}

	ServerErr  = &HttpError{500, "server error"}
	DBErr      = &HttpError{503, "db error"}
	BadRespErr = &HttpError{503, "response format error"}
)

func WrapResp(data interface{}) ([]byte, *HttpError) {

	output, err := json.Marshal(data)
	if err != nil {
		log.Warnln("http response marshal| ,", err)
		return output, BadRespErr
	}
	return output, nil
}
