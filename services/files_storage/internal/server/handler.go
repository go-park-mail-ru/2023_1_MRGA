package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
	"github.com/gorilla/mux"
)

func (server *Server) getRouter() *mux.Router {
	router := mux.NewRouter()

	// router.Use(SetCorsMiddleware)
	router.HandleFunc("/api/files/upload", server.uploadFile).Methods("POST")
	router.HandleFunc("/api/files/{id}", server.getFile).Methods("GET")

	return router
}

func (server Server) uploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath, err := server.service.UploadFile(file, fileHandler)
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	fileID, err := server.repo.UploadFile(filePath)
	if err != nil {
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	writer.Respond(w, r, map[string]interface{}{
		"photoID": fileID,
	})
}

func (server Server) getFile(w http.ResponseWriter, r *http.Request) {
	// Получаем ID файла из параметра запроса
	vars := mux.Vars(r)
	id := vars["id"]

	idUInt64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем содержимое файла из БД
	var filePath string
	if filePath, err = server.repo.GetFile(uint(idUInt64)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gotFile, filename, err := server.service.GetFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer gotFile.Close()

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Автоматически ставит заголовок image/jpg или image/png
	http.ServeContent(w, r, filename, time.Now(), gotFile)
}
