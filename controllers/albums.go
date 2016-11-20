package controllers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/tripper/models"
)

type Albums struct {
	*application
}

func (c *Albums) List(req *http.Request) ([]byte, *HttpError) {

	data, err := models.GetAlbums()
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}
