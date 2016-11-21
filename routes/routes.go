package routes

import (
	"encoding/json"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/controllers"
)

func WrapErrorResp(err *controllers.HttpError) []byte {
	output := []byte(`{
		"message": "response json error",
		"state_code": 503
		}`)
	output, _ = json.Marshal(map[string]interface{}{
		"state_code": err.Code,
		"message":    err.Message,
	})

	return output
}

func Instrument() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/projects", listProject)
	router.POST("/api/v1/projects", createProject)
	router.GET("/api/v1/projects/:id", showProject)

	router.POST("/api/v1/qntokens", createQntoken)

	return router
}
