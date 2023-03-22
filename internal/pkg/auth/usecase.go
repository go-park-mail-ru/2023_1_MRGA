package auth

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
)

type UseCase interface {
	Register(user *dataStruct.User) (string, error)
	Login(logInp delivery.LoginInput) (string, error)
	Logout(string) error
}
