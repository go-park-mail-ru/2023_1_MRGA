package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/mocks"
)

func TestNewInfoUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	if infoUsecase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestInfoUseCase_GetCities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	cities := []string{
		"Москва",
		"London",
	}
	infoRepoMock.EXPECT().GetCities().Return(cities, nil)

	citiesRes, err := infoUsecase.GetCities()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, citiesRes, cities)
}

func TestInfoUseCase_GetCities_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetCities().Return(nil, errRepo)

	_, err := infoUsecase.GetCities()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetJobs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	jobs := []string{
		"plumber",
		"worker",
	}
	infoRepoMock.EXPECT().GetJobs().Return(jobs, nil)

	jobsRes, err := infoUsecase.GetJobs()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, jobsRes, jobs)
}

func TestInfoUseCase_GetJobs_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetJobs().Return(nil, errRepo)

	_, err := infoUsecase.GetJobs()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetEducation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	education := []string{
		"university",
		"school",
	}
	infoRepoMock.EXPECT().GetEducation().Return(education, nil)

	educationRes, err := infoUsecase.GetEducation()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, educationRes, education)
}

func TestInfoUseCase_GetEducation_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetEducation().Return(nil, errRepo)

	_, err := infoUsecase.GetEducation()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetZodiacs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	zodiac := []string{
		"стрелец",
		"овен",
	}
	infoRepoMock.EXPECT().GetZodiac().Return(zodiac, nil)

	zodiacRes, err := infoUsecase.GetZodiacs()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, zodiacRes, zodiac)
}

func TestInfoUseCase_GetZodiacs_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetZodiac().Return(nil, errRepo)

	_, err := infoUsecase.GetZodiacs()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	hashtags := []string{
		"sport",
		"walk",
	}
	infoRepoMock.EXPECT().GetHashtags().Return(hashtags, nil)

	hashtagsRes, err := infoUsecase.GetHashtags()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, hashtagsRes, hashtags)
}

func TestInfoUseCase_GetHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetHashtags().Return(nil, errRepo)

	_, err := infoUsecase.GetHashtags()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetStatuses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	status := []string{
		"pro",
		"basic",
	}
	infoRepoMock.EXPECT().GetStatuses().Return(status, nil)

	statusRes, err := infoUsecase.GetStatuses()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, statusRes, status)
}

func TestInfoUseCase_GetStatuses_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetStatuses().Return(nil, errRepo)

	_, err := infoUsecase.GetStatuses()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetReasons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	reasons := []string{
		"test1",
		"test2",
	}
	infoRepoMock.EXPECT().GetReasons().Return(reasons, nil)

	reasonRes, err := infoUsecase.GetReasons()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, reasonRes, reasons)
}

func TestInfoUseCase_GetReasons_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetReasons().Return(nil, errRepo)

	_, err := infoUsecase.GetReasons()
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetCityId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := "test1"
	testRes := uint(1)
	infoRepoMock.EXPECT().GetCityId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetCityId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetCityId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := "test1"
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetCityId(testInp).Return(uint(0), errRepo)

	_, err := infoUsecase.GetCityId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetZodiacId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := "test1"
	testRes := uint(1)
	infoRepoMock.EXPECT().GetZodiacId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetZodiacId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetZodiacId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := "test1"
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetZodiacId(testInp).Return(uint(0), errRepo)

	_, err := infoUsecase.GetZodiacId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetEducationId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := "test1"
	testRes := uint(1)
	infoRepoMock.EXPECT().GetEducationId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetEducationId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetEducationId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := "test1"
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetJobId(testInp).Return(uint(0), errRepo)

	_, err := infoUsecase.GetJobId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetJobId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := "test1"
	testRes := uint(1)
	infoRepoMock.EXPECT().GetJobId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetJobId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetJobId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := "test1"
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetJobId(testInp).Return(uint(0), errRepo)

	_, err := infoUsecase.GetJobId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetStatusId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := "test1"
	testRes := uint(1)
	infoRepoMock.EXPECT().GetStatusId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetStatusId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetStatusId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := "test1"
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetStatusId(testInp).Return(uint(0), errRepo)

	_, err := infoUsecase.GetStatusId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetHashtagId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)

	testInp := []string{"test1", "test2"}
	testRes := []uint{1, 2}
	infoRepoMock.EXPECT().GetHashtagId(testInp).Return(testRes, nil)

	result, err := infoUsecase.GetHashtagId(testInp)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, testRes)
}

func TestInfoUseCase_GetHashtagId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	infoUsecase := NewInfoUseCase(infoRepoMock)
	testInp := []string{"test1", "test2"}
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetHashtagId(testInp).Return(nil, errRepo)

	_, err := infoUsecase.GetHashtagId(testInp)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}
