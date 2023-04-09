package info_user

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryInfo interface {
	AddInfoUser(userInfo *dataStruct.UserInfo) error
	AddUserPhoto(userPhoto *dataStruct.UserPhoto) error
	GetUserInfo(userId uint) (InfoStruct, error)
	GetAvatar(userId uint) (string, error)
	GetPhotos(userId uint) ([]dataStruct.UserPhoto, error)
	ChangeInfo(userInfo *dataStruct.UserInfo) error

	GetUserIdByEmail(email string) (uint, error)

	GetUserHashtags(userId uint) ([]dataStruct.UserHashtag, error)
	AddUserHashtag(hashtag dataStruct.UserHashtag) error
	GetHashtagById(hashtagId uint) (string, error)
	DeleteUserHashtag(userId, hashtagId uint) error

	GetHashtagId(nameHashtag string) (uint, error)
	GetEducationId(nameEducation string) (uint, error)
	GetJobId(nameJob string) (uint, error)
	GetZodiacId(nameZodiac string) (uint, error)
	GetCityId(nameCity string) (uint, error)

	GetHashtags() ([]dataStruct.Hashtag, error)
	GetCities() ([]dataStruct.City, error)
	GetJobs() ([]dataStruct.Job, error)
	GetEducation() ([]dataStruct.Education, error)
	GetZodiac() ([]dataStruct.Zodiac, error)
}
