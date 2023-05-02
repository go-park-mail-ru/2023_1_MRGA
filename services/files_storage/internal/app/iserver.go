package app

import "net/http"

type IServer interface {
	UploadFile(http.ResponseWriter, *http.Request)
	GetFile(http.ResponseWriter, *http.Request)
}
