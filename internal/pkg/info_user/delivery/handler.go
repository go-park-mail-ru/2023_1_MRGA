package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"

	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) CreateInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var infoInp info_user.InfoStruct
	err = json.Unmarshal(reqBody, &infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.AddInfo(uint(userId), infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, map[string]interface{}{})
	return
}

func (h *Handler) AddUserHashtags(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var hashtagInp info_user.HashtagInp
	err = json.Unmarshal(reqBody, &hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.AddHashtags(uint(userId), hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetUserHashtags(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) ChangeUserHashtags(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var hashtagInp info_user.HashtagInp
	err = json.Unmarshal(reqBody, &hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangeUserHashtags(uint(userId), hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetInfoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIdStr := params["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	infoBD, err := h.useCase.GetInfo(uint(userId))
	infoBD.Email = ""
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&infoBD)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, true)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	infoBD, err := h.useCase.GetInfo(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&infoBD)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) ChangeInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var infoInp info_user.InfoChange
	err = json.Unmarshal(reqBody, &infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	newInfo, err := h.useCase.ChangeInfo(uint(userId), infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&newInfo)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapInfo)
}

func (c *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	user, err := c.useCase.GetUserById(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	compInp := &complaintProto.UserId{UserId: userId}
	banned, err := c.compService.CheckBanned(r.Context(), compInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	user.Banned = banned.Banned

	mapUser := structs.Map(&user)
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, mapUser)
}
