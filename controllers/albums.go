package controllers

import (
	"encoding/json"
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

func (c *Albums) Create(req *http.Request) ([]byte, *HttpError) {
	var v models.Album
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}

	err = models.CreateAlbum(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}
