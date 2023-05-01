package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/logger"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/writer"
)

func (h *Handler) AddPhoto(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	err := r.ParseMultipartForm(32 << 20) // 32MB is the default size limit for a request
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["files[]"]
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	for idx, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}

		defer func() {
			err := file.Close()
			if err != nil {
				logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
				writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
				return
			}
		}()

		photoId, err := SendPhoto(file, fileHeader.Filename, uint(userId))
		if err != nil {
			logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
			err = fmt.Errorf("cant parse json")
			writer.ErrorRespond(w, r, err, http.StatusBadRequest)
			return
		}

		var avatar bool
		if idx == 0 {
			avatar = true
		} else {
			avatar = false
		}

		err = h.useCase.SavePhoto(uint(userId), photoId, avatar)
		if err != nil {
			logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusBadRequest)
			return
		}
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) GetPhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoId, ok := params["photo"]
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}
	bodyBytes, err := SendRequest(photoId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	body := &bytes.Buffer{}
	writerFile := multipart.NewWriter(body)
	fileField, err := writerFile.CreateFormFile("file", "filename")
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	_, err = fileField.Write(bodyBytes)
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	err = writerFile.Close()
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", writerFile.FormDataContentType())
	w.Header().Set("Content-Disposition", "attachment")
	_, err = w.Write(body.Bytes())
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoIdStr, ok := params["photo"]
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}
	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	err = h.useCase.DeletePhoto(uint(userId), photoId)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
	return
}

func (h *Handler) ChangePhoto(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()
	err := r.ParseMultipartForm(32 << 20) // 32MB is the default size limit for a request
	if err != nil {
		logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		logger.Log(http.StatusBadRequest, "", r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, nil, http.StatusBadRequest)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()
	params := mux.Vars(r)
	photoNumStr := params["photo"]
	photoNum, err := strconv.Atoi(photoNumStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	photoId, err := SendPhoto(file, "file", uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangePhoto(photoNum, photoId, uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path)
	writer.Respond(w, r, map[string]interface{}{})
}

func SendPhoto(file multipart.File, filename string, userID uint) (uint, error) {

	requestBody := &bytes.Buffer{}
	writerFile := multipart.NewWriter(requestBody)
	userIdField, err := writerFile.CreateFormField("userID")
	if err != nil {
		return 0, err
	}
	_, err = io.WriteString(userIdField, fmt.Sprintf("%v", userID))
	if err != nil {
		return 0, err
	}

	fileField, err := writerFile.CreateFormFile("file", filename)
	if err != nil {
		return 0, err
	}
	_, err = io.Copy(fileField, file)
	if err != nil {
		return 0, err
	}

	err = writerFile.Close()
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/api/files/upload", requestBody)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", writerFile.FormDataContentType())

	// Отправляем запрос и проверяем статус ответа
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status %v", resp.StatusCode)
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()

	answerBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var answer photo.AnswerPhoto
	err = json.Unmarshal(answerBody, &answer)
	if err != nil {
		return 0, err
	}

	return answer.Body.PhotoID, nil
}

func SendRequest(photoId string) ([]byte, error) {
	// Создаем HTTP-запрос на другой микросервис
	req, err := http.NewRequest("GET", "http://localhost:8081/api/files/"+photoId, nil)
	if err != nil {
		return nil, err
	}

	// Отправляем запрос и проверяем статус ответа
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, err
}
