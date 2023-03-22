package recommendation

type IRepositoryAuth interface {
	GetRecommendation(uint) ([]Recommendation, error)
	GetUserIdByToken(string) (uint, error)
}
