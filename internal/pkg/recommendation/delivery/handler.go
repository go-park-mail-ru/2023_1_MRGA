package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

//func (h *Handler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
//
//	token, err := cookie.GetValueCookie(r, _default.SessionTokenCookieName)
//	if err != nil {
//		if err == http.ErrNoCookie {
//			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
//			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
//			return
//		}
//		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
//		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
//		return
//	}
//
//	recs, err := h.useCase.GetRecommendation(token)
//	if err != nil {
//		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
//		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
//		return
//	}
//	mapResp := make(map[string]interface{})
//	mapResp["recommendations"] = recs
//
//	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
//	writer.Respond(w, r, mapResp)
//}

func (h *Handler) AddFilter(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var filterInp recommendation.FilterInput
	err = json.Unmarshal(reqBody, &filterInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.AddFilters(uint(userId), filterInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) GetReasons(w http.ResponseWriter, r *http.Request) {
	reasons, err := h.useCase.GetReasons()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	result := make(map[string]interface{})
	result["reasons"] = reasons
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)

}

func (h *Handler) GetFilter(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}
	filters, err := h.useCase.GetFilters(uint(userId))
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	result := make(map[string]interface{})
	result["filters"] = filters
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)

}

func (h *Handler) ChangeFilter(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var filterInp recommendation.FilterInput
	err = json.Unmarshal(reqBody, &filterInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangeFilters(uint(userId), filterInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	filters, err := h.useCase.GetFilters(uint(userId))
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	result := make(map[string]interface{})
	result["filters"] = filters
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}
