package routes

import (
	"encoding/json"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/controllers"
)

func WrapErrorResp(err *controllers.HttpError) string {
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

func Instrument() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/albums", listAlbum)
	router.POST("/api/v1/albums", createAlbum)

	router.GET("/api/v1/projects", listProject)
	router.POST("/api/v1/projects", createProject)
	router.GET("/api/v1/projects/:id", showProject)

	router.POST("/api/v1/qntokens", createQntoken)

	return router
}
