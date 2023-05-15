package app

import (
	"mime/multipart"
	"os"
)

type IService interface {
	UploadFile(multipart.File, string, uint) (uint, error)
	GetFile(uint) (*os.File, string, error)
	GetFileByPath(string) (*os.File, string, error)
}
