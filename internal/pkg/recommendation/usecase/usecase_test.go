package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	filterMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/mocks"
	infoUserMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/mocks"
	photoMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/mocks"
)

func TestNewRecUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recRepoMock := mock.NewMockIRepositoryRec(ctrl)
	filterUsecase := filterMock.NewMockUseCase(ctrl)
	photoUsecase := photoMock.NewMockUseCase(ctrl)
	infoUserUsecase := infoUserMock.NewMockUseCase(ctrl)
	recUsecase := NewRecUseCase(recRepoMock, filterUsecase, photoUsecase, infoUserUsecase)

	if recUsecase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestRecUseCase_GetRecommendations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recRepoMock := mock.NewMockIRepositoryRec(ctrl)
	filterUsecase := filterMock.NewMockUseCase(ctrl)
	photoUsecase := photoMock.NewMockUseCase(ctrl)
	infoUserUsecase := infoUserMock.NewMockUseCase(ctrl)
	recUsecase := NewRecUseCase(recRepoMock, filterUsecase, photoUsecase, infoUserUsecase)

	userId := uint(1)
	hashtags := []uint{1, 2}
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MaxAge:    40,
		MinAge:    20,
		SearchSex: uint(1),
	}
	rescUser := []recommendation.UserRecommend{
		{
			UserId: uint(2),
		},
	}
	recs := []recommendation.Recommendation{
		{
			Id:          uint(2),
			Name:        "test1",
			Job:         "test2",
			Photos:      []uint{1, 2, 3},
			Age:         25,
			Sex:         constform.Female,
			Description: "test",
			Hashtags:    []string{"test1", "test2"},
			Zodiac:      "test",
			Education:   "test",
		},
	}
	infoUserUsecase.EXPECT().GetUserHashtagsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserReasonsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserFilters(userId).Return(userFilter, nil)
	recRepoMock.EXPECT().GetUserHistory(userId).Return([]uint{3}, nil)
	recRepoMock.EXPECT().GetRecommendation(userId, []uint{3}, hashtags, hashtags, userFilter).Return(rescUser, nil)
	recRepoMock.EXPECT().GetRecommendedUser(recs[0].Id).Return(recs[0], nil)
	infoUserUsecase.EXPECT().GetUserHashtags(recs[0].Id).Return(recs[0].Hashtags, nil)
	photoUsecase.EXPECT().GetAllPhotos(recs[0].Id).Return(recs[0].Photos, nil)

	result, err := recUsecase.GetRecommendations(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, recs)
}

func TestRecUseCase_GetRecommendations_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recRepoMock := mock.NewMockIRepositoryRec(ctrl)
	filterUsecase := filterMock.NewMockUseCase(ctrl)
	photoUsecase := photoMock.NewMockUseCase(ctrl)
	infoUserUsecase := infoUserMock.NewMockUseCase(ctrl)
	recUsecase := NewRecUseCase(recRepoMock, filterUsecase, photoUsecase, infoUserUsecase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	hashtags := []uint{1, 2}
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MaxAge:    40,
		MinAge:    20,
		SearchSex: uint(1),
	}
	infoUserUsecase.EXPECT().GetUserHashtagsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserReasonsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserFilters(userId).Return(userFilter, nil)
	recRepoMock.EXPECT().GetUserHistory(userId).Return([]uint{3}, nil)
	recRepoMock.EXPECT().GetRecommendation(userId, []uint{3}, hashtags, hashtags, userFilter).Return(nil, errRepo)

	_, err := recUsecase.GetRecommendations(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestRecUseCase_CheckProStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recRepoMock := mock.NewMockIRepositoryRec(ctrl)
	filterUsecase := filterMock.NewMockUseCase(ctrl)
	photoUsecase := photoMock.NewMockUseCase(ctrl)
	infoUserUsecase := infoUserMock.NewMockUseCase(ctrl)
	recUsecase := NewRecUseCase(recRepoMock, filterUsecase, photoUsecase, infoUserUsecase)

	userId := uint(1)
	hashtags := []uint{1, 2}
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MaxAge:    40,
		MinAge:    20,
		SearchSex: uint(1),
	}
	rescUser := []recommendation.UserRecommend{
		{
			UserId: uint(2),
		},
	}
	recs := []recommendation.Recommendation{
		{
			Id:          uint(2),
			Name:        "test1",
			Job:         "test2",
			Photos:      []uint{1, 2, 3},
			Age:         25,
			Sex:         constform.Female,
			Description: "test",
			Hashtags:    []string{"test1", "test2"},
			Zodiac:      "test",
			Education:   "test",
		},
	}

	filterUsecase.EXPECT().GetUserReasonsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserFilters(userId).Return(userFilter, nil)
	recRepoMock.EXPECT().GetUserHistory(userId).Return([]uint{3}, nil)
	recRepoMock.EXPECT().GetLikes(userId, []uint{3}, hashtags, userFilter).Return(rescUser, nil)
	recRepoMock.EXPECT().GetRecommendedUser(recs[0].Id).Return(recs[0], nil)
	infoUserUsecase.EXPECT().GetUserHashtags(recs[0].Id).Return(recs[0].Hashtags, nil)
	photoUsecase.EXPECT().GetAllPhotos(recs[0].Id).Return(recs[0].Photos, nil)

	result, err := recUsecase.GetLikes(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, recs)
}

func TestRecUseCase_GetLikes_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recRepoMock := mock.NewMockIRepositoryRec(ctrl)
	filterUsecase := filterMock.NewMockUseCase(ctrl)
	photoUsecase := photoMock.NewMockUseCase(ctrl)
	infoUserUsecase := infoUserMock.NewMockUseCase(ctrl)
	recUsecase := NewRecUseCase(recRepoMock, filterUsecase, photoUsecase, infoUserUsecase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	hashtags := []uint{1, 2}
	userFilter := dataStruct.UserFilter{
		UserId:    userId,
		MaxAge:    40,
		MinAge:    20,
		SearchSex: uint(1),
	}
	filterUsecase.EXPECT().GetUserReasonsId(userId).Return(hashtags, nil)
	filterUsecase.EXPECT().GetUserFilters(userId).Return(userFilter, nil)
	recRepoMock.EXPECT().GetUserHistory(userId).Return([]uint{3}, nil)
	recRepoMock.EXPECT().GetLikes(userId, []uint{3}, hashtags, userFilter).Return(nil, errRepo)

	_, err := recUsecase.GetLikes(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}
