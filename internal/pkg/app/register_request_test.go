package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type pseudoRepo struct {
	Users      *[]dataStruct.User
	Cities     *[]dataStruct.City
	UserTokens *map[uint]string
}

func Checkout(ts *httptest.Server, inputJson string) (result map[string]interface{}, err error) {
	var resp *http.Response
	resp, err = http.Post(ts.URL, "application/json", bytes.NewBuffer([]byte(inputJson)))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var jsonStr []byte
	jsonStr, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func convertToString(val interface{}) (strVal string) {
	switch val.(type) {
	case string:
		strVal = val.(string)
	case float64:
		strVal = fmt.Sprint(val.(float64))
	case int:
		strVal = fmt.Sprint(val.(int))
	}
	return strVal
}

func mapEqual(got, expected map[string]interface{}) bool {
	for keyGot, valueGot := range got {
		valueExpected := expected[keyGot]

		var (
			strValueExpected string = convertToString(valueExpected)
			strValueGot      string = convertToString(valueGot)
		)
		if strValueExpected != strValueGot {
			return false
		}
	}
	return true
}

func TestRegister(t *testing.T) {
	tests := map[string]struct {
		inputJson  string
		outputJson map[string]interface{}
	}{
		"Обычный тест на создание пользователя": {
			inputJson:  `{"username": "masharpik", "email": "masharpik@gmail.com", "password": "masharpik2004", "age": 19}`,
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		"Создание пользователя с тем же ником": {
			inputJson:  `{"username": "masharpik", "email": "masharpik2@gmail.com", "password": "masharpik2004", "age": 19}`,
			outputJson: map[string]interface{}{"err": "username is not unique", "status": http.StatusBadRequest},
		},
	}

	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.Register))

	for testName, test := range tests {
		testName := testName
		test := test
		t.Run(testName, func(t *testing.T) {
			result, err := Checkout(ts, test.inputJson)
			if err != nil {
				t.Errorf("[%s] unexpected error: %#v", testName, err)
			}
			if !mapEqual(result, test.outputJson) {
				t.Errorf("[%s] wrong result, expected %#v, got %#v", testName, test.outputJson, result)
			}
		})
	}
}

// НИЖЕ РЕАЛИЗАЦИЯ ПСЕВДОРЕПЫ

func (r *pseudoRepo) AddUser(user *dataStruct.User) error {
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

func (r *pseudoRepo) DeleteToken(token string) error {
	var userId uint
	flagFound := false
	for indexUser, tokenDS := range *r.UserTokens {
		if tokenDS == token {
			userId = indexUser
			flagFound = true
			break
		}
	}

	if !flagFound {
		return fmt.Errorf("UnAuthorised")
	}

	delete(*r.UserTokens, userId)
	return nil
}

func (r *pseudoRepo) Login(input string, passwordInp string) (userId uint, err error) {
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

func (r *pseudoRepo) SaveToken(userId uint, token string) {
	tokenUser := *r.UserTokens
	tokenUser[userId] = token
	r.UserTokens = &tokenUser
}

func (pr *pseudoRepo) GetCities() ([]string, error) {
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

func (r *pseudoRepo) GetUserById(userId uint) (*UserRes, error) {

	for _, user := range *r.Users {
		if user.UserId == userId {
			userRes := UserRes{
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

func (r *pseudoRepo) GetUserIdByToken(InpToken string) (uint, error) {
	for userId, userToken := range *r.UserTokens {
		if userToken == InpToken {
			return userId, nil
		}
	}

	return 0, fmt.Errorf("user are not found")
}

func (r *pseudoRepo) GetRecommendation(userId uint) (recommendations []*Recommendation, err error) {
	count := 0

	for _, user := range *r.Users {
		if user.UserId != userId {
			recommendPerson := Recommendation{
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

func NewRepo() *pseudoRepo {
	var userDS []dataStruct.User
	var cityDS []dataStruct.City
	tokenDS := make(map[uint]string)
	r := pseudoRepo{&userDS, &cityDS, &tokenDS}

	return &r
}

func (r *pseudoRepo) CheckUsername(username string) error {
	for _, user := range *r.Users {
		if username == user.Username {

			return fmt.Errorf("username is not unique")
		}
	}

	return nil
}

func (r *pseudoRepo) CheckEmail(email string) error {
	for _, user := range *r.Users {
		if email == user.Email {

			return fmt.Errorf("email is not unique")
		}
	}

	return nil
}

func CheckAge(age int) error {
	if age > 150 || age < 18 {

		return fmt.Errorf("age is not correct")
	}
	return nil
}
