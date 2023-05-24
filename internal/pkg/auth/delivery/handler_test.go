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

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto/mocks"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/map_equal"
)

func TestAuthHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	AuthServiceMock := mock.NewMockAuthClient(ctrl)
	AuthHandler := NewHandler(AuthServiceMock)
	if AuthHandler == nil {
		t.Errorf("error")
	}
}

func TestHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	authInp := authProto.UserLoginInfo{
		Email:    "email",
		Password: "pass",
	}
	authOut := &authProto.UserResponse{
		Token: "token",
	}
	expected := []byte(`{"email": "email", "password": "pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/login", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	authServiceMock.EXPECT().Login(req.Context(), &authInp).Return(authOut, nil)

	authHandler.Login(w, req)
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

func TestHandler_Login_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)
	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}

	authInp := authProto.UserLoginInfo{
		Email:    "email",
		Password: "pass",
	}
	authOut := &authProto.UserResponse{
		Token: "token",
	}
	expected := []byte(`{"email": "email", "password": "pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	authServiceMock.EXPECT().Login(req.Context(), &authInp).Return(authOut, errRepo)

	authHandler.Login(w, req)
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

func TestHandler_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	userToken := "test"
	authInp := authProto.UserToken{
		Token: userToken,
	}
	expected := []byte(`{"email": "email", "password": "pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	req.AddCookie(&http.Cookie{Name: "session_token", Value: userToken})
	authServiceMock.EXPECT().Logout(req.Context(), &authInp).Return(nil, nil)

	authHandler.Logout(w, req)
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

func TestHandler_Logout_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 500,
	}
	userToken := "test"
	authInp := authProto.UserToken{
		Token: userToken,
	}
	expected := []byte(`{"email": "email", "password": "pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	req.AddCookie(&http.Cookie{Name: "session_token", Value: userToken})
	authServiceMock.EXPECT().Logout(req.Context(), &authInp).Return(nil, errRepo)

	authHandler.Logout(w, req)
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

func TestHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	authInp := authProto.UserRegisterInfo{
		Email:    "email",
		Password: "pass",
		Birthday: "01-01-1000",
	}
	authOut := &authProto.UserResponse{
		Token: "token",
	}
	expected := []byte(`{"email": "email", "password": "pass", "birthday": "01-01-1000"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	authServiceMock.EXPECT().Register(req.Context(), &authInp).Return(authOut, nil)

	authHandler.Register(w, req)
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

func TestHandler_Register_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	authInp := authProto.UserRegisterInfo{
		Email:    "email",
		Password: "pass",
		Birthday: "01-01-1000",
	}
	authOut := &authProto.UserResponse{
		Token: "token",
	}
	expected := []byte(`{"email": "email", "password": "pass", "birthday": "01-01-1000"}`)
	req := httptest.NewRequest(http.MethodPost, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()

	authServiceMock.EXPECT().Register(req.Context(), &authInp).Return(authOut, errRepo)

	authHandler.Register(w, req)
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

func TestHandler_ChangeUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	output := map[string]interface{}{
		"body":   map[string]interface{}{},
		"status": 200,
	}
	authInp := authProto.UserChangeInfo{
		UserId:   uint32(2),
		Email:    "email",
		Password: "pass",
		Birthday: "01-01-1000",
	}
	expected := []byte(`{"email": "email", "password": "pass", "birthday": "01-01-1000"}`)
	req := httptest.NewRequest(http.MethodDelete, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(2))
	req = req.WithContext(ctx)
	authServiceMock.EXPECT().ChangeUser(req.Context(), &authInp).Return(nil, nil)

	authHandler.ChangeUser(w, req)
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

func TestHandler_ChangeUser_GetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock.NewMockAuthClient(ctrl)
	authHandler := NewHandler(authServiceMock)

	errRepo := fmt.Errorf("something wrong")
	output := map[string]interface{}{
		"error":  errRepo.Error(),
		"status": 400,
	}
	authInp := authProto.UserChangeInfo{
		UserId:   uint32(2),
		Email:    "email",
		Password: "pass",
		Birthday: "01-01-1000",
	}
	expected := []byte(`{"email": "email", "password": "pass", "birthday": "01-01-1000"}`)
	req := httptest.NewRequest(http.MethodDelete, "/meetme/reaction", bytes.NewBuffer([]byte(expected)))
	w := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), "userId", uint32(2))
	req = req.WithContext(ctx)
	authServiceMock.EXPECT().ChangeUser(req.Context(), &authInp).Return(nil, errRepo)

	authHandler.ChangeUser(w, req)
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
