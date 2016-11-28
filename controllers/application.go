package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/tripper/models"
)

const (
	SSO_URL = "http://127.0.0.1:7000/api/v1/users"
)

type application struct {
	CurrentUser *models.User
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
	UnAuthErr           = &HttpError{403, "用户验证错误"}

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

func (a *application) AuthUser(req *http.Request) error {

	////for test
	a.CurrentUser = &models.User{
		Id:       "test",
		Avatar:   "http://i2.letvimg.com/lc07_user/201605/09/15/14/1923203311462778080_70_70.jpg",
		Nickname: "黑色的风",
		Gender:   1,
	}
	return nil
	////

	ssoReq, err := http.NewRequest("GET", SSO_URL, nil)
	if err != nil {
		return err
	}
	http.DefaultClient.Timeout = 5 * time.Second
	ssoReq.Header.Add("Sun-Token", req.Header.Get("Sun-Token"))

	resp, err := http.DefaultClient.Do(ssoReq)
	if err != nil {
		return fmt.Errorf("auth user error, %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("auth user error, s")
	}
	log.Debugln(resp)
	decoder := json.NewDecoder(resp.Body)
	var user models.User
	err = decoder.Decode(&user)
	if err != nil {
		return fmt.Errorf("auth user error, %s", err.Error())
	}
	a.CurrentUser = &user
	return nil
}
