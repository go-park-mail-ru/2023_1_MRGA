package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) *PhotoRepository {

	r := PhotoRepository{db}
	return &r
}

func (r *PhotoRepository) SavePhoto(row dataStruct.UserPhoto) error {
	err := r.CheckPhoto(row)
	if err != nil {
		return err
	}
	err = r.db.Create(&row).Error
	return err
}

func (r *PhotoRepository) DeletePhoto(row dataStruct.UserPhoto) error {
	err := r.db.First(&dataStruct.UserPhoto{}, "user_id =? AND photo =?", row.UserId, row.Photo).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserReaction{}, "user_id =? AND photo =?", row.UserId, row.Photo).Error
	return err
}
