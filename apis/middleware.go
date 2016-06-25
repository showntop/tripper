package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	httprouter *httprouter.Router
}

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("===========================")
	log.Println(req.Method + "------" + req.URL.Path)
	err := ParseParams(req.Body)
	rw.Header().Set("Content-Type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.ToJson()))
		return
	}
	log.Printf("%v", params)
	m.httprouter.ServeHTTP(rw, req)
	log.Println("===========================")
}
