package filter

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

//go:generate mockgen -source=repository.go -destination=mocks/repo.go -package=mock
type IRepositoryFilter interface {
	GetReasonById(reasonId []uint) ([]string, error)
	GetReasonId(reason []string) ([]uint, error)
	AddUserReason(reason []dataStruct.UserReason) error

	GetUserReasonsId(userId uint) ([]uint, error)
	GetUserReasons(userId uint) ([]string, error)
	DeleteUserReason(userId uint, reactionId []uint) error

	AddFilter(filter *dataStruct.UserFilter) error
	GetFilter(userId uint) (dataStruct.UserFilter, error)
	ChangeFilter(newFilter dataStruct.UserFilter) error
}
