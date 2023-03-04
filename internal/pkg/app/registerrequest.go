package app

import (
	"crypto/sha1"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/utils"
)

type LoginInput struct {
	Input    string `json:"input"`
	Password string `json:"password"`
}

const SessionTokenCookieName = "session_token"

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
	userToken := token.CreateToken()
	a.repo.SaveToken(userJson.UserId, userToken)
	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    userToken,
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

	if logInp.Input != "" {
		userId, err = a.repo.Login(logInp.Input, hashPass)
		if err != nil {
			logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
			http.Error(w, "error cant login", http.StatusBadRequest)

			return
		}
	} else {
		logger.Log(http.StatusBadRequest, "email and username are empty", r.Method, r.URL.Path)
		http.Error(w, "error cant login", http.StatusBadRequest)

		return
	}

	userToken := token.CreateToken()
	a.repo.SaveToken(userId, userToken)
	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    userToken,
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

	Stoken, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			http.Error(w, "error you are not authorised", http.StatusUnauthorized)

			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
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
		Name:     SessionTokenCookieName,
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
	cities, err := a.repo.GetCities()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error file cant read cities", http.StatusInternalServerError)

		return
	}
	mapResp := make(map[string][]string)
	mapResp["city"] = cities

	w.Header().Add("Content-Type", "application/json")
	jsonData, err := json.Marshal(mapResp)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "cant create json", http.StatusInternalServerError)

		return
	}
	w.Write(jsonData)
}

type UserRes struct {
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Age         int           `json:"age"`
	Sex         constform.Sex `json:"sex"`
	City        string        `json:"city"`
	Description string        `json:"description"`
	Avatar      string        `json:"avatar"`
}

func (a *Application) GetUserByCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Log(http.StatusNotFound, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusNotFound)

		return
	}

	Stoken, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			http.Error(w, "error you are not authorised", http.StatusUnauthorized)

			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		http.Error(w, "error", http.StatusInternalServerError)

		return
	}

	userId, err := a.repo.GetUserIdByToken(Stoken.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusBadRequest)

		return
	}
	user, err := a.repo.GetUserById(userId)
	if err != nil {
		logger.Log(http.StatusBadRequest, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusBadRequest)

		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		logger.Log(http.StatusInternalServerError, "Wrong method", r.Method, r.URL.Path)
		http.Error(w, "error method", http.StatusInternalServerError)

		return
	}

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonUser)

}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
