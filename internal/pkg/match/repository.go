package match

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryMatch interface {
	GetMatches(userId uint) ([]dataStruct.Match, error)
	GetUser(userId uint) (user UserRes, err error)
	GetIdReaction(reaction string) (uint, error)
	AddHistoryRow(row dataStruct.UserHistory) error
	GetUserReaction(userId, userToId uint) (dataStruct.UserReaction, error)
	AddUserReaction(row dataStruct.UserReaction) error
	DeleteUserReaction(rowId uint) error
	AddMatchRow(row dataStruct.Match) error
	ChangeStatusMatch(userId, profileId uint) error
	GetChat(userId uint) (ChatAnswer, error)
}
