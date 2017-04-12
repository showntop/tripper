package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Topics struct {
	application
}

func (p *Topics) Create(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	if err := p.AuthUser(req); err != nil {
		return nil, UnAuthErr
	}

	var v models.Topic
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}
	if err := v.Validate(); err != nil {
		return nil, BadRequestErr
	}
	v.Author = p.CurrentUser
	err = models.CreateTopic(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}

func (t *Topics) Show(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	v, err := models.GetTopicById(ps.ByName("id"))
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}

func (t *Topics) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	v, err := models.GetTopics()
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}
