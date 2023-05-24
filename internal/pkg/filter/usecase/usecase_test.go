package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/mocks"
	infoMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/mocks"
	infouserMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/mocks"
)

func TestNewFilterUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	if filterUseCase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestFilterUseCase_GetFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)

	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	reasons := []string{
		"test1",
		"test2",
	}

	filterRes := filter.FilterInput{
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
		Reason:    reasons,
	}

	filterRepoMock.EXPECT().GetFilter(userId).Return(userFilter, nil)
	filterRepoMock.EXPECT().GetUserReasons(userId).Return(reasons, nil)

	result, err := filterUseCase.GetFilters(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, filterRes)
}

func TestFilterUseCase_GetFilters_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	filterRepoMock.EXPECT().GetFilter(userId).Return(userFilter, nil)
	filterRepoMock.EXPECT().GetUserReasons(userId).Return(nil, errRepo)

	_, err := filterUseCase.GetFilters(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestFilterUseCase_AddFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)

	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	reasons := []string{
		"test1",
		"test2",
	}

	filterRes := filter.FilterInput{
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
		Reason:    reasons,
	}
	reasonsId := []uint{1, 2}
	reasonsAdd := []dataStruct.UserReason{
		{
			UserId:   userId,
			ReasonId: reasonsId[0],
		},
		{
			UserId:   userId,
			ReasonId: reasonsId[1],
		},
	}

	filterRepoMock.EXPECT().GetReasonId(reasons).Return(reasonsId, nil)
	filterRepoMock.EXPECT().AddUserReason(reasonsAdd).Return(nil)
	filterRepoMock.EXPECT().AddFilter(&userFilter).Return(nil)

	err := filterUseCase.AddFilters(userId, filterRes)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestFilterUseCase_AddInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	reasons := []string{
		"test1",
		"test2",
	}

	filterRes := filter.FilterInput{
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
		Reason:    reasons,
	}
	reasonsId := []uint{1, 2}
	reasonsAdd := []dataStruct.UserReason{
		{
			UserId:   userId,
			ReasonId: reasonsId[0],
		},
		{
			UserId:   userId,
			ReasonId: reasonsId[1],
		},
	}

	filterRepoMock.EXPECT().GetReasonId(reasons).Return(reasonsId, nil)
	filterRepoMock.EXPECT().AddUserReason(reasonsAdd).Return(nil)
	filterRepoMock.EXPECT().AddFilter(&userFilter).Return(errRepo)

	err := filterUseCase.AddFilters(userId, filterRes)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestFilterUseCase_ChangeFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)

	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	reasons := []string{
		"test1",
		"test2",
	}

	filterRes := filter.FilterInput{
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
		Reason:    reasons,
	}
	reasonsId := []uint{1, 2}
	reasonsUser := []uint{2, 3}
	reasonsAdd := []dataStruct.UserReason{
		{
			UserId:   userId,
			ReasonId: reasonsId[0],
		},
	}

	filterRepoMock.EXPECT().GetReasonId(reasons).Return(reasonsId, nil)
	filterRepoMock.EXPECT().GetUserReasonsId(userId).Return(reasonsUser, nil)
	filterRepoMock.EXPECT().AddUserReason(reasonsAdd).Return(nil)
	filterRepoMock.EXPECT().DeleteUserReason(userId, []uint{3}).Return(nil)
	filterRepoMock.EXPECT().ChangeFilter(userFilter).Return(nil)

	err := filterUseCase.ChangeFilters(userId, filterRes)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestFilterUseCase_ChangeInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	reasons := []string{
		"test1",
		"test2",
	}

	filterRes := filter.FilterInput{
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
		Reason:    reasons,
	}
	reasonsId := []uint{1, 2}
	reasonsUser := []uint{2, 3}
	reasonsAdd := []dataStruct.UserReason{
		{
			UserId:   userId,
			ReasonId: reasonsId[0],
		},
	}

	filterRepoMock.EXPECT().GetReasonId(reasons).Return(reasonsId, nil)
	filterRepoMock.EXPECT().GetUserReasonsId(userId).Return(reasonsUser, nil)
	filterRepoMock.EXPECT().AddUserReason(reasonsAdd).Return(nil)
	filterRepoMock.EXPECT().DeleteUserReason(userId, []uint{3}).Return(nil)
	filterRepoMock.EXPECT().ChangeFilter(userFilter).Return(errRepo)

	err := filterUseCase.ChangeFilters(userId, filterRes)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestFilterUseCase_GetUserReasonsId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	reasons := []uint{1, 2}
	userId := uint(1)

	filterRepoMock.EXPECT().GetUserReasonsId(userId).Return(reasons, nil)

	result, err := filterUseCase.GetUserReasonsId(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, reasons, result)
}

func TestFilterUseCase_GetUserReasonsId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	errRepo := fmt.Errorf("something wrong")
	reasons := []uint{1, 2}
	userId := uint(1)

	filterRepoMock.EXPECT().GetUserReasonsId(userId).Return(reasons, errRepo)

	_, err := filterUseCase.GetUserReasonsId(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestFilterUseCase_GetUserFilters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)

	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	filterRepoMock.EXPECT().GetFilter(userId).Return(userFilter, nil)

	result, err := filterUseCase.GetUserFilters(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, userFilter)
}

func TestFilterUseCase_GetUserFilters_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterRepoMock := mock.NewMockIRepositoryFilter(ctrl)
	infoUserUseCase := infouserMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	filterUseCase := NewFilterUseCase(filterRepoMock, infoUseCase, infoUserUseCase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MinAge:    20,
		MaxAge:    40,
		SearchSex: uint(1),
	}

	filterRepoMock.EXPECT().GetFilter(userId).Return(userFilter, errRepo)

	_, err := filterUseCase.GetUserFilters(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}
