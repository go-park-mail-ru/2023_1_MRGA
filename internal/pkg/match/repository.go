package match

type IRepositoryMatch interface {
	PostReaction(UserId, UserTo uint, reactionId uint) error
	ReactionIdByName(reaction string) (uint, error)
	GetMatches(userId uint) ([]UserRes, error)
}
