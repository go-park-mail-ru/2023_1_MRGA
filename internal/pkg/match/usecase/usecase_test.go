package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/mocks"
)

func TestNewMatchUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	if matchUsecase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestMatchUseCase_GetMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	userId := uint(1)
	matches := []dataStruct.Match{
		{Id: uint(2),
			UserFirstId:  userId,
			UserSecondId: uint(2),
			Shown:        false},
	}
	userMatch := []match.UserRes{
		{UserId: uint(2),
			Name:  "test",
			Age:   20,
			Photo: uint(1),
			Shown: false},
	}

	matchRepoMock.EXPECT().GetMatches(userId).Return(matches, nil)
	matchRepoMock.EXPECT().ChangeStatusMatch(userId, uint(2)).Return(nil)
	matchRepoMock.EXPECT().GetUser(uint(2)).Return(userMatch[0], nil)
	matchRepoMock.EXPECT().GetAge(userMatch[0].UserId).Return(userMatch[0].Age, nil)

	result, err := matchUsecase.GetMatches(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, userMatch)
}

func TestMatchUseCase_GetMatches_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)

	matchRepoMock.EXPECT().GetMatches(userId).Return(nil, errRepo)

	_, err := matchUsecase.GetMatches(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestMatchUseCase_PostReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	userId := uint(1)
	EvaluatedUserId := uint(2)
	reactionInp := match.ReactionInp{
		EvaluatedUserId: EvaluatedUserId,
		Reaction:        "like",
	}
	todayData := time.Now().Format("2006-01-02")
	historyRow := dataStruct.UserHistory{
		UserId:        userId,
		UserProfileId: EvaluatedUserId,
		ShowDate:      todayData,
	}
	reaction2 := dataStruct.UserReaction{
		UserId:     EvaluatedUserId,
		UserFromId: userId,
		ReactionId: 1,
	}

	match1 := dataStruct.Match{
		Shown:        false,
		UserFirstId:  userId,
		UserSecondId: EvaluatedUserId,
	}

	match2 := dataStruct.Match{
		Shown:        false,
		UserFirstId:  EvaluatedUserId,
		UserSecondId: userId,
	}

	matchRepoMock.EXPECT().CheckCountReaction(userId).Return(true, nil)
	matchRepoMock.EXPECT().IncrementLikeCount(userId).Return(nil)
	matchRepoMock.EXPECT().GetIdReaction(reactionInp.Reaction).Return(uint(1), nil)
	matchRepoMock.EXPECT().AddHistoryRow(historyRow).Return(nil)
	matchRepoMock.EXPECT().GetUserReaction(EvaluatedUserId, userId).Return(reaction2, nil)
	matchRepoMock.EXPECT().AddMatchRow(match1).Return(nil)
	matchRepoMock.EXPECT().AddMatchRow(match2).Return(nil)
	matchRepoMock.EXPECT().DeleteUserReaction(uint(0)).Return(nil)

	_, err := matchUsecase.PostReaction(userId, reactionInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestMatchUseCase_PostReaction_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	EvaluatedUserId := uint(2)
	reactionInp := match.ReactionInp{
		EvaluatedUserId: EvaluatedUserId,
		Reaction:        "like",
	}
	todayData := time.Now().Format("2006-01-02")
	historyRow := dataStruct.UserHistory{
		UserId:        userId,
		UserProfileId: EvaluatedUserId,
		ShowDate:      todayData,
	}
	reaction2 := dataStruct.UserReaction{
		UserId:     EvaluatedUserId,
		UserFromId: userId,
		ReactionId: 1,
	}

	match1 := dataStruct.Match{
		Shown:        false,
		UserFirstId:  userId,
		UserSecondId: EvaluatedUserId,
	}

	matchRepoMock.EXPECT().CheckCountReaction(userId).Return(true, nil)
	matchRepoMock.EXPECT().IncrementLikeCount(userId).Return(nil)
	matchRepoMock.EXPECT().GetIdReaction(reactionInp.Reaction).Return(uint(1), nil)
	matchRepoMock.EXPECT().AddHistoryRow(historyRow).Return(nil)
	matchRepoMock.EXPECT().GetUserReaction(EvaluatedUserId, userId).Return(reaction2, nil)
	matchRepoMock.EXPECT().AddMatchRow(match1).Return(errRepo)

	_, err := matchUsecase.PostReaction(userId, reactionInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestMatchUseCase_DeleteMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	userId := uint(1)
	matchId := uint(1)
	matchRepoMock.EXPECT().DeleteMatch(userId, matchId).Return(nil)

	err := matchUsecase.DeleteMatch(userId, matchId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestMatchUseCase_DeleteMatch_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchRepoMock := mock.NewMockIRepositoryMatch(ctrl)
	matchUsecase := NewMatchUseCase(matchRepoMock)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	matchId := uint(1)
	matchRepoMock.EXPECT().DeleteMatch(userId, matchId).Return(errRepo)

	err := matchUsecase.DeleteMatch(userId, matchId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}
