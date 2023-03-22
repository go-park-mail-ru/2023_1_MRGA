package delivery

import (
	"net/http"

	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetRecommendations(w http.ResponseWriter, r *http.Request) {

	token, err := r.Cookie(_default.SessionTokenCookieName)
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

	recs, err := h.useCase.GetRecommendation(token.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	mapResp := make(map[string]interface{})
	mapResp["recommendations"] = recs

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	writer.Respond(w, r, mapResp)
}
