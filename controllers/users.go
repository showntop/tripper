package controllers

import (
	"net/http"

	"github.com/showntop/tripper/models"
	"gopkg.in/mgo.v2/bson"
	log "qiniupkg.com/x/log.v7"
)

type Users struct {
	application
}

func (c *Users) ListAlbums(req *http.Request) ([]byte, *HttpError) {
	if err := c.AuthUser(req); err != nil {
		return nil, UnAuthErr
	}

	data, err := models.GetAlbums(bson.M{"owner.id": c.CurrentUser.Id})
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}
