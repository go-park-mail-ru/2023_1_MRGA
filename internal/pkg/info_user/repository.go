package info_user

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryInfo interface {
	AddInfoUser(userInfo *dataStruct.UserInfo) error
	AddUserPhoto(userPhoto *dataStruct.UserPhoto) error  //delete
	GetUserInfo(userId uint) (InfoStruct, error)         //ok
	GetUserPhoto(userId uint) (photos []uint, err error) //photo
	ChangeInfo(userInfo *dataStruct.UserInfo) error

	GetUserIdByEmail(email string) (uint, error)
	GetAge(userId uint) (int, error)     //ok
	GetAvatar(userId uint) (uint, error) //photo

	GetUserHashtags(userId uint) ([]uint, error)
	AddUserHashtag(hashtag []dataStruct.UserHashtag) error
	GetHashtagById(hashtagId []uint) ([]string, error) //delete
	DeleteUserHashtag(userId uint, hashtagId []uint) error
}
