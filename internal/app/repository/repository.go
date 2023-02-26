package repository

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type Repository struct {
	Users  *[]ds.User
	Cities *[]ds.City
}

func New() *Repository {
	var u []ds.User
	var c []ds.City
	r := Repository{&u, &c}
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
