package server

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	cors "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/app/cors_middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/repository"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/service"
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(cors.SetCorsMiddleware)

	router.HandleFunc("/api/files/upload", uploadFile).Methods("POST")
	router.HandleFunc("/api/files/{id}", getFile).Methods("GET")
	return router
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Не передан id пользователя", http.StatusBadRequest)
		return
	}
	defer file.Close()

	userID := r.FormValue("userID")
	if userID == "" {
		http.Error(w, "Не передан id пользователя", http.StatusBadRequest)
		return
	}

	userIDUInt64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filePath, err := service.UploadFile(file, handler, uint(userIDUInt64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileID, err := repository.UploadFile(filePath, uint(userIDUInt64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", fileID)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	// Получаем ID файла из параметра запроса
	vars := mux.Vars(r)
	id := vars["id"]

	idUInt64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем содержимое файла из БД
	var file repository.File
	if file, err = repository.GetFile(uint(idUInt64)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gotFile, err := os.Open(file.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer gotFile.Close()

	filename := path.Base(gotFile.Name())
	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Автоматически ставит заголовок image/jpg или image/png
	http.ServeContent(w, r, filename, time.Now(), gotFile)
}
