package match

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	GetMatches(userId uint) ([]UserRes, error)
	PostReaction(userId uint, reaction ReactionInp) error
	DeleteMatch(userId, userMatchId uint) error

	GetChatByEmail(usrId uint, userId uint) (ChatAnswer, error)
}
