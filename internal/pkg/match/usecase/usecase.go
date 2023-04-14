package usecase

import (
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
		matchUser, err := m.userRepo.GetUser(user.UserSecondId)
		matchUser.UserId = user.UserSecondId
		matchUser.Shown = user.Shown
		if err != nil {
			return nil, err
		}
		result = append(result, matchUser)
	}
	return result, nil

}

func (m *MatchUseCase) PostReaction(userId uint, reaction match.ReactionInp) error {
	userToId := reaction.EvaluatedUserId

	reactionId, err := m.userRepo.GetIdReaction(reaction.Reaction)
	if err != nil {
		return err
	}

	var historyRow dataStruct.UserHistory
	historyRow.UserId = userId
	historyRow.UserProfileId = userToId
	todayData := time.Now().Format("2006-01-02")
	historyRow.ShowDate = todayData
	err = m.userRepo.AddHistoryRow(historyRow)
	if err != nil {
		return err
	}

	reactionUser, err := m.userRepo.GetUserReaction(userToId, userId)
	if err == gorm.ErrRecordNotFound {
		var reactionRow dataStruct.UserReaction
		reactionRow.UserId = userId
		reactionRow.UserFromId = userToId
		reactionRow.ReactionId = reactionId
		err = m.userRepo.AddUserReaction(reactionRow)
		if err != nil {
			return err
		}
	} else {
		if reactionUser.ReactionId == 1 && reactionId == 1 {
			var match1 dataStruct.Match
			match1.Shown = false
			match1.UserFirstId = userId
			match1.UserSecondId = userToId
			err = m.userRepo.AddMatchRow(match1)
			if err != nil {
				return err
			}

			var match2 dataStruct.Match
			match2.Shown = false
			match2.UserFirstId = userToId
			match2.UserSecondId = userId
			err = m.userRepo.AddMatchRow(match2)
			if err != nil {
				return err
			}
		}
		err = m.userRepo.DeleteUserReaction(reactionUser.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MatchUseCase) GetChatByEmail(userId uint, matchUserId uint) (result match.ChatAnswer, err error) {
	if err != nil {
		return
	}

	err = m.userRepo.ChangeStatusMatch(userId, matchUserId)
	if err != nil {
		return
	}
	result, err = m.userRepo.GetChat(matchUserId)
	return
}
