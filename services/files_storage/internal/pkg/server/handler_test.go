package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	myMocks "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/pkg/mocks"
)

type UploadPhoto struct {
	PhotoID uint `json:"photoID"`
}

type UploadPhotoResp struct {
	Status int         `json:"status"`
	Body   UploadPhoto `json:"body"`
}

func TestUploadPhoto(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := myMocks.NewMockIService(mockCtrl)

	// Создаем временный файл и записываем в него некоторые данные
	tmpfile, err := os.CreateTemp("", "example*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name()) // Удаляем файл после завершения теста

	if _, err := tmpfile.Write([]byte("testdata")); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	var userID uint = 1

	const futurePhotoID uint = 1
	mockService.EXPECT().UploadPhoto(gomock.Any(), filepath.Base(tmpfile.Name()), userID).Return(futurePhotoID, nil)

	server := Server{
		service: mockService,
	}

	// Создание *multipart.Writer
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// запись файла
	part, err := writer.CreateFormFile("file", filepath.Base(tmpfile.Name()))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, tmpfile)
	if err != nil {
		t.Fatal(err)
	}

	userIdField, err := writer.CreateFormField("userID")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.WriteString(userIdField, fmt.Sprintf("%v", userID))
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	// Создаем запрос
	request := httptest.NewRequest(http.MethodPost, "/api/v1/files/upload", &b)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// Создаем запись ответа
	recorder := httptest.NewRecorder()

	server.UploadPhoto(recorder, request)

	body := recorder.Body

	var uploadFileAnswer UploadPhotoResp
	err = json.NewDecoder(body).Decode(&uploadFileAnswer)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, UploadPhotoResp{Status: 200, Body: UploadPhoto{
		PhotoID: futurePhotoID}}, uploadFileAnswer)

}

type uploadFile struct {
	PathToFile string `json:"pathToFile"`
}

type uploadFileResp struct {
	Status int        `json:"status"`
	Body   uploadFile `json:"body"`
}

func TestUploadFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := myMocks.NewMockIService(mockCtrl)

	// Создаем временный файл и записываем в него некоторые данные
	tmpfile, err := os.CreateTemp("", "example*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()
	defer os.Remove(tmpfile.Name()) // Удаляем файл после завершения теста

	if _, err := tmpfile.Write([]byte("testdata")); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	var userID uint = 1

	const futurePathToFile string = "path/to/file"
	mockService.EXPECT().UploadFile(gomock.Any(), filepath.Base(tmpfile.Name()), userID).Return(futurePathToFile, nil)

	server := Server{
		service: mockService,
	}

	// Создание *multipart.Writer
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// запись файла
	part, err := writer.CreateFormFile("file", filepath.Base(tmpfile.Name()))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, tmpfile)
	if err != nil {
		t.Fatal(err)
	}

	userIdField, err := writer.CreateFormField("userID")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.WriteString(userIdField, fmt.Sprintf("%v", userID))
	if err != nil {
		t.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	// Создаем запрос
	request := httptest.NewRequest(http.MethodPost, "/api/v2/files/upload", &b)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// Создаем запись ответа
	recorder := httptest.NewRecorder()

	server.UploadFile(recorder, request)

	body := recorder.Body

	var uploadFileAnswer uploadFileResp
	err = json.NewDecoder(body).Decode(&uploadFileAnswer)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uploadFileResp{Status: 200, Body: uploadFile{
		PathToFile: futurePathToFile}}, uploadFileAnswer)
}
