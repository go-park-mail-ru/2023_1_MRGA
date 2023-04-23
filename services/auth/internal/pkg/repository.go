package authRepo

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/pkg/data_struct"

type UserRepo interface {
	DeleteToken(string) error
	AddUser(d *dataStruct.User) (uint, error)
	Login(input string, pass string) (uint, error)
	GetUserIdByToken(string) (uint, error)
	ChangeUser(user dataStruct.User) error
	SaveToken(id uint, token string) error
	CheckSession(token string) (uint, error)
	GetAge(userId uint) (int, error)
}
