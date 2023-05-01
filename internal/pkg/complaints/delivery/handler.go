package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaints"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) Complain(w http.ResponseWriter, r *http.Request) {
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

	var logInp complaints.UserId
	err = json.Unmarshal(reqBody, &logInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	_, err = h.CompService.Complain(r.Context(), &logInp)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, _default.NameService, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, _default.NameService, false)
	writer.Respond(w, r, map[string]interface{}{})
}
