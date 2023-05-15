package server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
	"github.com/gorilla/mux"
)

func (server Server) UploadFileV1(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID := r.FormValue("userID")
	if userID == "" {
		writer.ErrorRespond(w, r, errors.New("Не передан id пользователя"), http.StatusBadRequest)
		return
	}

	userIDUInt64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	fileID, err := server.service.UploadFileV1(file, fileHandler.Filename, uint(userIDUInt64))
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	writer.Respond(w, r, map[string]interface{}{
		"photoID": fileID,
	})
}

func (server Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID := r.FormValue("userID")
	if userID == "" {
		writer.ErrorRespond(w, r, errors.New("Не передан id пользователя"), http.StatusBadRequest)
		return
	}

	userIDUInt64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	pathToFile, err := server.service.UploadFile(file, fileHandler.Filename, uint(userIDUInt64))
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	writer.Respond(w, r, map[string]interface{}{
		"pathToFile": pathToFile,
	})
}

func (server Server) GetFile(w http.ResponseWriter, r *http.Request) {
	// Получаем ID файла из параметра запроса
	vars := mux.Vars(r)
	id := vars["id"]

	idUInt64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gotFile, filename, err := server.service.GetFile(uint(idUInt64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer gotFile.Close()

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))

	// Автоматически ставит заголовок image/jpg или image/png
	http.ServeContent(w, r, filename, time.Now(), gotFile)
}

func (server Server) GetFileByPath(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pathToFile := vars["pathToFile"]

	gotFile, filename, err := server.service.GetFileByPath(pathToFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer gotFile.Close()

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))

	// Автоматически ставит заголовок image/jpg или image/png
	http.ServeContent(w, r, filename, time.Now(), gotFile)
}
