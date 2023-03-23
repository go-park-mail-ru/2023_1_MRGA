package repository

import (
	"fmt"
	"io"
	"os"
	"strings"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user"
)

type Repository struct {
	Users      []dataStruct.User
	Cities     []dataStruct.City
	UserTokens map[uint]string
}

func NewRepo() *Repository {
	var userDS []dataStruct.User
	var cityDS []dataStruct.City
	tokenDS := make(map[uint]string)
	r := Repository{userDS, cityDS, tokenDS}

	return &r
}

func (r *Repository) AddUser(user dataStruct.User) (uint, error) {
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

func (r *Repository) SaveToken(userId uint, token string) {
	tokenUser := r.UserTokens
	tokenUser[userId] = token
	r.UserTokens = tokenUser
}

func (r *Repository) Login(input string, passwordInp string) (userId uint, err error) {
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

func (r *Repository) DeleteToken(token string) error {
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

func (r *Repository) GetCities() ([]string, error) {
	fileCity, err := os.Open("./files/city.txt")
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

func (r *Repository) GetUserById(userId uint) (userRes user.UserRes, err error) {

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

func (r *Repository) GetUserIdByToken(InpToken string) (uint, error) {
	for userId, userToken := range r.UserTokens {
		if userToken == InpToken {
			return userId, nil
		}
	}

	return 0, fmt.Errorf("user are not found")
}

func (r *Repository) GetRecommendation(userId uint) (recommendations []recommendation.Recommendation, err error) {
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
