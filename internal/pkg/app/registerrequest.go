package app

import (
	"crypto/sha1"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
	token2 "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/utils"
)

func (a *Application) register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		logger.Log(http.StatusNotFound, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusNotFound)

		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant read json", http.StatusBadRequest)

		return

	}

	var userJson ds.User
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant parse json", http.StatusBadRequest)

		return
	}

	hashedPass := CreatePass(userJson.Password)
	userJson.Password = hashedPass

	err = a.repo.AddUser(&userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	token := token2.CreateToken()
	a.repo.SaveToken(userJson.UserId, token)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	m := utils.Message(true, "success")
	utils.Respond(w, m)
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte{})

	return string(bs)
}
