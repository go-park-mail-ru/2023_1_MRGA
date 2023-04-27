package usecase

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo"
)

type InfoUseCase struct {
	userRepo     info_user.IRepositoryInfo
	infoUseCase  info.UseCase
	photoUseCase photo.UseCase
}

func NewInfoUseCase(userRepo info_user.IRepositoryInfo, infoUC info.UseCase, photoUC photo.UseCase) *InfoUseCase {
	return &InfoUseCase{
		userRepo:     userRepo,
		infoUseCase:  infoUC,
		photoUseCase: photoUC,
	}
}

func (iu *InfoUseCase) AddInfo(userId uint, info info_user.InfoStruct) error {
	userInfo := dataStruct.UserInfo{
		UserId:      userId,
		Name:        info.Name,
		Description: info.Description,
		Sex:         info.Sex,
	}

	cityId, err := iu.infoUseCase.GetCityId(info.City)
	if err != nil {
		return err
	}
	zodiacId, err := iu.infoUseCase.GetZodiacId(info.Zodiac)
	if err != nil {
		return err
	}
	educationId, err := iu.infoUseCase.GetEducationId(info.Education)
	if err != nil {
		return err
	}
	jobId, err := iu.infoUseCase.GetJobId(info.Job)
	if err != nil {
		return err
	}

	userInfo.Zodiac = zodiacId
	userInfo.CityId = cityId
	userInfo.Education = educationId
	userInfo.Job = jobId

	err = iu.userRepo.AddInfoUser(&userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (iu *InfoUseCase) GetInfo(userId uint) (userInfo info_user.InfoStructAnswer, err error) {
	userInfoTemp, err := iu.userRepo.GetUserInfo(userId)
	if err != nil {
		return
	}
	userInfo = info_user.InfoStructAnswer{
		Email:       userInfoTemp.Email,
		Name:        userInfoTemp.Name,
		Sex:         userInfoTemp.Sex,
		Job:         userInfoTemp.Job,
		Education:   userInfoTemp.Education,
		Zodiac:      userInfoTemp.Zodiac,
		City:        userInfoTemp.City,
		Description: userInfoTemp.Description,
	}
	age, err := iu.userRepo.GetAge(userId)
	if err != nil {
		return
	}
	userInfo.Age = age

	photos, err := iu.photoUseCase.GetAllPhotos(userId)
	if err != nil {
		return
	}
	userInfo.Photos = photos
	return
}

func (iu *InfoUseCase) ChangeInfo(userId uint, infoInp info_user.InfoChange) (info_user.InfoStructAnswer, error) {
	userInfo := dataStruct.UserInfo{
		Sex:         infoInp.Sex,
		Description: infoInp.Description,
		Name:        infoInp.Name,
		UserId:      userId,
	}

	cityId, err := iu.infoUseCase.GetCityId(infoInp.City)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.CityId = cityId

	zodiacId, err := iu.infoUseCase.GetZodiacId(infoInp.Zodiac)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Zodiac = zodiacId

	educationId, err := iu.infoUseCase.GetEducationId(infoInp.Education)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Education = educationId

	jobId, err := iu.infoUseCase.GetJobId(infoInp.Job)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Job = jobId

	err = iu.userRepo.ChangeInfo(&userInfo)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	result, err := iu.GetInfo(userId)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	return result, nil
}

func (iu *InfoUseCase) GetUserById(userId uint) (user info_user.UserRes, err error) {

	userTemp, err := iu.userRepo.GetUserById(userId)
	if err != nil {
		return
	}
	user.Name = userTemp.Name
	age, err := iu.userRepo.GetAge(userId)
	if err != nil {
		return
	}
	user.Age = age
	avatar, err := iu.photoUseCase.GetAvatar(userId)
	if err != nil {
		return
	}
	user.Avatar = avatar

	if user.Name == "" {
		user.Step = constform.NoInfo
		return
	}

	hashtags, err := iu.GetUserHashtags(userId)
	if err != nil {
		return
	}
	if len(hashtags) == 0 {
		user.Step = constform.NoHashtags
		return
	}

	reasons, err := iu.userRepo.CheckFilter(userId)
	if err != nil {
		return
	}
	if !reasons {
		user.Step = constform.NoFilters
		return
	}

	if user.Avatar == 0 {
		user.Step = constform.NoPhotos
		return
	}
	user.Step = constform.FullInfo
	return
}

func (iu *InfoUseCase) AddHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	hashtagId, err := iu.infoUseCase.GetHashtagId(hashtagInp.Hashtag)
	if err != nil {
		return err
	}
	var addHashtags []dataStruct.UserHashtag
	for _, hashtag := range hashtagId {
		var userHashtag dataStruct.UserHashtag
		userHashtag.HashtagId = hashtag
		userHashtag.UserId = userId
		addHashtags = append(addHashtags, userHashtag)
	}
	err = iu.userRepo.AddUserHashtag(addHashtags)
	if err != nil {
		return err
	}

	return nil
}

func (iu *InfoUseCase) GetUserHashtags(userId uint) ([]string, error) {
	hashtags, err := iu.userRepo.GetUserHashtags(userId)
	if err != nil {
		return nil, err
	}
	return hashtags, nil

}

func (iu *InfoUseCase) ChangeUserHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	hashtagsBD, err := iu.userRepo.GetUserHashtagsId(userId)
	if err != nil {
		return err
	}

	hashtagsId, err := iu.infoUseCase.GetHashtagId(hashtagInp.Hashtag)
	if err != nil {
		return err
	}

	var addHashtags []uint
	if err != nil {
		return err
	}
	for _, hashtag := range hashtagsId {
		if !Contains(hashtagsBD, hashtag) {
			addHashtags = append(addHashtags, hashtag)
		}
	}
	if len(addHashtags) > 0 {
		var hashtagsAdd []dataStruct.UserHashtag
		for _, hashtagId := range addHashtags {
			var hashtag dataStruct.UserHashtag
			hashtag.UserId = userId
			hashtag.HashtagId = hashtagId
			hashtagsAdd = append(hashtagsAdd, hashtag)
		}

		err = iu.userRepo.AddUserHashtag(hashtagsAdd)
		if err != nil {
			return err
		}
	}

	var deletedHashtags []uint
	for _, hashtag := range hashtagsBD {
		if !Contains(hashtagsId, hashtag) {
			deletedHashtags = append(deletedHashtags, hashtag)
		}
	}
	if len(deletedHashtags) > 0 {
		err = iu.userRepo.DeleteUserHashtag(userId, deletedHashtags)
		if err != nil {
			return err
		}
	}
	return nil
}

func (iu *InfoUseCase) GetUserHashtagsId(userId uint) ([]uint, error) {
	hashtags, err := iu.userRepo.GetUserHashtagsId(userId)
	if err != nil {
		return nil, err
	}
	return hashtags, nil
}

func Contains(s []uint, elem uint) bool {
	for _, elemS := range s {
		if elem == elemS {
			return true
		}
	}
	return false
}
