package routes

import (
	"encoding/json"
	"net/http"
	"reflect"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	c "github.com/showntop/tripper/controllers"
)

func WrapErrorResp(err *c.HttpError) string {
	output := []byte(`{
		"message": "response json error",
		"status": 503
		}`)
	output, _ = json.Marshal(map[string]interface{}{
		"status":  err.Code,
		"message": err.Message,
	})

	return string(output)
}

func construct(controller interface{}, action string) func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	return func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")

		param1 := reflect.ValueOf(req)
		param2 := reflect.ValueOf(ps)

		rC := reflect.ValueOf(controller)
		results := rC.MethodByName(action).Call([]reflect.Value{param1, param2})
		log.Debugf("controller:%T, action:%s", controller, action)
		err := results[1].Interface().(*c.HttpError)
		if err != nil {
			http.Error(rw, err.Error(), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results[0].Interface().([]byte))
	}
}

func Instrument() *httprouter.Router {
	router := httprouter.New()

	// router.GET("/api/v1/feeds", construct(new(controller.Feeds, ""listFeed"")))

	router.GET("/api/v1/users/me/albums", construct(new(c.Users), "ListAlbums"))

	router.GET("/api/v1/albums", construct(new(c.Albums), "List"))
	router.GET("/api/v1/albums/:id", construct(new(c.Albums), "Show"))
	router.POST("/api/v1/albums", construct(new(c.Albums), "Create"))

	router.GET("/api/v1/categories", construct(new(c.Categories), "List"))

	router.GET("/api/v1/projects", construct(new(c.Projects), "List"))
	router.POST("/api/v1/projects", construct(new(c.Projects), "Create"))
	router.GET("/api/v1/projects/:id", construct(new(c.Projects), "Show"))

	router.GET("/api/v1/projects/:id/comments", construct(new(c.Projects), "listProjectComment"))
	router.POST("/api/v1/projects/:id/comments", construct(new(c.Projects), "CreateComment"))
	router.PUT("/api/v1/projects/:id/likes", construct(new(c.Projects), "CreateLike"))
	router.DELETE("/api/v1/projects/:id/likes", construct(new(c.Projects), "DeleteLike"))

	router.GET("/api/v1/posts", construct(new(c.Posts), "List"))
	router.POST("/api/v1/posts", construct(new(c.Posts), "Create"))
	router.GET("/api/v1/posts/:id", construct(new(c.Posts), "Show"))

	router.GET("/api/v1/topics", construct(new(c.Topics), "List"))
	router.POST("/api/v1/topics", construct(new(c.Topics), "Create"))
	router.GET("/api/v1/topics/:id", construct(new(c.Topics), "Show"))

	router.POST("/api/v1/qntokens", construct(new(c.Qntokens), "Create"))

	return router
}
