package usecase

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type MatchUseCase struct {
	userRepo match.IRepositoryMatch
}

func NewMatchUseCase(userRepo match.IRepositoryMatch) *MatchUseCase {
	return &MatchUseCase{
		userRepo: userRepo,
	}
}

func (m *MatchUseCase) GetMatches(userId uint) ([]match.UserRes, error) {
	users, err := m.userRepo.GetMatches(userId)

	if err != nil {
		return nil, err
	}
	var result []match.UserRes
	for _, user := range users {
		err = m.userRepo.ChangeStatusMatch(user.UserFirstId, user.UserSecondId)
		if err != nil {
			return nil, err
		}
		matchUser, err := m.userRepo.GetUser(user.UserSecondId)
		if err != nil {
			return nil, err
		}
		matchUser.UserId = user.UserSecondId
		matchUser.Shown = user.Shown
		age, err := m.userRepo.GetAge(user.UserSecondId)
		if err != nil {
			return nil, err
		}
		matchUser.Age = age
		result = append(result, matchUser)
	}
	return result, nil

}

func (m *MatchUseCase) PostReaction(userId uint, reaction match.ReactionInp) (result match.ReactionResult, err error) {
	if reaction.Reaction == "like" {
		ok, err := m.userRepo.CheckCountReaction(userId)
		if err != nil {
			return result, err
		}
		if !ok {
			return result, fmt.Errorf("you cant like people today. Try tomorrow")
		}
		err = m.userRepo.IncrementLikeCount(userId)
		if err != nil {
			return result, err
		}
	}

	userToId := reaction.EvaluatedUserId

	currentUserReactionCode, err := m.userRepo.GetIdReaction(reaction.Reaction)
	if err != nil {
		return result, err
	}

	var historyRow dataStruct.UserHistory
	historyRow.UserId = userId
	historyRow.UserProfileId = userToId
	todayData := time.Now().Format("2006-01-02")
	historyRow.ShowDate = todayData
	err = m.userRepo.AddHistoryRow(historyRow)
	if err != nil {
		return result, err
	}

	sideUserReaction, err := m.userRepo.GetUserReaction(userToId, userId)
	if err == gorm.ErrRecordNotFound {
		var reactionRow dataStruct.UserReaction
		reactionRow.UserId = userId
		reactionRow.UserFromId = userToId
		reactionRow.ReactionId = currentUserReactionCode
		err = m.userRepo.AddUserReaction(reactionRow)
		return match.ReactionResult{ResultCode: match.FirstReaction}, err
	}
	var like uint = dataStruct.LikeReaction
	var pass uint = dataStruct.PassReaction
	if sideUserReaction.ReactionId == like && currentUserReactionCode == like {
		var match1 dataStruct.Match
		match1.Shown = false
		match1.UserFirstId = userId
		match1.UserSecondId = userToId
		err = m.userRepo.AddMatchRow(match1)
		if err != nil {
			return result, err
		}

		var match2 dataStruct.Match
		match2.Shown = false
		match2.UserFirstId = userToId
		match2.UserSecondId = userId
		err = m.userRepo.AddMatchRow(match2)
		if err != nil {
			return result, err
		}
		result.ResultCode = match.NewMatch
	}
	fmt.Printf("sideUserReaction.ReactionId: %v, currentUserReactionCode:%v",
		sideUserReaction.ReactionId, currentUserReactionCode)
	if sideUserReaction.ReactionId == like && currentUserReactionCode == pass {
		result.ResultCode = match.MissedMatch
	}
	err = m.userRepo.DeleteUserReaction(sideUserReaction.Id)
	return result, err
}

func (m *MatchUseCase) DeleteMatch(userId uint, matchUserId uint) error {
	err := m.userRepo.DeleteMatch(userId, matchUserId)
	if err != nil {
		return err
	}

	return nil
}
