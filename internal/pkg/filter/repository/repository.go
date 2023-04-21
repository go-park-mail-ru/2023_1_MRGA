package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type FilterRepository struct {
	db *gorm.DB
}

func NewFilterRepo(db *gorm.DB) *FilterRepository {
	return &FilterRepository{
		db,
	}
}

func (r *FilterRepository) AddFilter(filter *dataStruct.UserFilter) error {
	err := r.db.Create(&filter).Error
	return err
}

func (r *FilterRepository) GetFilter(userId uint) (dataStruct.UserFilter, error) {
	filterDB := dataStruct.UserFilter{}
	err := r.db.First(&filterDB, "user_id = ?", userId).Error
	return filterDB, err
}

func (r *FilterRepository) ChangeFilter(newFilter dataStruct.UserFilter) error {
	filterDB := &dataStruct.UserFilter{}
	err := r.db.First(filterDB, "user_id = ?", newFilter.UserId).Error
	if err != nil {
		return err
	}
	filterDB.MaxAge = newFilter.MaxAge
	filterDB.MinAge = newFilter.MinAge
	filterDB.SearchSex = newFilter.SearchSex
	err = r.db.Save(&filterDB).Error
	return err

}

func (r *FilterRepository) GetUserReasonsId(userId uint) ([]uint, error) {
	var reasons []uint
	err := r.db.Table("user_reasons").Select("reason_id").Find(&reasons, "user_id =?", userId).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *FilterRepository) GetUserReasons(userId uint) ([]string, error) {
	var reasons []string
	err := r.db.Table("user_reasons ur").Select("r.reason").
		Joins("Join reasons r on ur.reason_id = r.id").Find(&reasons, "user_id =?", userId).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *FilterRepository) GetReasonById(reasonId []uint) ([]string, error) {
	var reasonDB []string
	err := r.db.Table("reasons").Select("reason").Find(&reasonDB, "id IN ?", reasonId).Error
	return reasonDB, err
}

func (r *FilterRepository) GetReasonId(reason []string) ([]uint, error) {
	var reasonDB []uint
	err := r.db.Table("reasons").Select("id").Find(&reasonDB, "reason IN ?", reason).Error
	return reasonDB, err
}

func (r *FilterRepository) AddUserReason(reason []dataStruct.UserReason) error {
	err := r.db.Create(&reason).Error
	return err
}

func (r *FilterRepository) DeleteUserReason(userId uint, reactionId []uint) error {
	err := r.db.Delete(&dataStruct.UserReason{}, "user_id = ? AND reason_id= IN ?", userId, reactionId).Error
	return err
}
