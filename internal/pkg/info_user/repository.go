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

	GetUserHashtags(userId uint) ([]dataStruct.UserHashtag, error)
	AddUserHashtag(hashtag dataStruct.UserHashtag) error
	GetHashtagById(hashtagId uint) (string, error) //delete
	DeleteUserHashtag(userId, hashtagId uint) error

	GetHashtagId(nameHashtag string) (uint, error)     //delete
	GetEducationId(nameEducation string) (uint, error) //delete
	GetJobId(nameJob string) (uint, error)             //delete
	GetZodiacId(nameZodiac string) (uint, error)       //delete
	GetCityId(nameCity string) (uint, error)           //delete

	GetHashtags() ([]dataStruct.Hashtag, error)    //info
	GetCities() ([]dataStruct.City, error)         //info
	GetJobs() ([]dataStruct.Job, error)            //info
	GetEducation() ([]dataStruct.Education, error) //info
	GetZodiac() ([]dataStruct.Zodiac, error)       //info
}
