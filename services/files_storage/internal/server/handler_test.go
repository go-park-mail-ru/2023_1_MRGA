package server

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type uploadFileResp struct {
	Status int `json:"status"`
	Body   struct {
		PhotoID uint `json:"photoID"`
	} `json:"body"`
}

func TestUploadFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepository := mocks.NewMockIRepository(mockCtrl)
	mockService := mocks.NewMockIService(mockCtrl)

	// Создаем временный файл и записываем в него некоторые данные
	tmpfile, err := ioutil.TempFile("", "example*.txt")
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

	mockRepository.EXPECT().UploadFile(tmpfile.Name()).Return(uint(1), nil)
	mockService.EXPECT().UploadFile(gomock.Any(), filepath.Base(tmpfile.Name())).Return(tmpfile.Name(), nil)

	server := Server{
		repo:    mockRepository,
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

	writer.Close()
	// Создаем запрос
	request := httptest.NewRequest(http.MethodPost, "/api/files/upload", &b)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	// Создаем запись ответа
	recorder := httptest.NewRecorder()

	server.UploadFile(recorder, request)

	body := recorder.Body

	var uploadFileAnswer uploadFileResp
	err = json.NewDecoder(body).Decode(&uploadFileAnswer)

	assert.Equal(t, uploadFileResp{Status: 200, Body: struct {
		PhotoID uint `json:"photoID"`
	}{1}}, uploadFileAnswer)

}
