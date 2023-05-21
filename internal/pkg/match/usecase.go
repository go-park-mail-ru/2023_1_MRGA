package match

type UseCase interface {
	GetMatches(userId uint) ([]UserRes, error)
	PostReaction(userId uint, reaction ReactionInp) (ReactionResult, error)
	DeleteMatch(userId, userMatchId uint) error

	GetChatByEmail(usrId uint, userId uint) (ChatAnswer, error)
}
