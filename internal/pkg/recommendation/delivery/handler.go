package delivery

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetRecommendationsHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	recs, err := h.useCase.GetRecommendations(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapResp := make(map[string]interface{})
	mapResp["recommendations"] = recs
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapResp)
}

func (h *Handler) GetLikes(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetLikesHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err := h.useCase.CheckProStatus(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("you do not have pro status")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	likes, err := h.useCase.GetLikes(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	result := make(map[string]interface{})
	result["likes"] = likes

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}
