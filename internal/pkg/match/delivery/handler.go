package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	matches, err := h.useCase.GetMatches(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	result := make(map[string]interface{})
	result["matches"] = matches

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}

func (h *Handler) AddReaction(w http.ResponseWriter, r *http.Request) {
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

	var reaction match.ReactionInp
	err = json.Unmarshal(reqBody, &reaction)
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

	err = h.useCase.PostReaction(uint(userId), reaction)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
	return
}

func (h *Handler) GetChatByUserId(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(int)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	matchUserIdStr := params["userId"]
	matchUserId, err := strconv.Atoi(matchUserIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	chat, err := h.useCase.GetChatByEmail(uint(userId), uint(matchUserId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	result := structs.Map(&chat)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, result)
}
