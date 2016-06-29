package apis

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/errors"
	"github.com/showntop/tripper/serializers"
)

func CreateSessionsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//validate allowed params key and value and type
	var login interface{}
	var ok bool
	if login, ok = params["login"]; ok {
		//string type
		if login, ok = login.(string); ok {

		} else {
			renderError(rw, NewParamsError("login param type error"), http.StatusBadRequest)
			return
		}
	} else {
		renderError(rw, NewParamsError("login param is needed"), http.StatusBadRequest)
		return
	}

	//logic
	user := store.User.FindByAny(params["login"].(string))
	if user.Id == "" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	matched := user.AuthPassword(params["password"].(string))
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
