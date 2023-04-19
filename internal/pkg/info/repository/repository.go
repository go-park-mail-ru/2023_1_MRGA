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

func (r *InfoRepository) GetCities() ([]dataStruct.City, error) {
	var cities []dataStruct.City
	err := r.db.Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *InfoRepository) GetHashtags() ([]dataStruct.Hashtag, error) {
	var hashtags []dataStruct.Hashtag
	err := r.db.Find(&hashtags).Error
	if err != nil {
		return nil, err
	}
	return hashtags, nil
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

func (r *InfoRepository) GetHashtagId(nameHashtag []string) ([]uint, error) {
	var hashtagId []uint

	err := r.db.Table("hashtags").Select("id").
		Where("hashtag IN ?", nameHashtag).
		Find(&hashtagId).Error
	return hashtagId, err
}

func (r *InfoRepository) GetEducationId(nameEducation string) (uint, error) {
	education := &dataStruct.Education{}
	err := r.db.First(education, "education = ?", nameEducation).Error
	return education.Id, err
}
