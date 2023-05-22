package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match/mocks"
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

func TestHandler(t *testing.T) {
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
	if !mapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}

func TestHandler_GetError(t *testing.T) {
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
	if !mapEqual(result, output) {
		t.Errorf(" wrong result, expected %#v, got %#v", output, result)
	}
}

func mapEqual(got, expected map[string]interface{}) bool {
	for keyGot, valueGot := range got {
		valueExpected := expected[keyGot]

		var (
			strValueExpected string = convertToString(valueExpected)
			strValueGot      string = convertToString(valueGot)
		)
		if strValueExpected != strValueGot {
			return false
		}
	}
	return true
}

func convertToString(val interface{}) (strVal string) {
	switch val.(type) {
	case string:
		strVal = val.(string)
	case float64:
		strVal = fmt.Sprint(val.(float64))
	case int:
		strVal = fmt.Sprint(val.(int))
	}
	return strVal
}
