package app

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/structs"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
)

type LoginInput struct {
	Input    string `json:"input"`
	Password string `json:"password"`
}

const SessionTokenCookieName = "session_token"

type Result struct {
	status int
	err    string
}

func Respond(w http.ResponseWriter, r *http.Request, res Result, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	data["status"] = res.status
	data["err"] = res.err

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)

	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		_, err = w.Write([]byte(fmt.Sprintf(`{"status": %d, "err": "%s"}`, http.StatusInternalServerError, err.Error())))
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			return
		}
		return
	}
}

// Register godoc
// @Summary      Register new user
// @Description  create new account with unique username and email
// @Tags         Registration
// @Param Body body dataStruct.User true "info about user"
// @Success      200 {object} map[string]interface{}
// @Router       /meetme/register [post]
func (a *Application) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		err := fmt.Errorf("only POST method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
		return
	}

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
			return
		}
	}()
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
		Respond(w, r, Result{http.StatusBadRequest, "cant parse json"}, map[string]interface{}{})
		return
	}

	hashedPass := CreatePass(userJson.Password)
	userJson.Password = hashedPass

	if userJson.Avatar == "" {
		userJson.Avatar = "https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/User-avatar.svg/2048px-User-avatar.svg.png"
	}

	err = a.repo.AddUser(userJson)
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
// @Param Body body LoginInput true "nickname/email password"
// @Success      200 {object} map[string]interface{}
// @Router       /meetme/login [post]
func (a *Application) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := fmt.Errorf("only POST method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
		return
	}

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusInternalServerError, err.Error()}, map[string]interface{}{})
			return
		}
	}()
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
		Respond(w, r, Result{http.StatusBadRequest, "cant parse json"}, map[string]interface{}{})
		return
	}

	hashPass := CreatePass(logInp.Password)

	var userId uint

	if logInp.Input != "" {
		userId, err = a.repo.Login(logInp.Input, hashPass)
		if err != nil {
			logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
			Respond(w, r, Result{http.StatusBadRequest, err.Error()}, map[string]interface{}{})
			return
		}
	} else {
		logger.Log(http.StatusBadRequest, "email and username are empty", r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusBadRequest, "empty login"}, map[string]interface{}{})
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
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}

// Logout godoc
// @Summary      Logout authorised user
// @Description  user can log out and end session
// @Tags         Registration
// @Success      200 {object} map[string]interface{}
// @Router       /meetme/logout [post]
func (a *Application) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := fmt.Errorf("only POST method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
		return
	}

	userToken, err := r.Cookie(SessionTokenCookieName)
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

	strToken := userToken.Value
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
	})

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}

// GetCities godoc
// @Summary      get list of cities for registration
// @Description  returned list of cities
// @Tags         info
// @Success      200  {object}  map[string][]string
// @Router       /meetme/city [get]
func (a *Application) GetCities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Errorf("only GET method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
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
	Username    string        `structs:"username"`
	Email       string        `structs:"email"`
	Age         int           `structs:"age"`
	Sex         constform.Sex `structs:"sex"`
	City        string        `structs:"city"`
	Description string        `structs:"description"`
	Avatar      string        `structs:"avatar"`
}

// GetCurrentUser godoc
// @Summary      get info about current user
// @Description  you can get info about current user
// @Tags         info
// @Produce      json
// @Success      200  {object}  UserRes
// @Router       /meetme/user [get]
func (a *Application) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := fmt.Errorf("only GET method is supported for this route")
		logger.Log(http.StatusNotFound, err.Error(), r.Method, r.URL.Path)
		Respond(w, r, Result{http.StatusNotFound, err.Error()}, map[string]interface{}{})
		return
	}

	UserToken, err := r.Cookie(SessionTokenCookieName)
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

	userId, err := a.repo.GetUserIdByToken(UserToken.Value)
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
	fmt.Print(mapUser)

	logger.Log(http.StatusOK, "give user information", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, mapUser)
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
