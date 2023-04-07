package info

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryInfo interface {
	AddInfoUser(userInfo *dataStruct.UserInfo) error
	AddUserPhoto(userPhoto *dataStruct.UserPhoto) error
	GetUserInfo(userId uint) (InfoStruct, error)
	GetAvatar(userId uint) (string, error)
	GetPhotos(userId uint) ([]dataStruct.UserPhoto, error)

	GetEducationId(nameEducation string) (uint, error)
	GetJobId(nameJob string) (uint, error)
	GetZodiacId(nameZodiac string) (uint, error)
	GetCityId(nameCity string) (uint, error)

	GetCities() ([]dataStruct.City, error)
	GetJobs() ([]dataStruct.Job, error)
	GetEducation() ([]dataStruct.Education, error)
	GetZodiac() ([]dataStruct.Zodiac, error)
}
