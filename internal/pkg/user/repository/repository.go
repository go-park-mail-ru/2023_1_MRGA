package repository

import (
	"fmt"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user"
)

type UserRepository struct {
	Users      []dataStruct.User
	UserTokens map[uint]string
}

func NewRepo() *UserRepository {
	var userDS []dataStruct.User
	tokenDS := make(map[uint]string)
	r := UserRepository{userDS, tokenDS}

	return &r
}

func (r *UserRepository) GetUserIdByToken(InpToken string) (uint, error) {
	for userId, userToken := range r.UserTokens {
		if userToken == InpToken {
			return userId, nil
		}
	}

	return 0, fmt.Errorf("user are not found")
}

func (r *UserRepository) GetUserById(userId uint) (userRes user.UserRes, err error) {

	for _, userdb := range r.Users {
		if userdb.UserId == userId {
			userRes = user.UserRes{
				Username:    userdb.Username,
				Avatar:      userdb.Avatar,
				City:        userdb.City,
				Age:         userdb.Age,
				Sex:         userdb.Sex,
				Email:       userdb.Email,
				Description: userdb.Description,
			}
			return userRes, nil
		}
	}

	return userRes, fmt.Errorf("user are not found")
}
