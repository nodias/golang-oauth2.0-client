package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"go.elastic.co/apm/module/apmgorilla"
	"net/http"
)

const (
	//fileServer
	fsPathPrefix = "/static/"
	fsDir        = "web/"

	//html
	htmlPathPrefix = "/"

	//userApi
	userApiPathPrefix = "/api/v1/users/"
	idPattern         = "/{id:[0-9]+}"
)

var renderer *render.Render
var templatesPath = fmt.Sprintf("%s/templates", fsDir)

func init() {
	renderer = render.New(render.Options{Directory: templatesPath})
}

func NewRouter() *mux.Router {
	return router()
}

func router() *mux.Router {
	r := mux.NewRouter()

	//file server router
	r.PathPrefix(fsPathPrefix).Handler(http.StripPrefix("/", http.FileServer(http.Dir(fsDir))))

	//template router
	tr := r.PathPrefix(htmlPathPrefix).Subrouter()
	tr.HandleFunc("/login", loginHandler).Methods("GET")

	//oauth2
	tr.HandleFunc("/reqauth", reqAuthHandler).Methods("GET")
	tr.HandleFunc("/oauth2", oauth2Handler).Methods("GET")
	tr.HandleFunc("/client", clientHandler).Methods("GET")
	tr.HandleFunc("/refresh", refreshHandler).Methods("GET")
	tr.HandleFunc("/try", tryHandler).Methods("GET")
	tr.HandleFunc("/pwd", pwdHandler).Methods("GET")


	//api router
	ar := r.PathPrefix(userApiPathPrefix).Subrouter()
	ar.HandleFunc(idPattern, getUserInfoHandler).Methods("GET")
	//ar.HandleFunc(idPattern, getUserInfoHandler).Methods("PUT")
	//ar.HandleFunc("/", getUserInfoHandler).Methods("POST")
	//ar.HandleFunc(idPattern, getUserInfoHandler).Methods("DELETE")
	ar.Use(apmgorilla.Middleware())
	return r
}
