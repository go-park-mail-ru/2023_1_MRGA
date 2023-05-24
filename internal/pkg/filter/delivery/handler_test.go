package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/filter/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestNewmHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)
	if filterHandler == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestHandler_GetFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)

	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	userId := uint(1)

	filterUsecaseMock.EXPECT().GetFilters(userId).Return(test, nil)
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"filters": test,
		},
		"status": 200,
	}
	req := httptest.NewRequest(http.MethodGet, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	filterHandler.GetFilter(w, req.WithContext(ctx))
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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

func TestHandler_GetFilter_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)

	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	filterUsecaseMock.EXPECT().GetFilters(userId).Return(test, errRepo)

	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 500,
	}
	req := httptest.NewRequest(http.MethodGet, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	keyContext := "userId"
	ctx := context.WithValue(req.Context(), keyContext, uint32(userId))
	filterHandler.GetFilter(w, req.WithContext(ctx))
	resp := w.Result()

	if resp.Status != "500 Internal Server Error" {
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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

func TestHandler_AddFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)
	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	filterUsecaseMock.EXPECT().AddFilters(userId, test).Return(nil)
	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	filterHandler.AddFilter(w, req.WithContext(ctx))
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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

func TestHandler_AddFilter_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)
	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	userId := uint(1)
	errRepo := fmt.Errorf("something wrong")
	filterUsecaseMock.EXPECT().AddFilters(userId, test).Return(errRepo)
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	filterHandler.AddFilter(w, req.WithContext(ctx))
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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

func TestHandler_ChangeFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)
	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	filterUsecaseMock.EXPECT().ChangeFilters(userId, test).Return(nil)
	filterUsecaseMock.EXPECT().GetFilters(userId).Return(test, nil)
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"filters": test,
		},
		"status": 200,
	}

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	filterHandler.ChangeFilter(w, req.WithContext(ctx))
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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

func TestHandler_ChangeFilter_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filterUsecaseMock := mock.NewMockUseCase(ctrl)
	filterHandler := NewHandler(filterUsecaseMock)
	errRepo := fmt.Errorf("something wrong")
	test := filter.FilterInput{
		MinAge:    20,
		MaxAge:    20,
		SearchSex: uint(0),
		Reason:    []string{"test"},
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	filterUsecaseMock.EXPECT().ChangeFilters(userId, test).Return(errRepo)

	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	filterHandler.ChangeFilter(w, req.WithContext(ctx))
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
	reqBody, err := ioutil.ReadAll(resp.Body)
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
