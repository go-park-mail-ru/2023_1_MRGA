package app

import (
	// "encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
)

type Recommendation struct {
	Username    string        `json:"username" structs:"username"`
	Avatar      string        `json:"avatar" structs:"avatar"`
	Age         int           `json:"age" structs:"age"`
	Sex         constform.Sex `json:"sex" structs:"sex"`
	Description string        `json:"description" structs:"description"`
	City        string        `json:"city" structs:"city"`
}

func (a *Application) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := "Only GET method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	token, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusUnauthorized, err.Error()}, map[string]interface{}{})
			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
		return
	}

	userId, err := a.repo.GetUserIdByToken(token.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	recomendation, err := a.repo.GetRecommendation(userId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	mapResp := make(map[string]interface{})
	mapResp["recommendations"] = recomendation

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, mapResp)
}
