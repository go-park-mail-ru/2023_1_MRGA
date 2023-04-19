package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
)

type InfoUseCase struct {
	userRepo info_user.IRepositoryInfo
	infoRepo info.IRepositoryInfo
}

func NewInfoUseCase(userRepo info_user.IRepositoryInfo, infoRepo info.IRepositoryInfo) *InfoUseCase {
	return &InfoUseCase{
		userRepo: userRepo,
		infoRepo: infoRepo,
	}
}

func (iu *InfoUseCase) AddInfo(userId uint, info info_user.InfoStruct) error {
	userInfo := dataStruct.UserInfo{
		UserId:      userId,
		Name:        info.Name,
		Description: info.Description,
		Sex:         info.Sex,
	}

	cityId, err := iu.infoRepo.GetCityId(info.City)
	if err != nil {
		return err
	}
	zodiacId, err := iu.infoRepo.GetZodiacId(info.Zodiac)
	if err != nil {
		return err
	}
	educationId, err := iu.infoRepo.GetEducationId(info.Education)
	if err != nil {
		return err
	}
	jobId, err := iu.infoRepo.GetJobId(info.Job)
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
} //ok

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
	avatar, err := iu.userRepo.GetAvatar(userId) //fix
	if err != nil {
		return
	}
	userInfo.Photos = append(userInfo.Photos, avatar) //fix
	photos, err := iu.userRepo.GetUserPhoto(userId)
	if err != nil {
		return
	}
	userInfo.Photos = append(userInfo.Photos, photos...)
	return
} //fix

func (iu *InfoUseCase) ChangeInfo(userId uint, infoInp info_user.InfoChange) (info_user.InfoStructAnswer, error) {
	userInfo := dataStruct.UserInfo{
		Sex:         infoInp.Sex,
		Description: infoInp.Description,
		Name:        infoInp.Name,
		UserId:      userId,
	}

	cityId, err := iu.infoRepo.GetCityId(infoInp.City)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.CityId = cityId

	zodiacId, err := iu.infoRepo.GetZodiacId(infoInp.Zodiac)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Zodiac = zodiacId

	educationId, err := iu.infoRepo.GetEducationId(infoInp.Education)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Education = educationId

	jobId, err := iu.infoRepo.GetJobId(infoInp.Job)
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
} //ok

func (iu *InfoUseCase) AddHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	hashtagId, err := iu.infoRepo.GetHashtagId(hashtagInp.Hashtag)
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

func (iu *InfoUseCase) GetUserHashtags(userId uint) (info_user.HashtagInp, error) {
	hashtags, err := iu.userRepo.GetUserHashtags(userId)
	if err != nil {
		return info_user.HashtagInp{}, err
	}

	var result info_user.HashtagInp
	result.Hashtag, err = iu.userRepo.GetHashtagById(hashtags)
	if err != nil {
		return info_user.HashtagInp{}, err
	}

	return result, nil
}

func (iu *InfoUseCase) ChangeUserHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	hashtagsBD, err := iu.userRepo.GetUserHashtags(userId)
	if err != nil {
		return err
	}

	hashtagsId, err := iu.infoRepo.GetHashtagId(hashtagInp.Hashtag)
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

func Contains(s []uint, elem uint) bool {
	for _, elemS := range s {
		if elem == elemS {
			return true
		}
	}
	return false
}
