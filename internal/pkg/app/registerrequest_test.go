package app

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type psevdoRepo struct {
	Users     *[]ds.User
	Cities    *[]ds.City
	UserToken *map[uint]string
}

func NewRepo() *psevdoRepo {
	var userDS []ds.User
	var cityDS []ds.City
	tokenDS := make(map[uint]string)
	r := psevdoRepo{&userDS, &cityDS, &tokenDS}

	return &r
}

func (pr *psevdoRepo) AddUser(user *ds.User) error {
	return nil
}

func (pr *psevdoRepo) DeleteToken(token string) error {
	return nil
}

func (pr *psevdoRepo) Login(emailInp string, usernameInp string, passwordInp string) (userId uint, err error) {
	return 0, nil
}

func (pr *psevdoRepo) SaveToken(userId uint, token string) {

}

func (pr *psevdoRepo) GetCities() ([]string, error) {
	return nil, nil
}

//func TestHealthCheckHandler(t *testing.T) {
//	reader := strings.NewReader("username=user1")
//	data := map[string]string{
//		"username": "user1",
//		"email":    "email1",
//		"password": "1234",
//	}
//	req, err := http.NewRequest("POST", "/login", data)
//	if err != nil {
//		t.Fatal(err)
//	}
//}

type cityType struct {
	City []string `json:"city"`
}

//
//func TestApplication_GetCities(t *testing.T) {
//	repo := NewRepo()
//	applicationTest := New(repo)
//	req, err := http.NewRequest(http.MethodGet, "/cities", nil)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(applicationTest.GetCities)
//
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	expected := map[string][]string{
//		"city": nil,
//	}
//	jsonExpected, _ := json.Marshal(expected)
//	var city cityType
//	_ = json.Unmarshal(rr.Body.Bytes(), &city)
//	if city != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//}
