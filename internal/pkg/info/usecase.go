package info

type UseCase interface {
	GetHashtags() ([]string, error)
	GetZodiacs() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetReasons() ([]string, error)
	GetEducation() ([]string, error)

	GetCityId(city string) (uint, error)
	GetZodiacId(zodiac string) (uint, error)
	GetEducationId(education string) (uint, error)
	GetJobId(job string) (uint, error)
	GetHashtagId(hashtagId []string) ([]uint, error)
}
