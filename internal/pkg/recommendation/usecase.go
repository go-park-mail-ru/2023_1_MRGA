package recommendation

type UseCase interface {
	GetRecommendation(string) ([]Recommendation, error)
}
