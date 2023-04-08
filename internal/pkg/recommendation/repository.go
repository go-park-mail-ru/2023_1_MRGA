package recommendation

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryRec interface {
	//GetRecommendation(uint) ([]Recommendation, error)
	GetReasonId(reason string) (uint, error)
	GetReasons() ([]dataStruct.Reason, error)
	AddUserReason(reason *dataStruct.UserReason) error
	AddFilter(filter *dataStruct.UserFilter) error
	GetFilter(userId uint) (dataStruct.UserFilter, error)
	GetUserReasons(userId uint) ([]dataStruct.UserReason, error)
	GetReasonById(reasonId uint) (string, error)
	DeleteUserReason(userId, reactionId uint) error
	ChangeFilter(newFilter dataStruct.UserFilter) error
}
