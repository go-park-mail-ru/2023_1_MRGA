package recommendation

type UseCase interface {
	GetRecommendations(userId uint) ([]Recommendation, error)
}
