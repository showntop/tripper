package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Categories struct {
	application
}

func (c *Categories) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	data, err := models.GetCategories()
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}
