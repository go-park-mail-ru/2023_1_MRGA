package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
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
