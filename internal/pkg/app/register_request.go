package app

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"

	"github.com/fatih/structs"
)

type LoginInput struct {
	Input    string `json:"input"`
	Password string `json:"password"`
}

type Result struct {
	status int
	err    string
}

const SessionTokenCookieName = "session_token"

func Respond(w http.ResponseWriter, r *http.Request, res Result, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	data["status"] = res.status
	data["err"] = res.err

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)

	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		w.Write([]byte(fmt.Sprintf(`{"status": %d, "err": "%s"}`, http.StatusInternalServerError, err.Error())))
		return
	}
}

// Register godoc
// @Summary      Register new user
// @Description  create new account with unique username and email
// @Tags         Registration
//@Param Body body dataStruct.User true "info about user"
// @Success      200
// @Router       /meetme/register [post]
func (a *Application) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := "Only POST method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	defer r.Body.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	var userJson dataStruct.User
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	hashedPass := HashPassword(userJson.Password)
	userJson.Password = hashedPass

	if userJson.Avatar == "" {
		userJson.Avatar = "https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/User-avatar.svg/2048px-User-avatar.svg.png"
	}

	err = a.repo.AddUser(&userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	userToken := token.CreateToken()
	a.repo.SaveToken(userJson.UserId, userToken)

	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    userToken,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}

// Login godoc
// @Summary      authorise user
// @Description  authorise existing user with username/email and password
// @Tags         Registration
//@Param Body body LoginInput true "nickname/email password"
// @Success      200
// @Router       /meetme/login [post]
func (a *Application) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := "Only POST method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	defer r.Body.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	var logInp LoginInput
	err = json.Unmarshal(reqBody, &logInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	hashPass := HashPassword(logInp.Password)

	var userId uint

	userId, err = a.repo.Login(logInp.Input, hashPass)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	userToken := token.CreateToken()
	a.repo.SaveToken(userId, userToken)

	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    userToken,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}

// Logout godoc
// @Summary      Logout authorised user
// @Description  user can log out and end session
// @Tags         Registration
// @Success      200
// @Router       /meetme/logout [get]
func (a *Application) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := "Only POST method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	token, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusUnauthorized, err.Error()}, map[string]interface{}{})
			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
		return
	}

	strToken := token.Value
	err = a.repo.DeleteToken(strToken)
	if err != nil {
		logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusUnauthorized, err.Error()}, map[string]interface{}{})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    "",
		Expires:  time.Now().Add(-120 * time.Second),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}

// GetCities godoc
// @Summary      get list of cities for registration
// @Description  returned list of cities
// @Tags         Registration
// @Success      200  {object}  map[string][]string
// @Router       /meetme/city [get]
func (a *Application) GetCities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := "Only GET method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	cities, err := a.repo.GetCities()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
		return
	}

	mapResp := make(map[string]interface{})
	mapResp["city"] = cities

	Respond(w, r, Result{http.StatusOK, ""}, mapResp)
}

type UserRes struct {
	Username    string        `json:"username" structs:"username"`
	Email       string        `json:"email" structs:"email"`
	Age         int           `json:"age" structs:"age"`
	Sex         constform.Sex `json:"sex" structs:"sex"`
	City        string        `json:"city" structs:"city"`
	Description string        `json:"description" structs:"description"`
	Avatar      string        `json:"avatar" structs:"avatar"`
}

// GetCurrentUser godoc
// @Summary      get info about current user
// @Description  you can get info about current user
// @Tags         Registration
// @Produce      json
// @Success      200  {object}  UserRes
// @Router       /meetme/user [get]
func (a *Application) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := "Only GET method is supported for this route"
		logger.Log(http.StatusNotFound, err, r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err}, map[string]interface{}{})
		return
	}

	token, err := r.Cookie(SessionTokenCookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusUnauthorized, err.Error()}, map[string]interface{}{})
			return
		}
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
		return
	}

	userId, err := a.repo.GetUserIdByToken(token.Value)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	user, err := a.repo.GetUserById(userId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
		return
	}

	mapUser := structs.Map(&user)

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, mapUser)
}

func HashPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
