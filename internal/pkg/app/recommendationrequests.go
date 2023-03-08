package app

import (
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

// GetRecommendations godoc
// @Summary      return recommendations for user
// @Description  now just return other 10 or fewer users
// @Tags         Tests
//@Param Body body []Recommendation true "list of users"
// @Success      200
// @Router       /meetme/recommendations [get]
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
