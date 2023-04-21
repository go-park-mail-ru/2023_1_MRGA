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

	err := r.db.First(&row, "user_id =? AND photo =?", row.UserId, row.Photo).Error
	if err != nil {
		return err
	}
	err = r.CheckDeletedPhoto(row)
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserPhoto{}, "user_id =? AND photo =?", row.UserId, row.Photo).Error
	return err
}

func (r *PhotoRepository) GetAvatar(userId uint) (uint, error) {
	var photoId uint
	err := r.db.Table("user_photos p").Select("photo").
		Where("user_id = ? AND avatar = ?", userId, true).Find(&photoId).Error
	return photoId, err
}

func (r *PhotoRepository) GetPhotos(userId uint) ([]uint, error) {
	var photos []uint
	err := r.db.Table("user_photos up").Select("up.photo").
		Where("user_id = ? AND avatar = ?", userId, false).
		Order("up.id ASC").Find(&photos).Error
	return photos, err
}
