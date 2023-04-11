package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Service struct {
}

func InitService() Service {
	return Service{}
}

func (service Service) UploadFile(file multipart.File, filename string) (string, error) {
	dir := filepath.Join("services", "files_storage", "saved_files")

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(filename)
	baseName := filename[0 : len(filename)-len(ext)]
	newFileName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), ext)

	filePath := filepath.Join(dir, newFileName)

	newFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	if _, err := io.Copy(newFile, file); err != nil {
		return "", err
	}

	return filePath, nil
}

func (service Service) GetFile(filePath string) (*os.File, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}

	filename := path.Base(file.Name())

	return file, filename, nil
}
