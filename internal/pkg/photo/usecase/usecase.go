package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
)

type PhotoUseCase struct {
	userRepo photo.IRepositoryPhoto
}

func NewPhotoUseCase(
	userRepo photo.IRepositoryPhoto,
) *PhotoUseCase {
	return &PhotoUseCase{
		userRepo: userRepo,
	}
}

func (p *PhotoUseCase) SavePhoto(userId uint, photoId uint, avatar bool) error {
	var rowPhoto dataStruct.UserPhoto
	rowPhoto.Photo = photoId
	rowPhoto.Avatar = avatar
	rowPhoto.UserId = userId

	err := p.userRepo.SavePhoto(rowPhoto)
	return err
}

func (p *PhotoUseCase) DeletePhoto(userId uint, photoId uint) error {
	var rowPhoto dataStruct.UserPhoto
	rowPhoto.Photo = photoId
	rowPhoto.UserId = userId

	err := p.userRepo.DeletePhoto(rowPhoto)
	return err
}
