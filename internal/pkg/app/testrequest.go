package app

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/logger"
)

func (a *Application) Return500(w http.ResponseWriter, r *http.Request) {
	logger.Log(http.StatusInternalServerError, "Give 500", r.Method, r.URL.Path)
	http.Error(w, "error method", http.StatusInternalServerError)
}
