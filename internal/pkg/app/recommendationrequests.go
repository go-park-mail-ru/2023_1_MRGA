package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
)

type Recommendation struct {
	Username    string        `json:"username"`
	Avatar      string        `json:"avatar"`
	Age         int           `json:"age"`
	Sex         constform.Sex `json:"sex"`
	Description string        `json:"description"`
	City        string        `json:"city"`
}

func (a *Application) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log(http.StatusNotFound, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusNotFound)

		return
	}

	Stoken, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			http.Error(w, "error you are not authorised", http.StatusUnauthorized)

			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error", http.StatusInternalServerError)

		return
	}

	userId, err := a.repo.GetUserIdByToken(Stoken.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusBadRequest)

		return
	}

	recomendation, err := a.repo.GetRecommendation(userId)

	mapResp := make(map[string][]*Recommendation)
	mapResp["recommendations"] = recomendation

	w.Header().Add("Content-Type", "application/json")
	jsonData, err := json.Marshal(mapResp)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "cant create json", http.StatusInternalServerError)

		return
	}
	w.Write(jsonData)
}
