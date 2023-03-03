package repository

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type Repository struct {
	Users     *[]ds.User
	Cities    *[]ds.City
	UserToken *map[uint]string
}

func NewRepo() *Repository {
	var userDS []ds.User
	var cityDS []ds.City
	tokenDS := make(map[uint]string)
	r := Repository{&userDS, &cityDS, &tokenDS}

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

func (r *Repository) Login(emailInp string, usernameInp string, passwordInp string) (userId uint, err error) {
	var userPassword string

	for _, user := range *r.Users {
		if user.Email == emailInp || user.Username == usernameInp {
			userPassword = user.Password
			userId = user.UserId
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

func (r *Repository) DeleteToken(token string) error {
	var userId uint
	flagFound := false
	for indexUser, tokenDS := range *r.UserToken {
		if tokenDS == token {
			userId = indexUser
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

func (r *Repository) GetCities() ([]string, error) {
	fileCity, err := os.Open("/Users/Staurran/GolandProjects/2023_1_MRGA3/files/city.txt")
	if err != nil {
		return nil, err
	}

	allCities, err := io.ReadAll(fileCity)
	if err != nil {
		return nil, err
	}

	allCitiesStr := string(allCities)
	cities := strings.Split(allCitiesStr, "\n")

	return cities, nil
}
