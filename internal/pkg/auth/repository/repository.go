package repository

import (
	"fmt"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type AuthRepository struct {
	Users      []dataStruct.User
	UserTokens map[uint]string
}

func NewRepo() *AuthRepository {
	var userDS []dataStruct.User
	tokenDS := make(map[uint]string)
	r := AuthRepository{userDS, tokenDS}

	return &r
}

func (r *AuthRepository) Login(input string, passwordInp string) (userId uint, err error) {
	var userPassword string

	for _, userdb := range r.Users {
		if userdb.Email == input || userdb.Username == input {
			userPassword = userdb.Password
			userId = userdb.UserId
			break
		}
	}
	switch userPassword {
	case "":
		err = fmt.Errorf("cant find user with such email")
		return
	case passwordInp:
		return
	}

	err = fmt.Errorf("password is not correct")
	return
}

func (r *AuthRepository) AddUser(user dataStruct.User) (uint, error) {
	userId := len(r.Users)
	user.UserId = uint(userId)

	if err := r.CheckUsername(user.Username); err != nil {
		return 0, err
	}

	if err := r.CheckEmail(user.Email); err != nil {
		return 0, err
	}

	if err := CheckAge(user.Age); err != nil {
		return 0, err
	}

	usersDB := r.Users
	usersDB = append(usersDB, user)
	r.Users = usersDB

	return user.UserId, nil
}

func (r *AuthRepository) DeleteToken(token string) error {
	var userId uint
	flagFound := false
	for indexUser, tokenDS := range r.UserTokens {
		if tokenDS == token {
			userId = indexUser
			flagFound = true
			break
		}
	}

	if !flagFound {
		return fmt.Errorf("UnAuthorised")
	}

	delete(r.UserTokens, userId)
	return nil
}

func (r *AuthRepository) SaveToken(userId uint, token string) {
	tokenUser := r.UserTokens
	tokenUser[userId] = token
	r.UserTokens = tokenUser
}
