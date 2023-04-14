package recommendation

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryRec interface {
	GetRecommendation(userId uint, history []uint, reasons []uint, hashtags []uint, filters dataStruct.UserFilter) (users []UserRecommend, err error)
	GetRecommendedUser(userId uint) (user Recommendation, err error)

	GetUserAge(userId uint) (int, error)
	GetUserHashtags(userId uint) ([]dataStruct.UserHashtag, error)
	GetUserNameHashtags(userId uint) ([]string, error)
	GetUserHistory(userId uint) ([]dataStruct.UserHistory, error)

	GetPhotos(userId uint) ([]Photo, error)

	GetReasonId(reason string) (uint, error)
	GetReasons() ([]dataStruct.Reason, error)
	AddUserReason(reason *dataStruct.UserReason) error

	GetUserReasons(userId uint) ([]dataStruct.UserReason, error)
	GetReasonById(reasonId uint) (string, error)
	DeleteUserReason(userId, reactionId uint) error

	AddFilter(filter *dataStruct.UserFilter) error
	GetFilter(userId uint) (dataStruct.UserFilter, error)
	ChangeFilter(newFilter dataStruct.UserFilter) error
}
