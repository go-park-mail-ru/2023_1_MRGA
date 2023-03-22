package app

import (
	"net/http"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type IRepository interface {
	DeleteToken(string) error
	AddUser(d dataStruct.User) (uint, error)
	Login(input string, pass string) (uint, error)
	SaveToken(id uint, token string)
	GetCities() ([]string, error)
	GetUserIdByToken(string) (uint, error)
	//GetUserById(uint) (dataStruct.User, error)
	GetRecommendation(uint) ([]recommendation.Recommendation, error)
}

type Application struct {
	Router *http.ServeMux
	repo   IRepository
}

func New(repo IRepository) *Application {
	router := http.NewServeMux()
	a := &Application{repo: repo, Router: router}

	return a
}
