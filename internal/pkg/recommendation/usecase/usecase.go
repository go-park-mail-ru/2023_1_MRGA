package usecase

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type RecUseCase struct {
	repo            recommendation.IRepositoryRec
	filterUseCase   filter.UseCase
	photoUseCase    photo.UseCase
	infoUserUseCase info_user.UseCase
}

func NewRecUseCase(userRepo recommendation.IRepositoryRec, filterUc filter.UseCase, photoUC photo.UseCase, infoUserUC info_user.UseCase) *RecUseCase {
	return &RecUseCase{
		repo:            userRepo,
		filterUseCase:   filterUc,
		photoUseCase:    photoUC,
		infoUserUseCase: infoUserUC,
	}
}

func (r *RecUseCase) GetRecommendations(userId uint) ([]recommendation.Recommendation, error) {
	hashtags, err := r.infoUserUseCase.GetUserHashtagsId(userId)
	if err != nil {
		return nil, err
	}

	reasons, err := r.filterUseCase.GetUserReasonsId(userId)
	if err != nil {
		return nil, err
	}

	filters, err := r.filterUseCase.GetUserFilters(userId)
	if err != nil {
		return nil, err
	}

	history, err := r.repo.GetUserHistory(userId)
	if err != nil {
		return nil, err
	}
	if len(history) == 0 {
		history = append(history, 0)
	}

	recs, err := r.repo.GetRecommendation(userId, history, reasons, hashtags, filters)
	if err != nil {
		return nil, err
	}

	var result []recommendation.Recommendation
	for _, rec := range recs {
		user, err := r.repo.GetRecommendedUser(rec.UserId)
		if err != nil {
			return nil, err
		}

		hashtagsUser, err := r.infoUserUseCase.GetUserHashtags(rec.UserId)
		if err != nil {
			return nil, err
		}
		user.Hashtags = hashtagsUser

		user.Photos, err = r.photoUseCase.GetAllPhotos(rec.UserId)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	return result, err
}

func (r *RecUseCase) CheckProStatus(userId uint) error {
	err := r.repo.CheckStatus(userId)
	return err
}

func (r *RecUseCase) GetStatus(userId uint) (string, error) {
	status, err := r.repo.GetStatus(userId)
	return status, err
}

func (r *RecUseCase) GetLikes(userId uint) ([]recommendation.Recommendation, error) {
	filters, err := r.filterUseCase.GetUserFilters(userId)
	if err != nil {
		return nil, err
	}

	reasons, err := r.filterUseCase.GetUserReasonsId(userId)
	if err != nil {
		return nil, err
	}

	history, err := r.repo.GetUserHistory(userId)
	if err != nil {
		return nil, err
	}
	if len(history) == 0 {
		history = append(history, 0)
	}

	recs, err := r.repo.GetLikes(userId, history, reasons, filters)
	if err != nil {
		return nil, err
	}

	var result []recommendation.Recommendation
	for _, rec := range recs {
		user, err := r.repo.GetRecommendedUser(rec.UserId)
		if err != nil {
			return nil, err
		}

		hashtagsUser, err := r.infoUserUseCase.GetUserHashtags(rec.UserId)
		if err != nil {
			return nil, err
		}
		user.Hashtags = hashtagsUser

		user.Photos, err = r.photoUseCase.GetAllPhotos(rec.UserId)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	return result, err
}

func (r *RecUseCase) GetLikesCount(userId uint) (count uint, err error) {
	filters, err := r.filterUseCase.GetUserFilters(userId)
	if err != nil {
		return
	}

	reasons, err := r.filterUseCase.GetUserReasonsId(userId)
	if err != nil {
		return
	}

	history, err := r.repo.GetUserHistory(userId)
	if err != nil {
		return
	}
	if len(history) == 0 {
		history = append(history, 0)
	}

	count, err = r.repo.GetLikesCount(userId, history, reasons, filters)
	if err != nil {
		return
	}

	return
}
