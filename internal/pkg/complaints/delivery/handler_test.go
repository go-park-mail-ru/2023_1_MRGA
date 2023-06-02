package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	mock "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	CompServiceMock := mock.NewMockComplaintsClient(ctrl)
	CompHandler := NewHandler(CompServiceMock)
	if CompHandler == nil {
		t.Errorf("error")
	}
}

func TestHandler_Complain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	CompServiceMock := mock.NewMockComplaintsClient(ctrl)
	CompHandler := NewHandler(CompServiceMock)

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	authInp := complaintProto.UserId{
		UserId: uint32(2),
	}
	expected := []byte(`{"userId": 2}`)
	req := httptest.NewRequest(http.MethodDelete, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	CompServiceMock.EXPECT().Complain(gomock.All(), &authInp).Return(nil, nil)

	CompHandler.Complain(w, req)
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

func TestHandler_Complain_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	CompServiceMock := mock.NewMockComplaintsClient(ctrl)
	CompHandler := NewHandler(CompServiceMock)

	errRepo := fmt.Errorf("something wrong")

	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	authInp := complaintProto.UserId{
		UserId: uint32(2),
	}
	expected := []byte(`{"userId": 2}`)
	req := httptest.NewRequest(http.MethodDelete, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	CompServiceMock.EXPECT().Complain(gomock.All(), &authInp).Return(nil, errRepo)

	CompHandler.Complain(w, req)
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
