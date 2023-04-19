package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
)

type FilterUseCase struct {
	Repo            filter.IRepositoryFilter
	infoUseCase     info.UseCase
	infoUserUseCase info_user.UseCase
}

func NewFilterUseCase(userRepo filter.IRepositoryFilter, infoUC info.UseCase, infoUserUC info_user.UseCase) *FilterUseCase {
	return &FilterUseCase{
		Repo:            userRepo,
		infoUseCase:     infoUC,
		infoUserUseCase: infoUserUC,
	}
}

func (f *FilterUseCase) AddFilters(userId uint, filterInp filter.FilterInput) error {
	reasonId, err := f.Repo.GetReasonId(filterInp.Reason)
	if err != nil {
		return err
	}
	var addReason []dataStruct.UserReason
	for _, reason := range reasonId {
		var userReason dataStruct.UserReason
		userReason.UserId = userId
		userReason.ReasonId = reason
		addReason = append(addReason, userReason)
	}
	err = f.Repo.AddUserReason(addReason)
	if err != nil {
		return err
	}

	var filter dataStruct.UserFilter
	filter.UserId = userId
	filter.SearchSex = filterInp.SearchSex
	filter.MinAge = filterInp.MinAge
	filter.MaxAge = filterInp.MaxAge

	err = f.Repo.AddFilter(&filter)
	if err != nil {
		return err
	}
	return nil
}

func (f *FilterUseCase) GetFilters(userId uint) (filter.FilterInput, error) {
	var filterRes filter.FilterInput
	filterUser, err := f.Repo.GetFilter(userId)
	if err != nil {
		return filter.FilterInput{}, err
	}

	filterRes.MinAge = filterUser.MinAge
	filterRes.MaxAge = filterUser.MaxAge
	filterRes.SearchSex = filterUser.SearchSex

	reasons, err := f.Repo.GetUserReasons(userId)
	if err != nil {
		return filter.FilterInput{}, err
	}

	filterRes.Reason = reasons

	return filterRes, nil
}

func (f *FilterUseCase) ChangeFilters(userId uint, filterInp filter.FilterInput) error {
	reasonsBD, err := f.Repo.GetUserReasonsId(userId)
	if err != nil {
		return err
	}

	reasonsId, err := f.Repo.GetReasonId(filterInp.Reason)
	if err != nil {
		return err
	}

	var addReasons []dataStruct.UserReason
	for _, reason := range reasonsId {
		if !Contains(reasonsBD, reason) {

			var reasonAdd dataStruct.UserReason
			reasonAdd.UserId = userId
			reasonAdd.ReasonId = reason
			addReasons = append(addReasons, reasonAdd)
		}
	}

	err = f.Repo.AddUserReason(addReasons)
	if err != nil {
		return err
	}

	var deleteReason []uint
	for _, reason := range reasonsBD {
		if !Contains(reasonsId, reason) {
			deleteReason = append(deleteReason, reason)
		}
	}
	err = f.Repo.DeleteUserReason(userId, deleteReason)
	if err != nil {
		return err
	}

	var newFilter dataStruct.UserFilter
	newFilter.UserId = userId
	newFilter.MaxAge = filterInp.MaxAge
	newFilter.SearchSex = filterInp.SearchSex

	err = f.Repo.ChangeFilter(newFilter)
	if err != nil {
		return err
	}
	return nil
}

func (f *FilterUseCase) GetUserReasonsId(userId uint) ([]uint, error) {
	reasonsBD, err := f.Repo.GetUserReasonsId(userId)
	if err != nil {
		return nil, err
	}
	return reasonsBD, err
}

func (f *FilterUseCase) GetUserFilters(userId uint) (dataStruct.UserFilter, error) {
	filterUser, err := f.Repo.GetFilter(userId)
	if err != nil {
		return dataStruct.UserFilter{}, err
	}

	return filterUser, nil
}

func Contains(s []uint, elem uint) bool {
	for _, elemS := range s {
		if elem == elemS {
			return true
		}
	}
	return false
}
