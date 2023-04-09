package match

type UseCase interface {
	GetMatches(userId uint) ([]UserRes, error)
	PostReaction(userId uint, reaction ReactionInp) error

	GetChatByEmail(usrId uint, email string) (ChatAnswer, error)
}
