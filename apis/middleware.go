package apis

import (
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	httprouter *httprouter.Router
}

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var err error
	log.Println("===========================")
	log.Println(req.Method + "------" + req.URL.Path)
	log.Printf("%v", req)
	contentType := req.Header.Get("Content-Type")
	switch {
	case strings.Contains(contentType, "multipart/form-data"):
		err = ParseMultiPart2(req)
		//_ = "breakpoint"

		// reader, err := req.MultipartReader()
		// if err != nil {
		// 	log.Printf("%v", err)
		// 	renderError(rw, err, http.StatusInternalServerError)
		// 	return
		// }
		// ParseMultiPart(reader)
	case strings.Contains(contentType, "application/json"):
		err = ParseJson(req.Body)
	}
	if err != nil {
		log.Printf("%v", err)
		renderError(rw, err, http.StatusInternalServerError)
		return
	}

	log.Printf("%v", params)
	m.httprouter.ServeHTTP(rw, req)
	log.Println("===========================")
}
