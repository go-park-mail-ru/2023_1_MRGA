package info

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryInfo interface {
	GetHashtagId(nameHashtag []string) ([]uint, error)
	GetEducationId(nameEducation string) (uint, error)
	GetJobId(nameJob string) (uint, error)
	GetZodiacId(nameZodiac string) (uint, error)
	GetCityId(nameCity string) (uint, error)

	GetHashtags() ([]dataStruct.Hashtag, error)
	GetReasons() ([]dataStruct.Reason, error)
	GetCities() ([]dataStruct.City, error)
	GetJobs() ([]dataStruct.Job, error)
	GetEducation() ([]dataStruct.Education, error)
	GetZodiac() ([]dataStruct.Zodiac, error)
}
