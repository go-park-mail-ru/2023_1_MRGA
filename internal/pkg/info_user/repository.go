package info_user

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

//go:generate mockgen -source=repository.go -destination=mocks/repo.go -package=mock
type IRepositoryInfo interface {
	AddInfoUser(userInfo *dataStruct.UserInfo) error
	GetUserInfo(userId uint) (InfoStruct, error)
	ChangeInfo(userInfo *dataStruct.UserInfo) error
	GetAge(userId uint) (int, error)

	GetUserById(uint) (UserRestTemp, error)
	CheckFilter(userId uint) (bool, error)

	GetUserHashtagsId(userId uint) ([]uint, error)
	GetUserHashtags(userId uint) ([]string, error)
	AddUserHashtag(hashtag []dataStruct.UserHashtag) error
	DeleteUserHashtag(userId uint, hashtagId []uint) error

	GetUserStatus(userId uint) (string, error)
	ChangeUserStatus(userId, statusId uint) error
}
