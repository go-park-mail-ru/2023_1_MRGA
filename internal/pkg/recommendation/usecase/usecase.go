package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
)

type RecUseCase struct {
	userRepo recommendation.IRepositoryRec
}

func NewRecUseCase(
	userRepo recommendation.IRepositoryRec) *RecUseCase {
	return &RecUseCase{
		userRepo: userRepo,
	}
}

func (r *RecUseCase) GetRecommendations(userId uint) ([]recommendation.Recommendation, error) {
	hashtags, err := r.userRepo.GetUserHashtags(userId)
	if err != nil {
		return nil, err
	}

	reasons, err := r.userRepo.GetUserReasons(userId)
	if err != nil {
		return nil, err
	}

	filters, err := r.userRepo.GetFilter(userId)
	if err != nil {
		return nil, err
	}

	hashtagsSlice := make([]uint, 0)
	for _, hashtagId := range hashtags {
		hashtagsSlice = append(hashtagsSlice, hashtagId.HashtagId)
	}

	reasonsSlice := make([]uint, 0)
	for _, reasonId := range reasons {
		reasonsSlice = append(reasonsSlice, reasonId.ReasonId)
	}

	history, err := r.userRepo.GetUserHistory(userId)
	historySlice := []uint{0}
	for _, historyId := range history {
		historySlice = append(historySlice, historyId.UserProfileId)
	}

	recs, err := r.userRepo.GetRecommendation(userId, historySlice, reasonsSlice, hashtagsSlice, filters)
	if err != nil {
		return nil, err
	}

	var result []recommendation.Recommendation
	for _, rec := range recs {
		user, err := r.userRepo.GetRecommendedUser(rec.UserId)
		if err != nil {
			return nil, err
		}

		hashtagsUser, err := r.userRepo.GetUserNameHashtags(rec.UserId)
		if err != nil {
			return nil, err
		}
		user.Hashtags = hashtagsUser
		avatar, err := r.userRepo.GetAvatar(rec.UserId)

		user.Photos = append(user.Photos, avatar)
		photos, err := r.userRepo.GetPhotos(rec.UserId)
		if err != nil {
			return nil, err
		}

		user.Photos = append(user.Photos, photos...)

		result = append(result, user)
	}

	return result, err
}

func (r *RecUseCase) AddFilters(userId uint, filterInp recommendation.FilterInput) error {
	for _, reason := range filterInp.Reason {
		reasonId, err := r.userRepo.GetReasonId(reason)
		if err != nil {
			return err
		}

		var userReason dataStruct.UserReason
		userReason.UserId = userId
		userReason.ReasonId = reasonId
		err = r.userRepo.AddUserReason(&userReason)
		if err != nil {
			return err
		}
	}

	var filter dataStruct.UserFilter
	filter.UserId = userId
	filter.SearchSex = filterInp.SearchSex
	filter.MinAge = filterInp.MinAge
	filter.MaxAge = filterInp.MaxAge

	err := r.userRepo.AddFilter(&filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecUseCase) GetReasons() ([]string, error) {
	reasons, err := r.userRepo.GetReasons()
	if err != nil {
		return nil, err
	}

	var reasonsResult []string
	for _, reason := range reasons {
		reasonsResult = append(reasonsResult, reason.Reason)
	}

	return reasonsResult, nil
}

func (r *RecUseCase) GetFilters(userId uint) (recommendation.FilterInput, error) {
	var filterRes recommendation.FilterInput
	filter, err := r.userRepo.GetFilter(userId)
	if err != nil {
		return recommendation.FilterInput{}, err
	}

	filterRes.MinAge = filter.MinAge
	filterRes.MaxAge = filter.MaxAge
	filterRes.SearchSex = filter.SearchSex

	reasons, err := r.userRepo.GetUserReasons(userId)
	if err != nil {
		return recommendation.FilterInput{}, err
	}

	for _, reasonId := range reasons {
		reason, err := r.userRepo.GetReasonById(reasonId.ReasonId)
		if err != nil {
			return recommendation.FilterInput{}, err
		}
		filterRes.Reason = append(filterRes.Reason, reason)
	}

	return filterRes, nil
}

func (r *RecUseCase) ChangeFilters(userId uint, filterInp recommendation.FilterInput) error {
	reasonsBD, err := r.userRepo.GetUserReasons(userId)
	if err != nil {
		return err
	}
	var reasonsSlice []string
	for _, reasonId := range reasonsBD {
		reason, err := r.userRepo.GetReasonById(reasonId.ReasonId)
		if err != nil {
			return err
		}
		reasonsSlice = append(reasonsSlice, reason)
	}

	for _, reason := range filterInp.Reason {
		if !Contains(reasonsSlice, reason) {
			var reasonAdd dataStruct.UserReason
			reasonId, err := r.userRepo.GetReasonId(reason)
			if err != nil {
				return err
			}
			reasonAdd.UserId = userId
			reasonAdd.ReasonId = reasonId
			err = r.userRepo.AddUserReason(&reasonAdd)
			if err != nil {
				return err
			}
		}
	}

	for _, reason := range reasonsSlice {
		if !Contains(filterInp.Reason, reason) {
			reasonId, err := r.userRepo.GetReasonId(reason)
			if err != nil {
				return err
			}
			err = r.userRepo.DeleteUserReason(userId, reasonId)
			if err != nil {
				return err
			}
		}
	}

	var newFilter dataStruct.UserFilter
	newFilter.UserId = userId
	newFilter.MaxAge = filterInp.MaxAge
	newFilter.SearchSex = filterInp.SearchSex

	err = r.userRepo.ChangeFilter(newFilter)
	if err != nil {
		return err
	}
	return nil
}

func Contains(s []string, elem string) bool {
	for _, elemS := range s {
		if elem == elemS {
			return true
		}
	}
	return false
}
