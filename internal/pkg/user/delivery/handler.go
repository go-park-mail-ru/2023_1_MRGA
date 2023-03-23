package delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/structs"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

func (c *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Errorf("only GET method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
		return
	}

	UserToken, err := r.Cookie(SessionTokenCookieName)
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

	userId, err := a.repo.GetUserIdByToken(UserToken.Value)
	log.Println(userId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	user, err := a.repo.GetUserById(userId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	mapUser := structs.Map(&user)
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, mapUser)
}
