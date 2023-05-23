package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestNewInfoHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoUsecaseMock := mock.NewMockUseCase(ctrl)
	infoHandler := NewHandler(infoUsecaseMock)
	if infoHandler == nil {
		t.Errorf("incorrect result")
		return
	}
}

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoUsecaseMock := mock.NewMockUseCase(ctrl)
	infoHandler := NewHandler(infoUsecaseMock)
	type callFunc func(w http.ResponseWriter, r *http.Request)
	testCases := []struct {
		OutputParam string
		Target      string
		Function    callFunc
	}{
		{
			"hashtags",
			"/api/hashtags",
			infoHandler.GetHashtags,
		},
		{
			"jobs",
			"/api/jobs",
			infoHandler.GetJobs,
		},
		{
			"education",
			"/api/education",
			infoHandler.GetEducation,
		},
		{
			"zodiac",
			"/api/zodiac",
			infoHandler.GetZodiac,
		},
		{
			"cities",
			"/api/cities",
			infoHandler.GetCities,
		},
		{
			"status",
			"/api/statuses",
			infoHandler.GetStatuses,
		},
		{
			"reasons",
			"/api/reasons",
			infoHandler.GetReasons,
		},
	}
	test := []string{"test1", "test2"}

	infoUsecaseMock.EXPECT().GetHashtags().Return(test, nil)
	infoUsecaseMock.EXPECT().GetJobs().Return(test, nil)
	infoUsecaseMock.EXPECT().GetEducation().Return(test, nil)
	infoUsecaseMock.EXPECT().GetZodiacs().Return(test, nil)
	infoUsecaseMock.EXPECT().GetStatuses().Return(test, nil)
	infoUsecaseMock.EXPECT().GetCities().Return(test, nil)
	infoUsecaseMock.EXPECT().GetReasons().Return(test, nil)

	for _, tCase := range testCases {
		output := map[string]interface{}{
			"body": map[string]interface{}{
				tCase.OutputParam: test,
			},
			"status": 200,
		}
		req := httptest.NewRequest(http.MethodGet, tCase.Target, nil)
		w := httptest.NewRecorder()
		tCase.Function(w, req)
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

}

func TestHandler_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	infoUsecaseMock := mock.NewMockUseCase(ctrl)
	infoHandler := NewHandler(infoUsecaseMock)
	type callFunc func(w http.ResponseWriter, r *http.Request)
	testCases := []struct {
		Target   string
		Function callFunc
	}{
		{
			"/api/hashtags",
			infoHandler.GetHashtags,
		},
		{
			"/api/jobs",
			infoHandler.GetJobs,
		},
		{
			"/api/education",
			infoHandler.GetEducation,
		},
		{
			"/api/zodiac",
			infoHandler.GetZodiac,
		},
		{
			"/api/cities",
			infoHandler.GetCities,
		},
		{
			"/api/status",
			infoHandler.GetStatuses,
		},
		{
			"/api/reasons",
			infoHandler.GetReasons,
		},
	}
	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 500,
	}

	infoUsecaseMock.EXPECT().GetHashtags().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetJobs().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetEducation().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetZodiacs().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetStatuses().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetCities().Return(nil, errRepo)
	infoUsecaseMock.EXPECT().GetReasons().Return(nil, errRepo)

	for _, tCase := range testCases {

		req := httptest.NewRequest(http.MethodGet, tCase.Target, nil)
		w := httptest.NewRecorder()
		tCase.Function(w, req)
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

}
