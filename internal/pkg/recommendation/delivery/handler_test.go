package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestNewRecHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recUsecaseMock := mock.NewMockUseCase(ctrl)
	recHandler := NewHandler(recUsecaseMock)
	if recHandler == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestHandler_GetRecommendations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recUsecaseMock := mock.NewMockUseCase(ctrl)
	recHandler := NewHandler(recUsecaseMock)

	test := []recommendation.Recommendation{
		{
			Id:   uint(1),
			Name: "test1",
		},
	}
	userId := uint(1)
	recUsecaseMock.EXPECT().GetRecommendations(userId).Return(test, nil)
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"recommendations": test,
		},
		"status": 200,
	}
	req := httptest.NewRequest(http.MethodGet, "/meetme/recommendation", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	recHandler.GetRecommendations(w, req.WithContext(ctx))
	resp := w.Result()

	if resp.Status != "200 OK" {
		t.Errorf("incorrect result")
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Errorf(err.Error())
			return
		}
	}()
	reqBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &result)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !map_equal.MapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}

func TestHandler_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recUsecaseMock := mock.NewMockUseCase(ctrl)
	recHandler := NewHandler(recUsecaseMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	userId := uint(1)
	recUsecaseMock.EXPECT().GetRecommendations(userId).Return(nil, errRepo)

	req := httptest.NewRequest(http.MethodGet, "/meetme/recommendation", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	recHandler.GetRecommendations(w, req.WithContext(ctx))
	resp := w.Result()

	if resp.Status != "400 Bad Request" {
		t.Errorf("incorrect result")
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Errorf(err.Error())
			return
		}
	}()
	reqBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &result)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !map_equal.MapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}

func TestHandler_GetLikes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recUsecaseMock := mock.NewMockUseCase(ctrl)
	recHandler := NewHandler(recUsecaseMock)

	test := []recommendation.Recommendation{
		{
			Id:   uint(1),
			Name: "test1",
		},
	}
	userId := uint(1)
	recUsecaseMock.EXPECT().GetLikes(userId).Return(test, nil)
	recUsecaseMock.EXPECT().CheckProStatus(userId).Return(nil)
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"likes": test,
		},
		"status": 200,
	}
	req := httptest.NewRequest(http.MethodGet, "/meetme/recommendation", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	recHandler.GetLikes(w, req.WithContext(ctx))
	resp := w.Result()

	if resp.Status != "200 OK" {
		t.Errorf("incorrect result")
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Errorf(err.Error())
			return
		}
	}()
	reqBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &result)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !map_equal.MapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}

func TestHandler_GetLikes_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	recUsecaseMock := mock.NewMockUseCase(ctrl)
	recHandler := NewHandler(recUsecaseMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	userId := uint(1)
	recUsecaseMock.EXPECT().GetLikes(userId).Return(nil, errRepo)
	recUsecaseMock.EXPECT().CheckProStatus(userId).Return(nil)

	req := httptest.NewRequest(http.MethodGet, "/meetme/recommendation", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	recHandler.GetLikes(w, req.WithContext(ctx))
	resp := w.Result()

	if resp.Status != "400 Bad Request" {
		t.Errorf("incorrect result")
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Errorf(err.Error())
			return
		}
	}()
	reqBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &result)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !map_equal.MapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}
