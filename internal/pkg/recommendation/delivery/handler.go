package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetRecommendations(w http.ResponseWriter, r *http.Request) {

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	recs, err := h.useCase.GetRecommendations(uint(userId))
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
