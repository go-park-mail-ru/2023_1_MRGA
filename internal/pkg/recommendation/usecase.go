package recommendation

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	GetRecommendations(userId uint) ([]Recommendation, error)
}
