package repository

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/app"
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

func (r *Repository) Login(input string, passwordInp string) (userId uint, err error) {
	var userPassword string

	for _, user := range *r.Users {
		if user.Email == input || user.Username == input {
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

func (r *Repository) GetUserById(userId uint) (*app.UserRes, error) {

	for _, user := range *r.Users {
		if user.UserId == userId {
			userRes := app.UserRes{
				Username:    user.Username,
				Avatar:      user.Avatar,
				City:        user.City,
				Age:         user.Age,
				Sex:         user.Sex,
				Email:       user.Email,
				Description: user.Description,
			}
			return &userRes, nil
		}
	}

	return nil, fmt.Errorf("user are not found")
}

func (r *Repository) GetUserIdByToken(InpToken string) (uint, error) {
	for userId, userToken := range *r.UserToken {
		if userToken == InpToken {
			return userId, nil
		}
	}

	return 0, fmt.Errorf("user are not found")
}

func (r *Repository) GetRecommendation(userId uint) (recommendations []*app.Recommendation, err error) {
	count := 0

	for _, user := range *r.Users {
		if user.UserId != userId {
			recommendPerson := app.Recommendation{
				City:        user.City,
				Username:    user.Username,
				Age:         user.Age,
				Avatar:      user.Avatar,
				Description: user.Description,
				Sex:         user.Sex,
			}
			recommendations = append(recommendations, &recommendPerson)
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
