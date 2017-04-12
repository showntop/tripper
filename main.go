package main

import (
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"

	. "github.com/showntop/tripper/config"
	"github.com/showntop/tripper/routes"
)

type Filter func(rw http.ResponseWriter, req *http.Request) bool
type Middleware struct {
	mmux          *http.ServeMux
	BeforeFilters []Filter
	// AfterFilters  []*Filter//TODO
}

func (m *Middleware) Route(method, path, controller, action string) {

}

func (m *Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	dump, _ := httputil.DumpRequest(req, true)

	log.Infoln()
	startTime := time.Now()
	for i, j := 0, len(m.BeforeFilters); i < j; i++ {
		funcName := runtime.FuncForPC(reflect.ValueOf(m.BeforeFilters[i]).Pointer()).Name()
		if rendered := m.BeforeFilters[i](rw, req); rendered {
			log.Infof("filted by %v", funcName)
			return
		}
	}
	m.mmux.ServeHTTP(rw, req)
	usedTime := time.Since(startTime)

	log.Infof("%s %s %s for %s used time: %s \n%s\n", time.Now().String(), req.Method, req.URL.Path, req.Host, usedTime, dump)
}

func (m *Middleware) AddBeforeFilter(f Filter) {
	m.BeforeFilters = append(m.BeforeFilters, f)
}

func main() {
	// init log
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		log.WithFields(log.Fields{
	// 			"gallery": true,
	// 			"err":     err,
	// 			"number":  100,
	// 		}).Fatal("The server breaks!")
	// 	}
	// }()
	log.SetFormatter(&log.TextFormatter{})

	if Config.Env == "develop" {
		log.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(Config.LogPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Error(err)
		}
		log.SetOutput(f)
	}
	log.SetLevel(Config.LogLevel)
	//init backend

	// init http server
	ms := &Middleware{mmux: http.DefaultServeMux}
	// ms.AddBeforeFilter(Filter(handlers.ParseLocale))
	// ms.AddBeforeFilter(Filter(handlers.AuthUser))
	mux := ms.mmux
	mux.Handle("/", routes.Instrument())
	log.WithField("time", time.Now()).Infof("starting in development on %s", Config.ServerPort)
	log.Fatal(http.ListenAndServe(Config.ServerPort, ms))
}
