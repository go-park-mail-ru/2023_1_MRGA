package recommendation

type IRepositoryRec interface {
	GetRecommendation(uint) ([]Recommendation, error)
	GetUserIdByToken(string) (uint, error)
}
