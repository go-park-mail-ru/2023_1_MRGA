package app

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type IRepository interface {
	DeleteToken(string) error
	AddUser(d *ds.User) error
	Login(email string, username string, pass string) (uint, error)
	SaveToken(id uint, token string)
	GetCities() ([]string, error)
	GetUserIdByToken(string) (uint, error)
	GetUserById(uint) (*ds.User, error)
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
