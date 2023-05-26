package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	ageCalc "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/age_calculator"
)

type RecRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *RecRepository {
	r := RecRepository{db}

	return &r
}

func (r *RecRepository) GetRecommendation(userId uint, history []uint, reasons []uint, hashtags []uint, filters dataStruct.UserFilter) (users []recommendation.UserRecommend, err error) {
	var sexSlice []uint
	switch filters.SearchSex {
	case 0:
		sexSlice = append(sexSlice, 0)
	case 1:
		sexSlice = append(sexSlice, 1)
	case 2:
		sexSlice = append(sexSlice, 0, 1)
	}
	err = r.db.Table("users u").Select("ui.user_id").
		Joins("JOIN user_infos ui on u.id = ui.user_id").
		//Joins("join user_hashtags uh on u.id = uh.user_id").
		//Joins("join user_reasons ur on u.id = ur.user_id").
		//Joins("LEFT Join complaints c on c.user_id = u.id").
		//Where("c.count is NULL or c.count < ?", 5).
		//Where("ui.user_id NOT IN ?", history).
		//Where("ui.sex IN ?", sexSlice).
		//Where("hashtag_id IN ?", hashtags).
		//Where("reason_id IN ?", reasons).
		Where("ui.user_id!=?", userId).
		Where("u.birth_day BETWEEN ? AND ?", ageCalc.CalculateBirthYear(filters.MaxAge), ageCalc.CalculateBirthYear(filters.MinAge)).
		Group("ui.user_id").
		//Order("COUNT(uh.hashtag_id) desc").
		Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *RecRepository) GetRecommendedUser(userId uint) (user recommendation.Recommendation, err error) {
	var filteredUser recommendation.DBRecommendation
	err = r.db.Table("users u").Select("u.id, ui.name, u.birth_day, ui.description, ui.sex, ed.education, z.zodiac, j.job, c.city").
		Where("u.id = ?", userId).
		Joins("Join user_infos ui on u.id = ui.user_id").
		Joins("Join educations ed on ed.id=ui.education").
		Joins("Join jobs j on j.id=ui.job").
		Joins("Join zodiacs z on z.id=ui.zodiac").
		Joins("Join cities c on c.id = ui.city_id").
		Find(&filteredUser).
		Error
	if err != nil {
		return user, err
	}
	age, err := ageCalc.CalculateAge(filteredUser.BirthDay)
	if err != nil {
		return user, err
	}

	user.Id = filteredUser.Id
	user.Name = filteredUser.Name
	user.Age = age
	user.Sex = filteredUser.Sex
	user.Description = filteredUser.Description
	user.City = filteredUser.City
	user.Zodiac = filteredUser.Zodiac
	user.Job = filteredUser.Job
	user.Education = filteredUser.Education

	return user, err
}

func (r *RecRepository) GetUserHistory(userId uint) ([]uint, error) {
	var users []uint
	err := r.db.Table("user_histories").Select("user_profile_id").Where("user_id = ? ", userId).Find(&users).Error
	return users, err
}
