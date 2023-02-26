package repository

import (
	"fmt"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type Repository struct {
	Users     *[]ds.User
	Cities    *[]ds.City
	UserToken *map[uint]string
}

func New() *Repository {
	var u []ds.User
	var c []ds.City
	t := make(map[uint]string)
	r := Repository{&u, &c, &t}
	return &r
}

func (r *Repository) AddUser(user *ds.User) error {
	userId := len(*r.Users)
	user.UserId = uint(userId)

	if err := r.CheckUsername(user.Username); err != nil {

		return err
	}

	if err := r.CheckEmail(user.Email); err != nil {

		return err
	}

	if err := CheckAge(user.Age); err != nil {
		return err
	}

	usersDB := *r.Users
	usersDB = append(usersDB, *user)
	r.Users = &usersDB

	return nil
}

func (r *Repository) SaveToken(userId uint, token string) {
	tokenUser := *r.UserToken
	tokenUser[userId] = token
	r.UserToken = &tokenUser
}

func (r *Repository) LoginEmail(email string, password string) (userId uint, err error) {
	var userpassword string

	for _, u := range *r.Users {
		if u.Email == email {
			userpassword = u.Password
			userId = u.UserId
			break
		}
	}

	if password == "" {

		err = fmt.Errorf("cant find user with such email")
		return
	}

	if userpassword == password {

		return
	}

	err = fmt.Errorf("password is not correct")
	return
}

func (r *Repository) LoginUsername(username string, password string) (userId uint, err error) {
	var userpassword string
	for _, u := range *r.Users {
		if u.Username == username {
			userpassword = u.Password
			userId = u.UserId
		}
	}

	if password == "" {
		err = fmt.Errorf("cant find user with such email")
		return
	}
	if userpassword == password {

		return userId, nil
	}

	err = fmt.Errorf("password is not correct")
	return
}

func (r *Repository) DeleteToken(token string) error {
	var userId uint
	flagFound := false
	for i, t := range *r.UserToken {
		if t == token {
			userId = i
			flagFound = true
			break
		}
	}

	if !flagFound {

		return fmt.Errorf("UnAuthorised")
	}

	delete(*r.UserToken, userId)

	return nil
}
