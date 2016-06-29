package apis

import (
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
	"github.com/showntop/tripper/serializers"
)

func ListSpotsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

}

func CreateSpotsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var err error
	//validate params
	err = Require("title", reflect.String)
	if err != nil {
		renderError(rw, err, http.StatusBadRequest)
		return
	}
	err = Optional("description", reflect.String)
	if err != nil {
		renderError(rw, err, http.StatusBadRequest)
		return
	}
	//logic
	spot := models.NewSpot(params["title"].(string), params["description"].(string))
	err = store.Spot.Save(spot)
	if err != nil {
		renderError(rw, err, http.StatusInternalServerError)
		return
	}
	data, _ := serializers.MarshalObjectPayload(spot)
	render(rw, data)
	//render
}
