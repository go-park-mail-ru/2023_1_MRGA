package delivery

import (
	"net/http"

	"github.com/fatih/structs"

	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (c *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

	userToken, err := r.Cookie(_default.SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	user, err := c.useCase.GetUserByToken(userToken.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapUser := structs.Map(&user)
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	writer.Respond(w, r, mapUser)
}
