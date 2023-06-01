package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) CreateInfo(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "CreateInfoHandler", nil)
	defer parentSpan.End()

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var infoInp info_user.InfoStruct
	err = json.Unmarshal(reqBody, &infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.AddInfo(uint(userId), infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) AddUserHashtags(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "AddUserHashtagsHandler", nil)
	defer parentSpan.End()

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var hashtagInp info_user.HashtagInp
	err = json.Unmarshal(reqBody, &hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.AddHashtags(uint(userId), hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetUserHashtags(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetUserHashtagsHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetUserStatus(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetUserStatusHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	statuses, err := h.useCase.GetUserStatus(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	result := make(map[string]interface{})
	result["status"] = statuses
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, result)
}

func (h *Handler) ChangeUserHashtags(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "ChangeUserHashtagsHandler", nil)
	defer parentSpan.End()

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var hashtagInp info_user.HashtagInp
	err = json.Unmarshal(reqBody, &hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangeUserHashtags(uint(userId), hashtagInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	hashtags, err := h.useCase.GetUserHashtags(uint(userId))
	var result info_user.HashtagInp
	result.Hashtag = hashtags
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&result)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) ChangeUserStatus(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "ChangeUserStatusHandler", nil)
	defer parentSpan.End()

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var statusInp info_user.StatusInp
	err = json.Unmarshal(reqBody, &statusInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangeUserStatus(uint(userId), statusInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) GetInfoById(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetInfoByIdHandler", nil)
	defer parentSpan.End()

	params := mux.Vars(r)
	userIdStr := params["userId"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	infoBD, err := h.useCase.GetInfo(uint(userId))
	infoBD.Email = ""
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&infoBD)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) GetInfo(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetInfoHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	infoBD, err := h.useCase.GetInfo(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&infoBD)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (h *Handler) ChangeInfo(w http.ResponseWriter, r *http.Request) {
	_, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "ChangeInfoHandler", nil)
	defer parentSpan.End()

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var infoInp info_user.InfoChange
	err = json.Unmarshal(reqBody, &infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	newInfo, err := h.useCase.ChangeInfo(uint(userId), infoInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	mapInfo := structs.Map(&newInfo)
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapInfo)
}

func (c *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	parentCtx, parentSpan := tracejaeger.NewSpan(r.Context(), "mainServer", "GetCurrentUserHandler", nil)
	defer parentSpan.End()

	userIdDB := r.Context().Value(middleware.ContextUserKey)
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	user, err := c.useCase.GetUserById(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	user.UserId = uint(userId)
	compInp := &complaintProto.UserId{UserId: userId}
	ctx, span := tracejaeger.NewSpan(parentCtx, "mainServer", "Complain", nil)
	banned, err := c.compService.CheckBanned(ctx, compInp)
	span.End()
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	user.Banned = banned.Banned

	mapUser := structs.Map(&user)
	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path, false)
	writer.Respond(w, r, mapUser)
}
