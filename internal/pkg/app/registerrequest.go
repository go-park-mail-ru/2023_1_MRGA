package app

import (
	"crypto/sha1"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
	token2 "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/utils"
)

type LoginInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Application) Register(w http.ResponseWriter, r *http.Request) {

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
	rsp := utils.Response(true, "success")
	utils.Respond(w, rsp)
}

func (a *Application) Login(w http.ResponseWriter, r *http.Request) {
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

	var logInp LoginInput
	err = json.Unmarshal(reqBody, &logInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant parse json", http.StatusBadRequest)

		return
	}

	hashPass := CreatePass(logInp.Password)

	var userId uint

	if logInp.Email != "" {
		userId, err = a.repo.LoginEmail(logInp.Email, hashPass)
	} else if logInp.Username != "" {
		userId, err = a.repo.LoginUsername(logInp.Username, hashPass)
	} else {
		logger.Log(http.StatusBadRequest, "email and username are empty", r.Method, r.URL.Path)
		http.Error(w, "error cant login", http.StatusBadRequest)
	}
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant login", http.StatusBadRequest)
	}

	token := token2.CreateToken()
	a.repo.SaveToken(userId, token)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
	})
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	rsp := utils.Response(true, "success")
	utils.Respond(w, rsp)
}

func (a *Application) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log(http.StatusNotFound, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusNotFound)

		return
	}

	Stoken, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			http.Error(w, "error you are not authorised", http.StatusUnauthorized)
			return
		}
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	Strtoken := Stoken.Value
	err = a.repo.DeleteToken(Strtoken)

	if err != nil {
		logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error you are not authorised", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-120 * time.Second),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	rsp := utils.Response(true, "success")
	utils.Respond(w, rsp)
}

func (a *Application) GetCities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log(http.StatusNotFound, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusNotFound)

		return
	}

	fileCity, err := os.Open("./files/city.txt")
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant read json", http.StatusInternalServerError)

		return
	}

	allCities, err := io.ReadAll(fileCity)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error cant read json", http.StatusInternalServerError)

		return
	}

	allCitiesStr := string(allCities)
	cities := strings.Split(allCitiesStr, "\n")
	mapResp := make(map[string][]string)
	mapResp["city"] = cities

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mapResp)
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
