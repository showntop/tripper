package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Feeds struct {
	application
}

func (c *Feeds) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	data, err := models.ListFeeds()
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}
