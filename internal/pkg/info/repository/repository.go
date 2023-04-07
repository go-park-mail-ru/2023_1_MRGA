package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
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

func (r *InfoRepository) GetUserInfo(userId uint) (info.InfoStruct, error) {
	var infoStruct info.InfoStruct
	err := r.db.Table("user_infos").Select("*").
		Where("user_infos.user_id =?", userId).
		Joins("JOIN jobs on jobs.id = user_infos.job").
		Joins("Join educations on educations.id = user_infos.education").
		Joins("Join zodiacs on zodiacs.id = user_infos.zodiac").
		Joins("Join cities on cities.id = user_infos.city_id").
		Find(&infoStruct).Error

	return infoStruct, err
}

func (r *InfoRepository) GetAvatar(userId uint) (string, error) {
	var avatar dataStruct.UserPhoto
	err := r.db.Table("user_photos").Where("user_id = ? AND avatar=?", userId, true).Find(&avatar).Error
	return avatar.Photo, err
}

func (r *InfoRepository) GetPhotos(userId uint) ([]dataStruct.UserPhoto, error) {
	var photos []dataStruct.UserPhoto
	err := r.db.Table("user_photos").Where("user_id = ? AND avatar=?", userId, false).Find(&photos).Error
	return photos, err
}

///getId

func (r *InfoRepository) GetCityId(nameCity string) (uint, error) {
	city := &dataStruct.City{}
	err := r.db.First(city, "name = ?", nameCity).Error
	return city.Id, err
}

func (r *InfoRepository) GetZodiacId(nameZodiac string) (uint, error) {
	zodiac := &dataStruct.Zodiac{}
	err := r.db.First(zodiac, "name = ?", nameZodiac).Error
	return zodiac.Id, err
}

func (r *InfoRepository) GetJobId(nameJob string) (uint, error) {
	job := &dataStruct.Job{}
	err := r.db.First(job, "name = ?", nameJob).Error
	return job.Id, err
}

func (r *InfoRepository) GetEducationId(nameEducation string) (uint, error) {
	education := &dataStruct.Education{}
	err := r.db.First(education, "name = ?", nameEducation).Error
	return education.Id, err
}

////getters

func (r *InfoRepository) GetCities() ([]dataStruct.City, error) {
	var cities []dataStruct.City
	err := r.db.Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
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
