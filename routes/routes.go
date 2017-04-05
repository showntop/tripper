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

	router.GET("/api/v1/users/me/albums", listUserAlbum)

	router.GET("/api/v1/albums", listAlbum)
	router.GET("/api/v1/albums/:id", showAlbum)
	router.POST("/api/v1/albums", createAlbum)

	router.GET("/api/v1/categories", listCategory)

	router.GET("/api/v1/projects", listProject)
	router.POST("/api/v1/projects", createProject)
	router.GET("/api/v1/projects/:id", showProject)

	router.GET("/api/v1/projects/:id/comments", listProjectComment)
	router.POST("/api/v1/projects/:id/comments", createProjectComment)
	router.PUT("/api/v1/projects/:id/likes", createProjectLike)
	router.DELETE("/api/v1/projects/:id/likes", deleteProjectLike)

	router.GET("/api/v1/posts", listPost)
	router.POST("/api/v1/posts", createPost)
	router.GET("/api/v1/posts/:id", showPost)

	router.GET("/api/v1/topics", listTopic)
	router.POST("/api/v1/topics", createTopic)
	router.GET("/api/v1/topics/:id", showTopic)

	router.POST("/api/v1/qntokens", createQntoken)

	return router
}
