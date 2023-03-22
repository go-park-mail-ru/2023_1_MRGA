package usecase

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type RecUseCase struct {
	userRepo recommendation.IRepositoryAuth
}

func NewRecUseCase(
	userRepo recommendation.IRepositoryAuth) *RecUseCase {
	return &RecUseCase{
		userRepo: userRepo,
	}
}

func (r *RecUseCase) GetRecommendation(token string) ([]recommendation.Recommendation, error) {

	userId, err := r.userRepo.GetUserIdByToken(token)
	if err != nil {
		return nil, err
	}

	recs, err := r.userRepo.GetRecommendation(userId)
	if err != nil {
		return nil, err
	}

	return recs, err
}
