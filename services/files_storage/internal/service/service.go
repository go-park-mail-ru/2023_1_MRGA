package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(file multipart.File, handler *multipart.FileHeader, userID uint) (string, error) {
	dir := filepath.Join("services", "files_storage", "saved_files", fmt.Sprintf("%d", userID))

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(handler.Filename)
	baseName := handler.Filename[0 : len(handler.Filename)-len(ext)]
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

func GetFile(filePath string) {

}
