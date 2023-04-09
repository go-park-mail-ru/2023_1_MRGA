package recommendation

type UseCase interface {
	GetRecommendations(userId uint) ([]Recommendation, error)
	AddFilters(userId uint, FilterInp FilterInput) error
	GetReasons() ([]string, error)
	GetFilters(userId uint) (FilterInput, error)
	ChangeFilters(userId uint, filterInp FilterInput) error
}
