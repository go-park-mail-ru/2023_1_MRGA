package repository

import (
	"gorm.io/gorm"
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
