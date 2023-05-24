package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestNewRecHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)
	if matchHandler == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestHandler_GetMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)

	test := []match.UserRes{
		{
			UserId: uint(1),
			Name:   "test",
			Age:    20,
			Photo:  uint(1),
			Shown:  false,
		},
	}
	userId := uint(1)

	matchUsecaseMock.EXPECT().GetMatches(userId).Return(test, nil)
	output := map[string]interface{}{
		"body": map[string]interface{}{
			"matches": test,
		},
		"status": 200,
	}
	req := httptest.NewRequest(http.MethodGet, "/meetme/match", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.GetMatches(w, req.WithContext(ctx))
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

func TestHandler_GetMatches_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	userId := uint(1)

	matchUsecaseMock.EXPECT().GetMatches(userId).Return(nil, errRepo)

	req := httptest.NewRequest(http.MethodGet, "/meetme/match", nil)
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.GetMatches(w, req.WithContext(ctx))
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

func TestHandler_AddReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)

	userId := uint(1)
	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	reaction := match.ReactionInp{
		EvaluatedUserId: uint(2),
		Reaction:        "like",
	}
	rawTest, err := easyjson.Marshal(reaction)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	matchUsecaseMock.EXPECT().PostReaction(userId, reaction).Return(match.ReactionResult{ResultCode: 1}, nil)

	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.AddReaction(w, req.WithContext(ctx))
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

func TestHandler_AddReaction_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)

	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	reaction := match.ReactionInp{
		EvaluatedUserId: uint(2),
		Reaction:        "like",
	}

	rawTest, err := easyjson.Marshal(reaction)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	matchUsecaseMock.EXPECT().PostReaction(userId, reaction).Return(match.ReactionResult{ResultCode: 1}, errRepo)

	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer(rawTest))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.AddReaction(w, req.WithContext(ctx))
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

func TestHandler_DeleteMatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)

	userId := uint(1)
	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}

	matchUsecaseMock.EXPECT().DeleteMatch(userId, uint(2)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/meetme/match/2", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"userId": "2",
	}
	req = mux.SetURLVars(req, vars)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.DeleteMatch(w, req.WithContext(ctx))
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

func TestHandler_DeleteMatch_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	matchUsecaseMock := mock.NewMockUseCase(ctrl)
	matchHandler := NewHandler(matchUsecaseMock)
	errRepo := fmt.Errorf("something wrong")
	userId := uint(1)
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	matchUsecaseMock.EXPECT().DeleteMatch(userId, uint(2)).Return(errRepo)

	req := httptest.NewRequest(http.MethodDelete, "/meetme/match/2", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"userId": "2",
	}
	req = mux.SetURLVars(req, vars)
	ctx := context.WithValue(req.Context(), "userId", uint32(userId))
	matchHandler.DeleteMatch(w, req.WithContext(ctx))
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
