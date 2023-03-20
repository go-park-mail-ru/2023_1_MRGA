package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

type Handler struct {
	useCase auth.UseCase
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

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

	if err := h.useCase.Register(userJson.Username, userJson.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	Respond(w, r, Result{http.StatusOK, ""}, map[string]interface{}{})
}
