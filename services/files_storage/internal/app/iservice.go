package app

import (
	"mime/multipart"
	"os"
)

type IService interface {
	UploadFile(multipart.File, string, uint) (string, error)
	GetFile(string) (*os.File, string, error)
}
