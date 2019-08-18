package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nodias/golang-oauth2.0-client/service"
	"github.com/nodias/golang-oauth2.0-common/models"
	"github.com/nodias/golang-oauth2.0-common/shared/logger"
	"go.elastic.co/apm"
	"net/http"
	"strings"
	"unicode"
)

//getUserInfoHandler is a function, gets the information of one User
func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log := logger.New(ctx)

	id := mux.Vars(req)["id"]
	log.WithField("id", id).Debug("handling hello request")
	if strings.IndexFunc(id, func(r rune) bool { return r >= unicode.MaxASCII }) >= 0 {
		panic("non-ASCII id!")
	}

	user, rerr := service.GetUserInfo(req.Context(), id)
	if rerr != nil {
		w.WriteHeader(rerr.Code)
	}
	err := json.NewEncoder(w).Encode(models.Response{
		Id:    models.ID(id),
		User:  user,
		Error: rerr,
	})
	if err != nil {
		rerr2 := models.NewResponseError(err, 500)
		log.WithError(rerr2).Error("failed to GetUserInfoHandler")
		//apm server에 에러를 업로드 시켜줍니다.
		apm.CaptureError(ctx, rerr2.Err).Send()
		http.Error(w, "failed encode to json", 9999)
		return
	}
}
