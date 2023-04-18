package info_user

type UseCase interface {
	AddInfo(userId uint, info InfoStruct) error                        //ok
	ChangeInfo(userId uint, info InfoChange) (InfoStructAnswer, error) //ok
	GetInfo(userId uint) (InfoStructAnswer, error)                     //fix

	AddHashtags(userId uint, inp HashtagInp) error
	GetUserHashtags(userId uint) (HashtagInp, error)
	ChangeUserHashtags(userId uint, inp HashtagInp) error
}
