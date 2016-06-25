package apis

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/showntop/tripper/stores"
)

var (
	store *stores.Store
)

func Setup() {
	//init the store
	store = stores.NewStore()
	log.Println("the store started...")

	router := httprouter.New()
	router.POST("/", Home)
	router.POST("/Users", CreateUsersHandler)
	router.POST("/Sessions", CreateSessionsHandler)
	router.GET("/Spots", ListSpotsHandler)

	m := &Middleware{router}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Fatal(http.ListenAndServe(":"+port, m))
}

func Home(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	rw.Write([]byte("this is the root path"))
}
