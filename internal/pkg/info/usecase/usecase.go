package usecase

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info"

type InfoUseCase struct {
	userRepo info.IRepositoryInfo
}

func NewInfoUseCase(userRepo info.IRepositoryInfo) *InfoUseCase {
	return &InfoUseCase{
		userRepo: userRepo,
	}
}

func (iu *InfoUseCase) GetHashtags() ([]string, error) {
	hashtags, err := iu.userRepo.GetHashtags()
	if err != nil {
		return nil, err
	}

	return hashtags, nil
}

func (iu *InfoUseCase) GetCities() ([]string, error) {
	cities, err := iu.userRepo.GetCities()
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func (iu *InfoUseCase) GetJobs() ([]string, error) {
	jobs, err := iu.userRepo.GetJobs()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (iu *InfoUseCase) GetEducation() ([]string, error) {
	education, err := iu.userRepo.GetEducation()
	if err != nil {
		return nil, err
	}

	return education, nil
}

func (iu *InfoUseCase) GetZodiacs() ([]string, error) {
	zodiac, err := iu.userRepo.GetZodiac()
	if err != nil {
		return nil, err
	}

	return zodiac, nil
}

func (iu *InfoUseCase) GetReasons() ([]string, error) {
	reasons, err := iu.userRepo.GetReasons()
	if err != nil {
		return nil, err
	}

	return reasons, nil
}

func (iu *InfoUseCase) GetCityId(city string) (uint, error) {
	cityId, err := iu.userRepo.GetCityId(city)
	if err != nil {
		return 0, err
	}
	return cityId, nil
}

func (iu *InfoUseCase) GetZodiacId(zodiac string) (uint, error) {
	zodiacId, err := iu.userRepo.GetZodiacId(zodiac)
	if err != nil {
		return 0, err
	}
	return zodiacId, nil
}

func (iu *InfoUseCase) GetEducationId(education string) (uint, error) {
	educationId, err := iu.userRepo.GetEducationId(education)
	if err != nil {
		return 0, err
	}
	return educationId, nil
}

func (iu *InfoUseCase) GetJobId(job string) (uint, error) {
	jobId, err := iu.userRepo.GetJobId(job)
	if err != nil {
		return 0, err
	}
	return jobId, nil
}

func (iu *InfoUseCase) GetHashtagId(hashtagId []string) ([]uint, error) {
	hashtags, err := iu.userRepo.GetHashtagId(hashtagId)
	if err != nil {
		return nil, err
	}
	return hashtags, nil
}
