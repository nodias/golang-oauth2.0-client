package router

import (
	"github.com/nodias/golang-oauth2.0-common/models"
	"log"
	"net/http"
)


//templates
func loginHandler(w http.ResponseWriter, req *http.Request) {
	err := renderer.HTML(w, http.StatusOK, "index", &models.Response{
		Id:    "",
		User:  nil,
		Error: nil,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

