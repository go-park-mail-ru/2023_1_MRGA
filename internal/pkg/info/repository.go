package info

type IRepositoryInfo interface {
	GetHashtagId(nameHashtag []string) ([]uint, error)
	GetEducationId(nameEducation string) (uint, error)
	GetJobId(nameJob string) (uint, error)
	GetZodiacId(nameZodiac string) (uint, error)
	GetCityId(nameCity string) (uint, error)
	GetStatusId(status string) (uint, error)

	GetHashtags() ([]string, error)
	GetStatuses() ([]string, error)
	GetReasons() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetEducation() ([]string, error)
	GetZodiac() ([]string, error)
}
