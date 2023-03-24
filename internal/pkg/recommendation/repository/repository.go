package repository

import (
	"fmt"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type RecRepository struct {
	Users      []dataStruct.User
	UserTokens map[uint]string
}

func NewRepo() *RecRepository {
	var userDS []dataStruct.User
	tokenDS := make(map[uint]string)
	r := RecRepository{userDS, tokenDS}

	return &r
}

func (r *RecRepository) GetRecommendation(userId uint) (recommendations []recommendation.Recommendation, err error) {
	count := 0

	for _, userdb := range r.Users {
		if userdb.UserId != userId {
			recommendPerson := recommendation.Recommendation{
				City:        userdb.City,
				Username:    userdb.Username,
				Age:         userdb.Age,
				Avatar:      userdb.Avatar,
				Description: userdb.Description,
				Sex:         userdb.Sex,
			}
			recommendations = append(recommendations, recommendPerson)
			count += 1
			if count == 10 {
				break
			}
		}
	}

	if count == 0 {
		return nil, fmt.Errorf("no users yet")
	}

	return recommendations, nil
}

func (r *RecRepository) GetUserIdByToken(InpToken string) (uint, error) {
	for userId, userToken := range r.UserTokens {
		if userToken == InpToken {
			return userId, nil
		}
	}

	return 0, fmt.Errorf("user are not found")
}
