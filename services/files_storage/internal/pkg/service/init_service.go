package service

import "github.com/go-park-mail-ru/2023_1_MRGA.git/services/files_storage/internal/app"

type Service struct {
	repository app.IRepository
}

func InitService(repository app.IRepository) (service Service, err error) {
	service = Service{
		repository: repository,
	}
	return
}
