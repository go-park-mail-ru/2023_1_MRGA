package auth

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type UseCase interface {
	Register(user *dataStruct.User) (string, error)
	Login(logInp LoginInput) (string, error)
	Logout(string) error
	GetUserById(uint) (UserRes, error)
}
