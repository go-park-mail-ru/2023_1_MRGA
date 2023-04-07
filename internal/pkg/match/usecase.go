package match

type UseCase interface {
	GetMatches(userId uint) ([]UserRes, error)
	PostReaction(userId, userTo uint, reaction string) error
}
