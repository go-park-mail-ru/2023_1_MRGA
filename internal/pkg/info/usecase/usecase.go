package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"
)

type InfoUseCase struct {
	userRepo info.IRepositoryInfo
}

func NewInfoUseCase(userRepo info.IRepositoryInfo) *InfoUseCase {
	return &InfoUseCase{
		userRepo: userRepo,
	}
}

func (iu *InfoUseCase) AddInfo(userId uint, info info.InfoStruct) error {
	var userInfo dataStruct.UserInfo
	var avatar dataStruct.UserPhoto
	avatar.Avatar = true
	avatar.UserId = userId
	avatar.Photo = info.Avatar
	err := iu.userRepo.AddUserPhoto(&avatar)
	if err != nil {
		return err
	}
	for _, photo := range info.Photo {
		var photoDB dataStruct.UserPhoto
		photoDB.Avatar = false
		photoDB.UserId = userId
		photoDB.Photo = photo
		err = iu.userRepo.AddUserPhoto(&photoDB)
		if err != nil {
			return err
		}
	}

	userInfo.UserId = userId
	userInfo.Name = info.Name
	userInfo.Description = info.Description
	userInfo.Sex = info.Sex

	cityId, err := iu.userRepo.GetCityId(info.City)
	if err != nil {
		return err
	}
	zodiacId, err := iu.userRepo.GetZodiacId(info.Zodiac)
	if err != nil {
		return err
	}
	educationId, err := iu.userRepo.GetEducationId(info.Education)
	if err != nil {
		return err
	}
	jobId, err := iu.userRepo.GetJobId(info.Job)
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

func (iu *InfoUseCase) GetInfo(userId uint) (userInfo info.InfoStruct, err error) {
	userInfo, err = iu.userRepo.GetUserInfo(userId)
	if err != nil {
		return
	}

	avatar, err := iu.userRepo.GetAvatar(userId)
	if err != nil {
		return
	}
	userInfo.Avatar = avatar

	photos, err := iu.userRepo.GetPhotos(userId)
	for _, photo := range photos {
		userInfo.Photo = append(userInfo.Photo, photo.Photo)
	}
	return
}

func (iu *InfoUseCase) GetCities() ([]string, error) {
	cities, err := iu.userRepo.GetCities()
	if err != nil {
		return nil, err
	}

	var citiesResult []string
	for _, city := range cities {
		citiesResult = append(citiesResult, city.City)
	}

	return citiesResult, nil
}

func (iu *InfoUseCase) GetJobs() ([]string, error) {
	jobs, err := iu.userRepo.GetJobs()
	if err != nil {
		return nil, err
	}

	var jobsResult []string
	for _, job := range jobs {
		jobsResult = append(jobsResult, job.Job)
	}

	return jobsResult, nil
}

func (iu *InfoUseCase) GetEducation() ([]string, error) {
	education, err := iu.userRepo.GetEducation()
	if err != nil {
		return nil, err
	}

	var educationResult []string
	for _, ed := range education {
		educationResult = append(educationResult, ed.Education)
	}

	return educationResult, nil
}

func (iu *InfoUseCase) GetZodiacs() ([]string, error) {
	zodiac, err := iu.userRepo.GetZodiac()
	if err != nil {
		return nil, err
	}

	var zodiacResult []string
	for _, z := range zodiac {
		zodiacResult = append(zodiacResult, z.Zodiac)
	}

	return zodiacResult, nil
}
