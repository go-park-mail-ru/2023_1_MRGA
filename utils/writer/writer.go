package writer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
)

type SuccessResult struct {
	Status int                    `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

const NameServise = "writer"

func Respond(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	result := SuccessResult{
		http.StatusOK,
		data,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)

	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
		_, err = w.Write([]byte(fmt.Sprintf(`{"status": %d, "err": "%s"}`, http.StatusInternalServerError, err.Error())))
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			return
		}
		return
	}
}

type ErrorResult struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func ErrorRespond(w http.ResponseWriter, r *http.Request, servarErr error, status int) {
	w.Header().Add("Content-Type", "application/json")

	result := ErrorResult{
		status,
		servarErr.Error(),
	}
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)

	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)

		_, err = w.Write([]byte(fmt.Sprintf(`{"status": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())))
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			return
		}
		return
	}
}

type ErrorResultWithData struct {
	Status int                    `json:"status"`
	Error  string                 `json:"error"`
	Data   map[string]interface{} `json:"body"`
}

func ErrorRespondWithData(w http.ResponseWriter, r *http.Request, servarErr error, status int, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	result := ErrorResultWithData{
		status,
		servarErr.Error(),
		data,
	}
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)

	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)

		_, err = w.Write([]byte(fmt.Sprintf(`{"status": %d, "error": "%s"}`, http.StatusInternalServerError, err.Error())))
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			return
		}
		return
	}
}
