package usecase

import (
	"fmt"

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

func (p *PhotoUseCase) DeletePhoto(userId uint, num int) error {
	photos, err := p.GetAllPhotos(userId)
	if err != nil {
		return err
	}
	if len(photos) <= num {
		err = fmt.Errorf("wrong number of photo")
		return err
	}

	var rowPhoto dataStruct.UserPhoto
	rowPhoto.Photo = photos[num]
	rowPhoto.UserId = userId

	err = p.userRepo.DeletePhoto(rowPhoto)
	return err
}

func (p *PhotoUseCase) GetAllPhotos(userId uint) ([]uint, error) {
	avatar, err := p.userRepo.GetAvatar(userId)
	if err != nil {
		return nil, err
	}
	var result []uint
	result = append(result, avatar)
	photos, err := p.userRepo.GetPhotos(userId)
	if err != nil {
		return nil, err
	}
	result = append(result, photos...)
	return result, nil
}

func (p *PhotoUseCase) ChangePhoto(num int, photoId uint, userId uint) error {
	photos, err := p.GetAllPhotos(userId)
	if err != nil {
		return err
	}
	if len(photos) <= num {
		err = p.SavePhoto(userId, photoId, false)
		if err != nil {
			return err
		}
		return nil
	}
	oldPhoto := photos[num]
	err = p.userRepo.ChangePhoto(oldPhoto, userId, photoId)
	if err != nil {
		return err
	}
	return nil
}

func (p *PhotoUseCase) GetAvatar(userId uint) (uint, error) {
	avatar, err := p.userRepo.GetAvatar(userId)
	if err != nil {
		return 0, err
	}
	return avatar, nil
}
