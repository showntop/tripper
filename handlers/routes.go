package handlers

import (
	"log"

	"github.com/julienschmidt/httprouter"
)

func init() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
