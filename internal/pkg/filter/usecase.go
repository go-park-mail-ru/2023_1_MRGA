package filter

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	AddFilters(userId uint, FilterInp FilterInput) error
	GetFilters(userId uint) (FilterInput, error)
	ChangeFilters(userId uint, filterInp FilterInput) error

	GetUserReasonsId(userId uint) ([]uint, error)
	GetUserFilters(userId uint) (dataStruct.UserFilter, error)
}
