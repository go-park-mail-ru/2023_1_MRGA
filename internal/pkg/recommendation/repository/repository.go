package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
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
		Joins("join user_hashtags uh on u.id = uh.user_id").
		Joins("join user_reasons ur on u.id = ur.user_id").
		Where("ui.user_id NOT IN ?", history).
		Where("hashtag_id IN ?", hashtags).
		Where("reason_id IN ?", reasons).
		Where("ui.user_id!=?", userId).
		Where("u.birth_day BETWEEN ? AND ?", calculateBirthYear(filters.MaxAge), calculateBirthYear(filters.MinAge)).
		Group("ui.user_id").
		Order("COUNT(uh.hashtag_id) desc").
		Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *RecRepository) GetRecommendedUser(userId uint) (user recommendation.Recommendation, err error) {
	var filteredUser recommendation.DBRecommendation
	err = r.db.Table("users u").Select("ui.name, u.birth_day, ui.description, ui.sex, ed.education, z.zodiac, j.job, c.city").
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
	age, err := calculateAge(filteredUser.BirthDay)
	if err != nil {
		return user, err
	}

	user.Name = filteredUser.Name
	user.Age = age
	user.Sex = filteredUser.Sex
	user.Description = filteredUser.Description
	user.City = filteredUser.City
	user.Zodiac = filteredUser.Zodiac
	user.Job = filteredUser.Job
	user.Education = filteredUser.Education

	var photos []dataStruct.UserPhoto
	err = r.db.Table("user_photos up").Where("user_id = ?", userId).Order("id DESC").Find(&photos).Error
	if err != nil {
		return user, err
	}

	var photosId []uint
	for _, photoItem := range photos {
		photosId = append(photosId, photoItem.Photo)
	}
	user.Photos = photosId

	var hashtags []string
	err = r.db.Table("user_hashtags uh").Select("h.hashtag").
		Where("uh.user_id = ?", userId).
		Joins("Join hashtags h on h.id = uh.hashtag_id").
		Order("h.hashtag DESC").
		Find(&hashtags).Error
	if err != nil {
		return user, err
	}
	user.Hashtags = hashtags

	return user, err
}

func (r *RecRepository) GetUserAge(userId uint) (int, error) {
	var user dataStruct.User
	err := r.db.First(&user, "id=?", userId).Error
	if err != nil {
		return 0, err
	}
	age, err := calculateAge(user.BirthDay)
	return age, err
}

func (r *RecRepository) GetUserHistory(userId uint) ([]dataStruct.UserHistory, error) {
	var users []dataStruct.UserHistory
	err := r.db.Table("user_histories").Where("user_id = ? ", userId).Find(&users).Error
	return users, err
}

func (r *RecRepository) GetUserHashtags(userId uint) ([]dataStruct.UserHashtag, error) {
	var hashtags []dataStruct.UserHashtag
	err := r.db.Table("user_hashtags").Where("user_id = ? ", userId).Find(&hashtags).Error
	return hashtags, err
}

func (r *RecRepository) GetReasonId(reason string) (uint, error) {
	reasonDB := &dataStruct.Reason{}
	err := r.db.First(reasonDB, "reason = ?", reason).Error
	return reasonDB.Id, err
}

func (r *RecRepository) AddUserReason(reason *dataStruct.UserReason) error {
	err := r.db.Create(&reason).Error
	return err
}

func (r *RecRepository) AddFilter(filter *dataStruct.UserFilter) error {
	err := r.db.Create(&filter).Error
	return err
}

func (r *RecRepository) GetReasons() ([]dataStruct.Reason, error) {
	var reasons []dataStruct.Reason
	err := r.db.Find(&reasons).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *RecRepository) GetFilter(userId uint) (dataStruct.UserFilter, error) {
	filterDB := dataStruct.UserFilter{}
	err := r.db.First(&filterDB, "user_id = ?", userId).Error
	return filterDB, err
}

func (r *RecRepository) ChangeFilter(newFilter dataStruct.UserFilter) error {
	filterDB := &dataStruct.UserFilter{}
	err := r.db.First(filterDB, "user_id = ?", newFilter.UserId).Error
	if err != nil {
		return err
	}
	if filterDB.MaxAge != newFilter.MaxAge {
		filterDB.MaxAge = newFilter.MaxAge
	}
	if filterDB.MinAge != newFilter.MinAge {
		filterDB.MinAge = newFilter.MinAge
	}
	if filterDB.SearchSex != newFilter.SearchSex {
		filterDB.SearchSex = newFilter.SearchSex
	}

	err = r.db.Save(&filterDB).Error
	return err

}

func (r *RecRepository) GetUserReasons(userId uint) ([]dataStruct.UserReason, error) {
	var reasons []dataStruct.UserReason
	err := r.db.Find(&reasons, "user_id =?", userId).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *RecRepository) GetReasonById(reasonId uint) (string, error) {
	reasonDB := dataStruct.Reason{}
	err := r.db.First(&reasonDB, "id = ?", reasonId).Error
	return reasonDB.Reason, err
}

func (r *RecRepository) DeleteUserReason(userId, reactionId uint) error {
	err := r.db.First(&dataStruct.UserReason{}, "user_id = ? AND reason_id=?", userId, reactionId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserReason{}, "user_id = ? AND reason_id=?", userId, reactionId).Error
	return err
}

func calculateBirthYear(age int) string {
	return fmt.Sprintf("%d-01-01", time.Now().Year()-age)
}

func calculateAge(birthDay string) (int, error) {
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
