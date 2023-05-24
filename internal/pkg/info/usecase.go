package info

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	GetHashtags() ([]string, error)
	GetZodiacs() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetReasons() ([]string, error)
	GetEducation() ([]string, error)
	GetStatuses() ([]string, error)

	GetCityId(city string) (uint, error)
	GetZodiacId(zodiac string) (uint, error)
	GetEducationId(education string) (uint, error)
	GetJobId(job string) (uint, error)
	GetHashtagId(hashtagId []string) ([]uint, error)
	GetStatusId(status string) (uint, error)
}
