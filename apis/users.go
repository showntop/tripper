package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/errors"
	"github.com/showntop/tripper/models"
	"github.com/showntop/tripper/serializers"
)

func CreateUsersHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("route handle")
	user, err := models.NewUser(params["mobile"], params["password"])

	err = store.User.Save(user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.ToJson()))
		return
	}
	response, errs := serializers.MarshalObjectPayload(user)
	if errs != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(errors.NewServerError(errs.Error()).ToJson()))
		return
	}
	rw.Write(response)
}
