package match

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

//go:generate mockgen -source=repository.go -destination=mocks/repo.go -package=mock
type IRepositoryMatch interface {
	GetMatches(userId uint) ([]dataStruct.Match, error)
	GetUser(userId uint) (user UserRes, err error) //user_info
	GetIdReaction(reaction string) (uint, error)
	AddHistoryRow(row dataStruct.UserHistory) error
	GetUserReaction(userId, userToId uint) (dataStruct.UserReaction, error)
	AddUserReaction(row dataStruct.UserReaction) error
	DeleteUserReaction(rowId uint) error
	DeleteMatch(userId, userMatchId uint) error
	AddMatchRow(row dataStruct.Match) error
	ChangeStatusMatch(userId, profileId uint) error
	GetAge(userId uint) (int, error) //user_info
	CheckCountReaction(userId uint) (ok bool, err error)
	IncrementLikeCount(userId uint) error
}
