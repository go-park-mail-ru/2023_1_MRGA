package app

import "net/http"

type IServer interface {
	UploadFileV1(http.ResponseWriter, *http.Request)
	UploadFile(http.ResponseWriter, *http.Request)
	GetFile(http.ResponseWriter, *http.Request)
	GetFileByPath(http.ResponseWriter, *http.Request)
}
