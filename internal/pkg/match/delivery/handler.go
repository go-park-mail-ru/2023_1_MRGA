package delivery

import (
	"net/http"

	"github.com/fatih/structs"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/cookie"
	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetMatches(w http.ResponseWriter, r *http.Request) {

	userToken, err := cookie.GetValueCookie(r, _default.SessionTokenCookieName)

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
	repoAuth := repository.NewRepo(h.useCase.db)
	user, err := h.useCase.GetUserByToken(userToken)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapUser := structs.Map(&user)
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	writer.Respond(w, r, mapUser)
}
