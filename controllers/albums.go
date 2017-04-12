package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	"github.com/showntop/tripper/models"
)

type Albums struct {
	application
}

func (c *Albums) Show(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	data, err := models.GetAlbumById(bson.ObjectIdHex(ps.ByName("id")))
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}

func (c *Albums) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	data, err := models.GetAlbums(nil)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}

func (c *Albums) Create(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	if err := c.AuthUser(req); err != nil {
		return nil, UnAuthErr
	}

	var v models.Album
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}

	v.Owner = c.CurrentUser

	err = models.CreateAlbum(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}
