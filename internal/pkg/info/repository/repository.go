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

func (r *InfoRepository) GetZodiac() ([]string, error) {
	var zodiac []string
	err := r.db.Table("zodiacs").Select("zodiac").Find(&zodiac).Error
	if err != nil {
		return nil, err
	}
	return zodiac, nil
}

func (r *InfoRepository) GetJobs() ([]string, error) {
	var jobs []string
	err := r.db.Table("jobs").Select("job").Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *InfoRepository) GetEducation() ([]string, error) {
	var education []string
	err := r.db.Table("educations").Select("education").Find(&education).Error
	if err != nil {
		return nil, err
	}
	return education, nil
}

func (r *InfoRepository) GetCities() ([]string, error) {
	var cities []string
	err := r.db.Table("cities").Select("city").Find(&cities).Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *InfoRepository) GetReasons() ([]string, error) {
	var reasons []string
	err := r.db.Table("reasons").Select("reason").Find(&reasons).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *InfoRepository) GetStatuses() ([]string, error) {
	var statuses []string
	err := r.db.Table("statuses").Select("name").Find(&statuses).Error
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (r *InfoRepository) GetHashtags() ([]string, error) {
	var hashtags []string
	err := r.db.Table("hashtags").Select("hashtag").Find(&hashtags).Error
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

func (r *InfoRepository) GetStatusId(status string) (uint, error) {
	statusDB := &dataStruct.Status{}
	err := r.db.First(statusDB, "name = ?", status).Error
	return statusDB.Id, err
}
