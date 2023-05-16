package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	ageCalc "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/age_calculator"
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
	infoDB.CityId = userInfo.CityId
	infoDB.Zodiac = userInfo.Zodiac
	infoDB.Job = userInfo.Job
	infoDB.Education = userInfo.Education
	infoDB.Description = userInfo.Description
	infoDB.Name = userInfo.Name
	infoDB.Sex = userInfo.Sex

	err = r.db.Save(&infoDB).Error
	return err

}

func (r *InfoRepository) GetUserHashtags(userId uint) ([]string, error) {
	var hashtags []string
	err := r.db.Table("user_hashtags uh").
		Select("h.hashtag").
		Joins("Join hashtags h on h.id= uh.hashtag_id").
		Where("uh.user_id = ? ", userId).Find(&hashtags).Error
	return hashtags, err
}

func (r *InfoRepository) GetUserHashtagsId(userId uint) ([]uint, error) {
	var hashtags []uint
	err := r.db.Table("user_hashtags ").
		Select("hashtag_id").
		Where("user_id = ? ", userId).Find(&hashtags).Error
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

func (r *InfoRepository) GetUserById(userId uint) (userRes info_user.UserRestTemp, err error) {

	user := info_user.UserRestTemp{}
	err = r.db.Table("users").Select("user_infos.name").
		Where("users.id =?", userId).
		Joins("Join user_infos on users.id = user_infos.user_id").
		Find(&user).Error
	if err != nil {
		return
	}

	return user, nil
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

	age, err := ageCalc.CalculateAge(birthday)
	if err != nil {
		return 0, err
	}

	return age, nil
}

func (r *InfoRepository) CheckFilter(userId uint) (bool, error) {
	var reasons []uint
	err := r.db.Table("user_reasons").Select("reason_id").Where("user_id = ?", userId).Find(&reasons).Error
	if err != nil {
		return false, err
	}
	if len(reasons) == 0 {
		return false, nil
	}
	return true, nil
}

func (r *InfoRepository) GetUserStatus(userId uint) (string, error) {
	var status string
	err := r.db.Table("users u").Select("s.name").
		Where("u.id=?", userId).
		Joins("Join statuses s on u.status=s.id").
		Find(&status).Error
	return status, err
}

func (r *InfoRepository) ChangeUserStatus(userId, statusId uint) error {
	var user dataStruct.User
	err := r.db.First(&user, "id=?", userId).Error
	if err != nil {
		return err
	}
	user.Status = statusId
	err = r.db.Save(&user).Error
	return err
}
