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

type IHandler interface {
	uploadFile(http.ResponseWriter, *http.Request)
	getFile(http.ResponseWriter, *http.Request)
}

type IService interface {
	UploadFile(multipart.File, *multipart.FileHeader) (string, error)
	GetFile(string) (*os.File, string, error)
}
