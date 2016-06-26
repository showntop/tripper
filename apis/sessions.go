package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/errors"
	"github.com/showntop/tripper/serializers"
)

func CreateSessionsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("aaaaa1")

	//validate allowed params key and value and type
	user := store.User.FindByAny(params["login"])
	if user.Id == "" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	matched := user.AuthPassword(params["password"])
	if !matched {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(errors.NewAuthError("用户名或密码错误").ToJson()))
		return
	}
	response, err := serializers.MarshalObjectPayload(user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(errors.NewServerError(err.Error()).ToJson()))
		return
	}
	rw.Write(response)
}
