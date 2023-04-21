package photo

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryPhoto interface {
	SavePhoto(row dataStruct.UserPhoto) error
	DeletePhoto(row dataStruct.UserPhoto) error

	GetAvatar(userId uint) (uint, error)
	GetPhotos(userId uint) ([]uint, error)
}
