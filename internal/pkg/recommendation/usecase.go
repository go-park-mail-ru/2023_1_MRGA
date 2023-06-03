package recommendation

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	GetRecommendations(userId uint) ([]Recommendation, error)
	GetLikes(userId uint) ([]Recommendation, error)
	GetLikesCount(userId uint) (uint, error)
	CheckProStatus(userId uint) error
	GetStatus(userId uint) (string, error)
}
