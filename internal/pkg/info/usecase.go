package info

type UseCase interface {
	GetHashtags() ([]string, error)
	GetZodiacs() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetEducation() ([]string, error)
}
