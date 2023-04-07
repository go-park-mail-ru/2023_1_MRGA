package info

type UseCase interface {
	AddInfo(userId uint, info InfoStruct) error
	ChangeInfo(userId uint, info InfoChange) (InfoStruct, error)
	GetInfo(userId uint) (InfoStruct, error)
	///getters
	GetZodiacs() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetEducation() ([]string, error)
}
