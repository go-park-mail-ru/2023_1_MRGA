package app

import (
	"mime/multipart"
	"os"
)

type IService interface {
	UploadPhoto(multipart.File, string, uint) (uint, error)
	UploadFile(multipart.File, string, uint) (string, error)
	GetFile(uint) (*os.File, string, error)
	GetFileByPath(string) (*os.File, string, error)
}
