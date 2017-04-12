package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Projects struct {
	application
}

func (p *Projects) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	query := req.URL.Query()

	var data []*models.Project
	var err error
	if category := query.Get("category"); category != "" {
		categoryi, _ := strconv.Atoi(category)
		data, err = models.GetPorjectsByCategory(categoryi)
	} else {
		data, err = models.GetPorjectsSelected()
	}
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

func (p *Projects) Create(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
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
	if err := v.Validate(); err != nil {
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

func (p *Projects) Recommend(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	data, err := models.GetProjectDaily()
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(data)
}

func (p *Projects) CreateComment(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	projectId := ps.ByName("id")
	if projectId == "" {
		return nil, &HttpError{Code: 429, Message: "content not exist"}
	}
	var v map[string]string
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	if err != nil {
		return nil, BadRequestErr
	}
	comment := models.ProjectComment{ProjectId: projectId, Content: v["content"]}
	err = models.CreateProjectComment(&comment)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(comment)
}

func (p *Projects) CreateLike(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	if err := p.AuthUser(req); err != nil {
		log.Warnln(err)
		return nil, UnAuthErr
	}
	projectId := ps.ByName("id")
	if projectId == "" {
		return nil, &HttpError{Code: 429, Message: "content not exist"}
	}
	var v models.ProjectLike = models.ProjectLike{ProjectId: projectId, Liker: p.CurrentUser}

	err := models.CreateProjectLike(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}

func (p *Projects) DeleteLike(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	if err := p.AuthUser(req); err != nil {
		log.Warnln(err)
		return nil, UnAuthErr
	}
	projectId := ps.ByName("id")
	if projectId == "" {
		return nil, &HttpError{Code: 429, Message: "content not exist"}
	}
	var v models.ProjectLike = models.ProjectLike{ProjectId: projectId, Liker: p.CurrentUser}
	err := models.DeleteProjectLike(&v)
	if err != nil {
		log.Error(err)
		return nil, DBErr
	}
	return WrapResp(v)
}
