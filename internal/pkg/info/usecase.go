package info

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type UseCase interface {
	AddInfo(user *dataStruct.User) (string, error)
	ChangeInfo(logInp LoginInput) (string, error)
	Logout(string) error
	GetUserById(uint) (UserRes, error)
}
