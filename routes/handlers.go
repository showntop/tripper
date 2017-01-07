package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/controllers"
)

func listUserAlbum(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	usersC := new(controllers.Users)
	results, err := usersC.ListAlbums(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func listAlbum(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	albumsC := new(controllers.Albums)
	results, err := albumsC.List(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func createAlbum(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	albumsC := new(controllers.Albums)
	results, err := albumsC.Create(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func listProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.List(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func createProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.Create(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func showProject(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.Show(req, ps)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func listProjectComment(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {

}

func createProjectComment(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	projectsC := new(controllers.Projects)
	results, err := projectsC.CreateComment(req, ps)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func listTopic(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	postsC := new(controllers.Topics)
	results, err := postsC.List(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func createTopic(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	postsC := new(controllers.Topics)
	results, err := postsC.Create(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func showTopic(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	postsC := new(controllers.Topics)
	results, err := postsC.Show(req, ps)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func listPost(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// rw.Header().Set("Content-Type", "application/json")
	// postsC := new(controllers.Posts)
	// results, err := postsC.List(req)
	// if err != nil {
	// 	http.Error(rw, err.Error(), err.Code)
	// 	// rw.Write(WrapErrorResp(err))
	// 	return
	// }
	// rw.Write(results)
}

func createPost(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	postsC := new(controllers.Posts)
	results, err := postsC.Create(req)
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}

func showPost(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// rw.Header().Set("Content-Type", "application/json")
	// postsC := new(controllers.Posts)
	// results, err := postsC.Show(req, ps)
	// if err != nil {
	// 	http.Error(rw, err.Error(), err.Code)
	// 	// rw.Write(WrapErrorResp(err))
	// 	return
	// }
	// rw.Write(results)
}

func createQntoken(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	qntokensC := new(controllers.Qntokens)
	results, err := qntokensC.Create()
	if err != nil {
		http.Error(rw, err.Error(), err.Code)
		// rw.Write(WrapErrorResp(err))
		return
	}
	rw.Write(results)
}
