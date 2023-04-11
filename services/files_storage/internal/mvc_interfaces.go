package internal

import (
	"mime/multipart"
	"net/http"
	"os"
)

type IRepository interface {
	UploadFile(string) (uint, error)
	GetFile(uint) (string, error)
}

type IServer interface {
	UploadFile(http.ResponseWriter, *http.Request)
	GetFile(http.ResponseWriter, *http.Request)
}

type IService interface {
	UploadFile(multipart.File, string) (string, error)
	GetFile(string) (*os.File, string, error)
}
