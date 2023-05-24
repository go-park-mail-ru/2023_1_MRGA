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
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/middleware"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	mockProto "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)
	if compHandler == nil {
		t.Errorf("error")

	}
}

func TestHandler_AddUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}
	test := info_user.HashtagInp{
		Hashtag: testInp,
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"hashtags": testInp,
		},
		"status": 200,
	}
	compUsecaseMock.EXPECT().AddHashtags(userId, test).Return(nil)
	compUsecaseMock.EXPECT().GetUserHashtags(userId).Return(testInp, nil)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	compHandler.AddUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_AddUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}
	test := info_user.HashtagInp{
		Hashtag: testInp,
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	compUsecaseMock.EXPECT().AddHashtags(userId, test).Return(errRepo)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	compHandler.AddUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_ChangeUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}
	test := info_user.HashtagInp{
		Hashtag: testInp,
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"hashtags": testInp,
		},
		"status": 200,
	}
	compUsecaseMock.EXPECT().ChangeUserHashtags(userId, test).Return(nil)
	compUsecaseMock.EXPECT().GetUserHashtags(userId).Return(testInp, nil)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), middleware.ContextUserKey, uint32(userId))
	compHandler.ChangeUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_ChangeUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}
	test := info_user.HashtagInp{
		Hashtag: testInp,
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	compUsecaseMock.EXPECT().ChangeUserHashtags(userId, test).Return(errRepo)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.ChangeUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_ChangeUserStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	test := info_user.StatusInp{
		Status: "test",
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	compUsecaseMock.EXPECT().ChangeUserStatus(userId, test).Return(nil)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.ChangeUserStatus(w, req.WithContext(ctx))
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

func TestHandler_ChangeUserStatus_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)
	errRepo := fmt.Errorf("something wrong")
	test := info_user.StatusInp{
		Status: "test",
	}
	userId := uint(1)
	rawTest, err := easyjson.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	compUsecaseMock.EXPECT().ChangeUserStatus(userId, test).Return(errRepo)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.ChangeUserStatus(w, req.WithContext(ctx))
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

func TestHandler_GetUserHashtags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}

	userId := uint(1)

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"hashtags": testInp,
		},
		"status": 200,
	}

	compUsecaseMock.EXPECT().GetUserHashtags(userId).Return(testInp, nil)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.GetUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_GetUserHashtags_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	testInp := []string{
		"test1",
		"test2",
	}

	userId := uint(1)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().GetUserHashtags(userId).Return(testInp, errRepo)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.GetUserHashtags(w, req.WithContext(ctx))
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

func TestHandler_GetUserStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	test := info_user.StatusInp{
		Status: "test",
	}
	userId := uint(1)

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"statuses": test.Status,
		},
		"status": 200,
	}

	compUsecaseMock.EXPECT().GetUserStatus(userId).Return(test.Status, nil)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.GetUserStatus(w, req.WithContext(ctx))
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

func TestHandler_GetUserStatus_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	test := info_user.StatusInp{
		Status: "test",
	}

	userId := uint(1)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().GetUserStatus(userId).Return(test.Status, errRepo)
	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	compHandler.GetUserStatus(w, req.WithContext(ctx))
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

func TestHandler_GetCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)
	test := info_user.UserRes{
		UserId: userId,
		Name:   "test",
		Age:    20,
		Avatar: uint(1),
		Step:   constform.FullInfo,
		Banned: false,
	}

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"userId": test.UserId,
			"name":   test.Name,
			"age":    test.Age,
			"avatar": test.Avatar,
			"step":   test.Step,
			"banned": false,
		},
		"status": 200,
	}

	compInp := &complaintProto.UserId{UserId: uint32(userId)}
	banned := complaintProto.Response{
		Banned: false,
	}

	compUsecaseMock.EXPECT().GetUserById(userId).Return(test, nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)
	compServiceMock.EXPECT().CheckBanned(req.Context(), compInp).Return(&banned, nil)
	w := httptest.NewRecorder()

	compHandler.GetCurrentUser(w, req)
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

