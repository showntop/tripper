package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Projects struct {
	application
}

func (p *Projects) List(req *http.Request) ([]byte, *HttpError) {
	query := req.URL.Query()
	if recommended := query.Get("daily"); recommended != "" {
		return p.Recommend(req)
	}

	data, err := models.GetPorjectsSelected()
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}

func (p *Projects) Show(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	data, err := models.GetProjectById(ps.ByName("id"))
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}

func (p *Projects) Create(req *http.Request) ([]byte, *HttpError) {
	if err := p.AuthUser(req); err != nil {
		log.Warnln(err)
		return nil, UnAuthErr
	}
	var v models.Project
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}
	v.Author = p.CurrentUser
	err = models.CreateProject(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}

func (p *Projects) Recommend(req *http.Request) ([]byte, *HttpError) {

	data, err := models.GetProjectDaily()
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}
