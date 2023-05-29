package match

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock UseCase
type UseCase interface {
	GetMatches(userId uint) ([]UserRes, error)
	PostReaction(userId uint, reaction ReactionInp) (ReactionResult, error)
	DeleteMatch(userId, userMatchId uint) error
}
