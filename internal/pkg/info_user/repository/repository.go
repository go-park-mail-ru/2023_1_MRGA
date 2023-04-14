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
		Joins("JOIN jobs on jobs.id = user_infos.job").
		Joins("Join educations on educations.id = user_infos.education").
		Joins("Join zodiacs on zodiacs.id = user_infos.zodiac").
		Joins("Join cities on cities.id = user_infos.city_id").
		Find(&infoStruct).Error

	return infoStruct, err
}

func (r *InfoRepository) GetUserPhoto(userId uint) (photos []dataStruct.UserPhoto, err error) {
	err = r.db.Find(&photos, "user_id = ?", userId).Error
	return
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

func (r *InfoRepository) GetHashtags() ([]dataStruct.Hashtag, error) {
	var hashtags []dataStruct.Hashtag
	err := r.db.Find(&hashtags).Error
	if err != nil {
		return nil, err
	}
	return hashtags, nil
}

func (r *InfoRepository) GetUserHashtags(userId uint) ([]dataStruct.UserHashtag, error) {
	var hashtags []dataStruct.UserHashtag
	err := r.db.Table("user_hashtags").Where("user_id = ? ", userId).Find(&hashtags).Error
	return hashtags, err
}

func (r *InfoRepository) AddUserHashtag(hashtag dataStruct.UserHashtag) error {
	err := r.db.Create(&hashtag).Error
	return err
}

func (r *InfoRepository) DeleteUserHashtag(userId, hashtagId uint) error {
	err := r.db.First(&dataStruct.UserHashtag{}, "user_id = ? AND hashtag_id=?", userId, hashtagId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserHashtag{}, "user_id = ? AND hashtag_id=?", userId, hashtagId).Error
	return err
}

func (r *InfoRepository) GetHashtagId(nameHashtag string) (uint, error) {
	hashtag := &dataStruct.Hashtag{}
	err := r.db.First(hashtag, "hashtag = ?", nameHashtag).Error
	return hashtag.Id, err
}

func (r *InfoRepository) GetHashtagById(hashtagId uint) (string, error) {
	hashtagDB := dataStruct.Hashtag{}
	err := r.db.First(&hashtagDB, "id = ?", hashtagId).Error
	return hashtagDB.Hashtag, err
}

func (r *InfoRepository) GetCityId(nameCity string) (uint, error) {
	city := &dataStruct.City{}
	err := r.db.First(city, "city = ?", nameCity).Error
	return city.Id, err
}

func (r *InfoRepository) GetZodiacId(nameZodiac string) (uint, error) {
	zodiac := &dataStruct.Zodiac{}
	err := r.db.First(zodiac, "zodiac = ?", nameZodiac).Error
	return zodiac.Id, err
}

func (r *InfoRepository) GetJobId(nameJob string) (uint, error) {
	job := &dataStruct.Job{}
	err := r.db.First(job, "job = ?", nameJob).Error
	return job.Id, err
}

func (r *InfoRepository) GetEducationId(nameEducation string) (uint, error) {
	education := &dataStruct.Education{}
	err := r.db.First(education, "education = ?", nameEducation).Error
	return education.Id, err
}

func (r *InfoRepository) GetCities() ([]dataStruct.City, error) {
	var cities []dataStruct.City
	err := r.db.Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *InfoRepository) GetAge(userId uint) (int, error) {
	var user dataStruct.User
	err := r.db.First(&user, "id=?", userId).Error
	if err != nil {
		return 0, err
	}
	age, err := CalculateAge(user.BirthDay)
	if err != nil {
		return 0, err
	}
	return age, nil
}

func (r *InfoRepository) GetZodiac() ([]dataStruct.Zodiac, error) {
	var zodiac []dataStruct.Zodiac
	err := r.db.Find(&zodiac).Error
	if err != nil {
		return nil, err
	}
	return zodiac, nil
}

func (r *InfoRepository) GetJobs() ([]dataStruct.Job, error) {
	var jobs []dataStruct.Job
	err := r.db.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *InfoRepository) GetEducation() ([]dataStruct.Education, error) {
	var education []dataStruct.Education
	err := r.db.Find(&education).Error
	if err != nil {
		return nil, err
	}
	return education, nil
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
