package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/controllers"
)

func listProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.List(req)
	if err != nil {
		// http.Error(rw, err.Error(), err.Code)
		rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func createProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.Create(req)
	if err != nil {
		// http.Error(rw, err.Error(), err.Code)
		rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func showProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.Show(req, ps)
	if err != nil {
		// http.Error(rw, err.Error(), err.Code)
		rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func createQntoken(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	qntokensC := new(controllers.Qntokens)
	results, err := qntokensC.Create()
	if err != nil {
		// http.Error(rw, err.Error(), err.Code)
		rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}
