package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	infoMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/mocks"
	photoMock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/mocks"
)

func TestNewInfoUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	if infoUserUseCase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestInfoUseCase_GetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	userAge := 20
	photos := []uint{1, 2}
	user := info_user.InfoStruct{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userRes := info_user.InfoStructAnswer{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
		Age:         userAge,
		Photos:      photos,
	}
	infoRepoMock.EXPECT().GetUserInfo(userId).Return(user, nil)
	infoRepoMock.EXPECT().GetAge(userId).Return(userAge, nil)
	photoUseCase.EXPECT().GetAllPhotos(userId).Return(photos, nil)

	result, err := infoUserUseCase.GetInfo(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, userRes)
}

func TestInfoUseCase_GetInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	infoRepoMock.EXPECT().GetUserInfo(userId).Return(info_user.InfoStruct{}, errRepo)

	_, err := infoUserUseCase.GetInfo(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_AddInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	user := info_user.InfoStruct{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userRes := dataStruct.UserInfo{
		UserId:      userId,
		Name:        "test",
		CityId:      uint(1),
		Sex:         uint(1),
		Description: "test",
		Zodiac:      uint(1),
		Job:         uint(1),
		Education:   uint(1),
	}

	infoUseCase.EXPECT().GetCityId(user.City).Return(uint(1), nil)
	infoUseCase.EXPECT().GetEducationId(user.Education).Return(uint(1), nil)
	infoUseCase.EXPECT().GetZodiacId(user.Zodiac).Return(uint(1), nil)
	infoUseCase.EXPECT().GetJobId(user.Job).Return(uint(1), nil)
	infoRepoMock.EXPECT().AddInfoUser(&userRes).Return(nil)

	err := infoUserUseCase.AddInfo(userId, user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestInfoUseCase_AddInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	user := info_user.InfoStruct{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userRes := dataStruct.UserInfo{
		UserId:      userId,
		Name:        "test",
		CityId:      uint(1),
		Sex:         uint(1),
		Description: "test",
		Zodiac:      uint(1),
		Job:         uint(1),
		Education:   uint(1),
	}
	errRepo := fmt.Errorf("something wrong")

	infoUseCase.EXPECT().GetCityId(user.City).Return(uint(1), nil)
	infoUseCase.EXPECT().GetEducationId(user.Education).Return(uint(1), nil)
	infoUseCase.EXPECT().GetZodiacId(user.Zodiac).Return(uint(1), nil)
	infoUseCase.EXPECT().GetJobId(user.Job).Return(uint(1), nil)
	infoRepoMock.EXPECT().AddInfoUser(&userRes).Return(errRepo)

	err := infoUserUseCase.AddInfo(userId, user)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_ChangeInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userAge := 20
	photos := []uint{1, 2}
	userGet := info_user.InfoStruct{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userResGet := info_user.InfoStructAnswer{
		Name:        "test",
		City:        "test",
		Email:       "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
		Age:         userAge,
		Photos:      photos,
	}

	userId := uint(1)
	user := info_user.InfoChange{
		Name:        "test",
		City:        "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userRes := dataStruct.UserInfo{
		UserId:      userId,
		Name:        "test",
		CityId:      uint(1),
		Sex:         uint(1),
		Description: "test",
		Zodiac:      uint(1),
		Job:         uint(1),
		Education:   uint(1),
	}

	infoUseCase.EXPECT().GetCityId(user.City).Return(uint(1), nil)
	infoUseCase.EXPECT().GetEducationId(user.Education).Return(uint(1), nil)
	infoUseCase.EXPECT().GetZodiacId(user.Zodiac).Return(uint(1), nil)
	infoUseCase.EXPECT().GetJobId(user.Job).Return(uint(1), nil)

	infoRepoMock.EXPECT().ChangeInfo(&userRes).Return(nil)

	infoRepoMock.EXPECT().GetUserInfo(userId).Return(userGet, nil)
	infoRepoMock.EXPECT().GetAge(userId).Return(userAge, nil)
	photoUseCase.EXPECT().GetAllPhotos(userId).Return(photos, nil)

	result, err := infoUserUseCase.ChangeInfo(userId, user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	require.EqualValues(t, result, userResGet)
}

func TestInfoUseCase_ChangeInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	user := info_user.InfoChange{
		Name:        "test",
		City:        "test",
		Sex:         uint(1),
		Description: "test",
		Zodiac:      "test",
		Job:         "test",
		Education:   "test",
	}
	userRes := dataStruct.UserInfo{
		UserId:      userId,
		Name:        "test",
		CityId:      uint(1),
		Sex:         uint(1),
		Description: "test",
		Zodiac:      uint(1),
		Job:         uint(1),
		Education:   uint(1),
	}

	infoUseCase.EXPECT().GetCityId(user.City).Return(uint(1), nil)
	infoUseCase.EXPECT().GetEducationId(user.Education).Return(uint(1), nil)
	infoUseCase.EXPECT().GetZodiacId(user.Zodiac).Return(uint(1), nil)
	infoUseCase.EXPECT().GetJobId(user.Job).Return(uint(1), nil)

	infoRepoMock.EXPECT().ChangeInfo(&userRes).Return(errRepo)

	_, err := infoUserUseCase.ChangeInfo(userId, user)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	userAge := 20
	avatar := uint(1)
	hashtags := []string{
		"test1",
		"test2",
	}
	user := info_user.UserRestTemp{
		Name: "test",
	}

	userRes := info_user.UserRes{
		Name:   "test",
		Age:    userAge,
		Avatar: avatar,
		Step:   constform.FullInfo,
		Banned: false,
	}
	infoRepoMock.EXPECT().GetUserById(userId).Return(user, nil)
	infoRepoMock.EXPECT().GetAge(userId).Return(userAge, nil)
	photoUseCase.EXPECT().GetAvatar(userId).Return(avatar, nil)
	infoRepoMock.EXPECT().GetUserHashtags(userId).Return(hashtags, nil)
	infoRepoMock.EXPECT().CheckFilter(userId).Return(true, nil)

	result, err := infoUserUseCase.GetUserById(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, userRes)
}

func TestInfoUseCase_GetUserById_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	userId := uint(1)
	user := info_user.UserRestTemp{
		Name: "test",
	}
	errRepo := fmt.Errorf("something wrong")

	infoRepoMock.EXPECT().GetUserById(userId).Return(user, errRepo)

	_, err := infoUserUseCase.GetUserById(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	hashtags := []string{
		"test1",
		"test2",
	}
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserHashtags(userId).Return(hashtags, nil)
	result, err := infoUserUseCase.GetUserHashtags(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, hashtags)
}

func TestInfoUseCase_GetUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserHashtags(userId).Return(nil, errRepo)
	_, err := infoUserUseCase.GetUserHashtags(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_AddUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	hashtags := info_user.HashtagInp{
		Hashtag: []string{
			"test1",
			"test2"},
	}
	hashtagsId := []uint{
		1, 2,
	}
	userId := uint(1)
	hashtagsAdd := []dataStruct.UserHashtag{
		{
			UserId:    userId,
			HashtagId: hashtagsId[0],
		},
		{
			UserId:    userId,
			HashtagId: hashtagsId[1],
		},
	}

	infoUseCase.EXPECT().GetHashtagId(hashtags.Hashtag).Return(hashtagsId, nil)
	infoRepoMock.EXPECT().AddUserHashtag(hashtagsAdd).Return(nil)
	err := infoUserUseCase.AddHashtags(userId, hashtags)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestInfoUseCase_AddUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")

	hashtags := info_user.HashtagInp{
		Hashtag: []string{
			"test1",
			"test2"},
	}
	hashtagsId := []uint{
		1, 2,
	}
	userId := uint(1)
	hashtagsAdd := []dataStruct.UserHashtag{
		{
			UserId:    userId,
			HashtagId: hashtagsId[0],
		},
		{
			UserId:    userId,
			HashtagId: hashtagsId[1],
		},
	}

	infoUseCase.EXPECT().GetHashtagId(hashtags.Hashtag).Return(hashtagsId, nil)
	infoRepoMock.EXPECT().AddUserHashtag(hashtagsAdd).Return(errRepo)
	err := infoUserUseCase.AddHashtags(userId, hashtags)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_ChangeUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	hashtags := info_user.HashtagInp{
		Hashtag: []string{
			"test1",
			"test2"},
	}
	hashtagsId := []uint{
		2, 3,
	}
	hashtagsId2 := []uint{
		1, 2,
	}
	userId := uint(1)
	hashtagsAdd := []dataStruct.UserHashtag{
		{
			UserId:    userId,
			HashtagId: 1,
		},
	}

	infoUseCase.EXPECT().GetHashtagId(hashtags.Hashtag).Return(hashtagsId2, nil)
	infoRepoMock.EXPECT().GetUserHashtagsId(userId).Return(hashtagsId, nil)
	infoRepoMock.EXPECT().AddUserHashtag(hashtagsAdd).Return(nil)
	infoRepoMock.EXPECT().DeleteUserHashtag(userId, []uint{3}).Return(nil)
	err := infoUserUseCase.ChangeUserHashtags(userId, hashtags)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestInfoUseCase_ChangeUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")

	hashtags := info_user.HashtagInp{
		Hashtag: []string{
			"test1",
			"test2"},
	}
	hashtagsId := []uint{
		2, 3,
	}
	hashtagsId2 := []uint{
		1, 2,
	}
	userId := uint(1)
	hashtagsAdd := []dataStruct.UserHashtag{
		{
			UserId:    userId,
			HashtagId: 1,
		},
	}

	infoUseCase.EXPECT().GetHashtagId(hashtags.Hashtag).Return(hashtagsId2, nil)
	infoRepoMock.EXPECT().GetUserHashtagsId(userId).Return(hashtagsId, nil)
	infoRepoMock.EXPECT().AddUserHashtag(hashtagsAdd).Return(nil)
	infoRepoMock.EXPECT().DeleteUserHashtag(userId, []uint{3}).Return(errRepo)
	err := infoUserUseCase.ChangeUserHashtags(userId, hashtags)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetUserHashtagsId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	hashtags := []uint{
		1,
		2,
	}
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserHashtagsId(userId).Return(hashtags, nil)
	result, err := infoUserUseCase.GetUserHashtagsId(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, hashtags)
}

func TestInfoUseCase_GetUserHashtagsId_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserHashtagsId(userId).Return(nil, errRepo)
	_, err := infoUserUseCase.GetUserHashtagsId(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_GetUserStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	hashtags := "test1"
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserStatus(userId).Return(hashtags, nil)
	result, err := infoUserUseCase.GetUserStatus(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, result, hashtags)
}

func TestInfoUseCase_GetUserStatus_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	infoRepoMock.EXPECT().GetUserStatus(userId).Return("", errRepo)
	_, err := infoUserUseCase.GetUserStatus(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}

func TestInfoUseCase_ChangeUserStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	newStatus := info_user.StatusInp{Status: "test2"}
	statusId := uint(1)
	userId := uint(1)
	infoUseCase.EXPECT().GetStatusId(newStatus.Status).Return(statusId, nil)
	infoRepoMock.EXPECT().ChangeUserStatus(userId, statusId).Return(nil)

	err := infoUserUseCase.ChangeUserStatus(userId, newStatus)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestInfoUseCase_ChangeUserStatus_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoRepoMock := mock.NewMockIRepositoryInfo(ctrl)
	photoUseCase := photoMock.NewMockUseCase(ctrl)
	infoUseCase := infoMock.NewMockUseCase(ctrl)
	infoUserUseCase := NewInfoUseCase(infoRepoMock, infoUseCase, photoUseCase)

	errRepo := fmt.Errorf("something wrong")

	newStatus := info_user.StatusInp{Status: "test2"}
	statusId := uint(1)
	userId := uint(1)
	infoUseCase.EXPECT().GetStatusId(newStatus.Status).Return(statusId, nil)
	infoRepoMock.EXPECT().ChangeUserStatus(userId, statusId).Return(errRepo)

	err := infoUserUseCase.ChangeUserStatus(userId, newStatus)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, errRepo.Error())
}
