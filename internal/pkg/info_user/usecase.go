package info_user

type UseCase interface {
	AddInfo(userId uint, info InfoStruct) error
	ChangeInfo(userId uint, info InfoChange) (InfoStructAnswer, error)
	GetInfo(userId uint) (InfoStructAnswer, error)
	GetInfoByEmail(userId uint) (userInfo InfoStructAnswer, err error)

	AddHashtags(userId uint, inp HashtagInp) error
	GetUserHashtags(userId uint) (HashtagInp, error)
	ChangeUserHashtags(userId uint, inp HashtagInp) error

	GetHashtags() ([]string, error)
	GetZodiacs() ([]string, error)
	GetCities() ([]string, error)
	GetJobs() ([]string, error)
	GetEducation() ([]string, error)
}