func TestHandler_GetCurrentUser_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)
	test := info_user.UserRes{
		UserId: userId,
		Name:   "test",
		Age:    20,
		Avatar: uint(1),
		Step:   constform.FullInfo,
		Banned: false,
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().GetUserById(userId).Return(test, errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.GetCurrentUser(w, req)
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

func TestHandler_GetInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"name":        test.Name,
			"age":         test.Age,
			"photos":      test.Photos,
			"email":       test.Email,
			"eduction":    test.Education,
			"zodiac":      test.Zodiac,
			"city":        test.City,
			"job":         test.Job,
			"sex":         test.Sex,
			"description": test.Description,
		},
		"status": 200,
	}

	compUsecaseMock.EXPECT().GetInfo(userId).Return(test, nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.GetInfo(w, req)
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

func TestHandler_GetInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().GetInfo(userId).Return(test, errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.GetInfo(w, req)
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

func TestHandler_GetInfoById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"name":        test.Name,
			"age":         test.Age,
			"photos":      test.Photos,
			"email":       test.Email,
			"eduction":    test.Education,
			"zodiac":      test.Zodiac,
			"city":        test.City,
			"job":         test.Job,
			"sex":         test.Sex,
			"description": test.Description,
		},
		"status": 200,
	}

	compUsecaseMock.EXPECT().GetInfo(userId).Return(test, nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)
	vars := map[string]string{
		"userId": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	compHandler.GetInfoById(w, req)
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

func TestHandler_GetInfoById_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().GetInfo(userId).Return(test, errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", nil)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)
	vars := map[string]string{
		"userId": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	compHandler.GetInfoById(w, req)
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

func TestHandler_ChangeInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	testInp := info_user.InfoChange{
		Description: "test",
	}
	rawTest, err := easyjson.Marshal(testInp)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	output := map[string]interface{}{
		"body": map[string]interface{}{
			"name":        test.Name,
			"age":         test.Age,
			"photos":      test.Photos,
			"email":       test.Email,
			"eduction":    test.Education,
			"zodiac":      test.Zodiac,
			"city":        test.City,
			"job":         test.Job,
			"sex":         test.Sex,
			"description": test.Description,
		},
		"status": 200,
	}

	compUsecaseMock.EXPECT().ChangeInfo(userId, testInp).Return(test, nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.ChangeInfo(w, req)
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

func TestHandler_ChangeInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)
	testInp := info_user.InfoChange{
		Description: "test",
	}
	rawTest, err := easyjson.Marshal(testInp)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	test := info_user.InfoStructAnswer{
		Name:        "test",
		Age:         20,
		Photos:      []uint{1},
		Email:       "",
		City:        "test",
		Education:   "test",
		Zodiac:      "test",
		Job:         "test",
		Sex:         uint(1),
		Description: "test",
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().ChangeInfo(userId, testInp).Return(test, errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.ChangeInfo(w, req)
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

func TestHandler_CreateInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)

	testInp := info_user.InfoStruct{
		Description: "test",
	}
	rawTest, err := easyjson.Marshal(testInp)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}

	compUsecaseMock.EXPECT().AddInfo(userId, testInp).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.CreateInfo(w, req)
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

func TestHandler_CreateInfo_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	compServiceMock := mockProto.NewMockComplaintsClient(ctrl)
	compUsecaseMock := mock.NewMockUseCase(ctrl)
	compHandler := NewHandler(compUsecaseMock, compServiceMock)

	userId := uint(1)
	testInp := info_user.InfoStruct{
		Description: "test",
	}
	rawTest, err := easyjson.Marshal(testInp)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	compUsecaseMock.EXPECT().AddInfo(userId, testInp).Return(errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/filters", bytes.NewBuffer(rawTest))
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	compHandler.CreateInfo(w, req)
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
