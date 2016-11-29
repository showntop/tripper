package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/tripper/models"
)

type Posts struct {
	application
}

func (p *Posts) Create(req *http.Request) ([]byte, *HttpError) {
	if err := p.AuthUser(req); err != nil {
		return nil, UnAuthErr
	}

	var v models.Post
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}
	if err := v.Validate(); err != nil {
		return nil, BadRequestErr
	}
	v.Author = p.CurrentUser
	err = models.CreatePost(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}
