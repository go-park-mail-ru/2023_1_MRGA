package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type RecRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *RecRepository {
	r := RecRepository{db}

	return &r
}

//
//func (r *RecRepository) GetRecommendation(userId uint) (recommendations []recommendation.Recommendation, err error) {
//	count := 0
//
//	for _, userdb := range r.Users {
//		if userdb.UserId != userId {
//			recommendPerson := recommendation.Recommendation{
//				City:        userdb.City,
//				Username:    userdb.Username,
//				Age:         userdb.Age,
//				Avatar:      userdb.Avatar,
//				Description: userdb.Description,
//				Sex:         userdb.Sex,
//			}
//			recommendations = append(recommendations, recommendPerson)
//			count += 1
//			if count == 10 {
//				break
//			}
//		}
//	}
//
//	if count == 0 {
//		return nil, fmt.Errorf("no users yet")
//	}
//
//	return recommendations, nil
//}
//
//func (r *RecRepository) GetUserIdByToken(InpToken string) (uint, error) {
//	for userId, userToken := range r.UserTokens {
//		if userToken == InpToken {
//			return userId, nil
//		}
//	}
//
//	return 0, fmt.Errorf("user are not found")
//}

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
	err := r.db.First(&dataStruct.UserReason{}, "user_id = ? AND reason_is=?", userId, reactionId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserReason{}, "user_id = ? AND reason_is=?", userId, reactionId).Error
	return err
}
