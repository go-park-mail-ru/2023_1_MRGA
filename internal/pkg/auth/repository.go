package auth

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryAuth interface {
	DeleteToken(string) error
	AddUser(d *dataStruct.User) (uint, error)
	Login(input string, pass string) (uint, error)
	GetUserIdByToken(string) (uint, error)
	GetUserPhoto(userId uint) (photos []dataStruct.UserPhoto, err error)
	GetUserById(uint) (UserRestTemp, error)
	ChangeUser(user dataStruct.User) error
	SaveToken(id uint, token string) error
	GetAge(userId uint) (int, error)
}
