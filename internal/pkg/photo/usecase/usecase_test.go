package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/photo/mocks"
)

func TestNewPhotoUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	if photoUsecase == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestPhotoUseCase_GetAllPhotos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := []uint{1, 2, 3}

	userId := uint(1)

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{2, 3}, nil)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)

	photoResult, err := photoUsecase.GetAllPhotos(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, photoResult, photos)
}

func TestPhotoUseCase_GetAllPhotos_GetPhotoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{}, errRepo)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)

	_, err := photoUsecase.GetAllPhotos(userId)
	require.EqualError(t, err, errRepo.Error())
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestPhotoUseCase_GetAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := uint(1)

	userId := uint(1)

	photoRepoMock.EXPECT().GetAvatar(userId).Return(photos, nil)

	photoResult, err := photoUsecase.GetAvatar(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualValues(t, photoResult, photos)
}

func TestPhotoUseCase_GetAvatar_GetAvatarError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)
	errRepo := fmt.Errorf("something wrong")

	userId := uint(1)

	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(0), errRepo)

	_, err := photoUsecase.GetAvatar(userId)
	require.EqualError(t, err, errRepo.Error())
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestPhotoUseCase_SavePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photoId := uint(1)
	avatar := false
	userId := uint(1)

	photoRow := dataStruct.UserPhoto{
		Photo:  photoId,
		Avatar: avatar,
		UserId: userId,
	}

	photoRepoMock.EXPECT().SavePhoto(photoRow).Return(nil)

	err := photoUsecase.SavePhoto(userId, photoId, avatar)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestPhotoUseCase_SavePhoto_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photoId := uint(1)
	avatar := false
	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")

	photoRow := dataStruct.UserPhoto{
		Photo:  photoId,
		Avatar: avatar,
		UserId: userId,
	}

	photoRepoMock.EXPECT().SavePhoto(photoRow).Return(errRepo)

	err := photoUsecase.SavePhoto(userId, photoId, avatar)
	require.EqualError(t, err, errRepo.Error())
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestPhotoUseCase_DeletePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := []uint{1, 2, 3}

	num := 1
	userId := uint(1)

	photoRow := dataStruct.UserPhoto{
		Photo:  photos[num],
		UserId: userId,
	}

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{2, 3}, nil)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)
	photoRepoMock.EXPECT().DeletePhoto(photoRow).Return(nil)

	err := photoUsecase.DeletePhoto(userId, num)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestPhotoUseCase_DeletePhoto_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := []uint{1, 2, 3}
	errRepo := fmt.Errorf("something wrong")
	num := 1
	userId := uint(1)

	photoRow := dataStruct.UserPhoto{
		Photo:  photos[num],
		UserId: userId,
	}

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{2, 3}, nil)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)
	photoRepoMock.EXPECT().DeletePhoto(photoRow).Return(errRepo)

	err := photoUsecase.DeletePhoto(userId, num)
	require.EqualError(t, err, errRepo.Error())
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestPhotoUseCase_ChangePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := []uint{1, 2, 3}

	num := 2
	photoId := uint(1)
	userId := uint(1)

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{2, 3}, nil)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)
	photoRepoMock.EXPECT().ChangePhoto(photos[num], userId, photoId).Return(nil)

	err := photoUsecase.ChangePhoto(num, photoId, userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestPhotoUseCase_ChangePhoto_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	photoRepoMock := mock.NewMockIRepositoryPhoto(ctrl)
	photoUsecase := NewPhotoUseCase(photoRepoMock)

	photos := []uint{1, 2, 3}

	num := 2
	photoId := uint(1)
	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")

	photoRepoMock.EXPECT().GetPhotos(userId).Return([]uint{2, 3}, nil)
	photoRepoMock.EXPECT().GetAvatar(userId).Return(uint(1), nil)
	photoRepoMock.EXPECT().ChangePhoto(photos[num], userId, photoId).Return(errRepo)

	err := photoUsecase.ChangePhoto(num, photoId, userId)
	require.EqualError(t, err, errRepo.Error())
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
