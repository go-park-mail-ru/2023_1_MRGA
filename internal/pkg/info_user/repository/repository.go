package repository

import (
	"time"

	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
)

type InfoRepository struct {
	db *gorm.DB
}

func NewInfoRepo(db *gorm.DB) *InfoRepository {
	return &InfoRepository{
		db,
	}
}

func (r *InfoRepository) AddInfoUser(userInfo *dataStruct.UserInfo) error {
	err := r.db.Create(userInfo).Error
	return err
}

func (r *InfoRepository) AddUserPhoto(userPhoto *dataStruct.UserPhoto) error {
	err := r.db.Create(userPhoto).Error
	return err
}

func (r *InfoRepository) GetUserInfo(userId uint) (info_user.InfoStruct, error) {
	var infoStruct info_user.InfoStruct
	err := r.db.Table("user_infos").Select("*").
		Where("user_infos.user_id =?", userId).
		Joins("JOIN users on users.id = user_infos.user_id").
		Joins("JOIN jobs on jobs.id = user_infos.job").
		Joins("Join educations on educations.id = user_infos.education").
		Joins("Join zodiacs on zodiacs.id = user_infos.zodiac").
		Joins("Join cities on cities.id = user_infos.city_id").
		Find(&infoStruct).Error

	return infoStruct, err
}

func (r *InfoRepository) GetUserPhoto(userId uint) ([]uint, error) {
	var photos []uint
	err := r.db.Table("user_photos").Select("photo").Where("user_id = ? and avatar=?", userId, false).Find(&photos).Error
	return photos, err
}

func (r *InfoRepository) ChangeInfo(userInfo *dataStruct.UserInfo) error {
	infoDB := &dataStruct.UserInfo{}
	err := r.db.First(infoDB, "user_id = ?", userInfo.UserId).Error
	if err != nil {
		return err
	}
	if userInfo.CityId != infoDB.CityId {
		infoDB.CityId = userInfo.CityId
	}
	if userInfo.Zodiac != infoDB.Zodiac {
		infoDB.Zodiac = userInfo.Zodiac
	}
	if userInfo.Job != infoDB.Job {
		infoDB.Job = userInfo.Job
	}
	if userInfo.Education != infoDB.Education {
		infoDB.Education = userInfo.Education
	}

	if userInfo.Description != "" {
		infoDB.Description = userInfo.Description
	}

	if userInfo.Name != "" {
		infoDB.Name = userInfo.Name
	}

	if userInfo.Sex != userInfo.Sex {
		infoDB.Sex = userInfo.Sex
	}

	err = r.db.Save(&infoDB).Error
	return err

}

func (r *InfoRepository) GetUserHashtags(userId uint) ([]uint, error) {
	var hashtags []uint
	err := r.db.Table("user_hashtags").Select("hashtag_id").Where("user_id = ? ", userId).Find(&hashtags).Error
	return hashtags, err
}

func (r *InfoRepository) AddUserHashtag(hashtag []dataStruct.UserHashtag) error {
	err := r.db.Create(&hashtag).Error
	return err
}

func (r *InfoRepository) DeleteUserHashtag(userId uint, hashtagId []uint) error {
	err := r.db.First(&dataStruct.UserHashtag{}, "user_id = ? AND hashtag_id IN?", userId, hashtagId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserHashtag{}, "user_id = ? AND hashtag_id IN ?", userId, hashtagId).Error
	return err
}

func (r *InfoRepository) GetHashtagById(hashtagId []uint) ([]string, error) {
	var hashtagDB []string
	err := r.db.Table("hashtags").Select("hashtag").
		Where("id IN ?", hashtagId).Find(&hashtagDB).Error
	return hashtagDB, err
}

func (r *InfoRepository) GetAvatar(userId uint) (uint, error) {
	var photo dataStruct.UserPhoto
	err := r.db.First(&photo, "user_id=? AND avatar = ?", userId, true).Error
	if err != nil {
		return 0, err
	}
	return photo.Photo, nil
}

func (r *InfoRepository) GetAge(userId uint) (int, error) {
	var birthday string
	err := r.db.Table("users").
		Select("birth_day").
		Where("id=?", userId).
		Find(&birthday).Error
	if err != nil {
		return 0, err
	}

	age, err := CalculateAge(birthday)
	if err != nil {
		return 0, err
	}

	return age, nil
}

func (r *InfoRepository) GetUserIdByEmail(email string) (uint, error) {
	var user dataStruct.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return 0, err
	}
	return user.Id, err
}

func CalculateAge(birthDay string) (int, error) {
	birth, err := time.Parse("2006-01-02", birthDay[:10])
	if err != nil {
		return 0, err
	}
	now := time.Now()
	age := now.Year() - birth.Year()
	if now.Month() > birth.Month() {
		age -= 1
	}
	if now.Month() == birth.Month() && now.Day() < birth.Day() {
		age -= 1
	}
	return age, nil
}
