package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

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

	var userJson dataStruct.User
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userToken, err := h.useCase.Register(&userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     _default.SessionTokenCookieName,
		Value:    userToken,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

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

	var logInp auth.LoginInput
	err = json.Unmarshal(reqBody, &logInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	userToken, err := h.useCase.Login(logInp)
	http.SetCookie(w, &http.Cookie{
		Name:     _default.SessionTokenCookieName,
		Value:    userToken,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
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

	err = h.useCase.Logout(userToken.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     _default.SessionTokenCookieName,
		Value:    "",
		Expires:  time.Now().Add(-120 * time.Second),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}
