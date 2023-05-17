package info_user

type UseCase interface {
	AddInfo(userId uint, info InfoStruct) error
	ChangeInfo(userId uint, info InfoChange) (InfoStructAnswer, error)
	GetInfo(userId uint) (InfoStructAnswer, error)

	GetUserById(uint) (UserRes, error)

	AddHashtags(userId uint, inp HashtagInp) error
	GetUserHashtags(userId uint) ([]string, error)
	ChangeUserHashtags(userId uint, inp HashtagInp) error

	GetUserStatus(userId uint) (string, error)
	ChangeUserStatus(userId uint, statusInp StatusInp) error

	GetUserHashtagsId(userId uint) ([]uint, error)
}
