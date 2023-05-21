package info_user

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	AddInfo(userId uint, info InfoStruct) error                        //test
	ChangeInfo(userId uint, info InfoChange) (InfoStructAnswer, error) //test
	GetInfo(userId uint) (InfoStructAnswer, error)                     //test

	GetUserById(uint) (UserRes, error) //test

	AddHashtags(userId uint, inp HashtagInp) error //test
	GetUserHashtags(userId uint) ([]string, error) //test
	ChangeUserHashtags(userId uint, inp HashtagInp) error

	GetUserStatus(userId uint) (string, error) //test
	ChangeUserStatus(userId uint, statusInp StatusInp) error

	GetUserHashtagsId(userId uint) ([]uint, error) //test
}
