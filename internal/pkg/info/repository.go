package info

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryInfo interface {
	GetHashtagId(nameHashtag []string) ([]uint, error) //delete
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
