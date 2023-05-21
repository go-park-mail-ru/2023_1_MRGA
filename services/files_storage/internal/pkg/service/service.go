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

func (service Service) UploadPhoto(file multipart.File, filename string, userID uint) (fileID uint, err error) {
	dir := filepath.Join("services", "files_storage", "saved_files", fmt.Sprintf("%d", userID))

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	ext := filepath.Ext(filename)
	baseName := filename[0 : len(filename)-len(ext)]
	newFileName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), ext)

	filePath := filepath.Join(dir, newFileName)

	newFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer newFile.Close()

	if _, err = file.Seek(0, 0); err != nil {
		return
	}

	if _, err = io.Copy(newFile, file); err != nil {
		return
	}

	fileID, err = service.repository.UploadPhoto(filePath, userID)
	if err != nil {
		return
	}

	return
}

func (service Service) UploadFile(file multipart.File, filename string, userID uint) (pathToFile string, err error) {
	dir := filepath.Join("services", "files_storage", "saved_files", fmt.Sprintf("%d", userID))

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	ext := filepath.Ext(filename)
	baseName := filename[0 : len(filename)-len(ext)]
	newFileName := fmt.Sprintf("%s_%d%s", baseName, time.Now().UnixNano(), ext)

	pathToFile = filepath.Join(dir, newFileName)

	newFile, err := os.Create(pathToFile)
	if err != nil {
		return
	}
	defer newFile.Close()

	if _, err = file.Seek(0, 0); err != nil {
		return
	}

	if _, err = io.Copy(newFile, file); err != nil {
		return
	}

	err = service.repository.UploadFile(pathToFile, userID)
	if err != nil {
		return
	}

	return
}

func (service Service) GetFile(userID uint) (file *os.File, filename string, err error) {
	var filePath string
	if filePath, err = service.repository.GetFile(userID); err != nil {
		return
	}

	file, filename, err = openFileByPath(filePath)
	if err != nil {
		return
	}

	return
}

func (service Service) GetFileByPath(filePath string) (file *os.File, filename string, err error) {
	file, filename, err = openFileByPath(filePath)
	if err != nil {
		return
	}

	return
}

func openFileByPath(pathToFile string) (file *os.File, filename string, err error) {
	file, err = os.Open(pathToFile)
	if err != nil {
		return
	}

	filename = path.Base(file.Name())

	return
}
