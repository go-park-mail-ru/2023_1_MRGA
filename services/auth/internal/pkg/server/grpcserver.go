package server

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/app/token"
	authRepo "github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/internal/pkg"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/auth/pkg/data_struct"
	auth "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/authProto"
)

type GRPCServer struct {
	AuthRepo authRepo.UserRepo
}

func NewGPRCServer(authRepo authRepo.UserRepo) *GRPCServer {
	return &GRPCServer{
		AuthRepo: authRepo,
	}
}

func (s *GRPCServer) Register(ctx context.Context, req *auth.UserRegisterInfo) (*auth.UserResponse, error) {
	user := dataStruct.User{Email: req.Email, Password: req.Password, BirthDay: req.Birthday}

	hashedPass, err := CreatePass(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass

	userId, err := s.AuthRepo.AddUser(&user)
	if err != nil {
		return nil, err
	}

	userToken := token.CreateToken()
	err = s.AuthRepo.SaveToken(userId, userToken)
	if err != nil {
		return nil, err
	}

	return &auth.UserResponse{
		UserId: uint32(userId),
		Token:  userToken,
		Ok:     true,
	}, nil
}

func (s *GRPCServer) Login(ctx context.Context, req *auth.UserLoginInfo) (*auth.UserResponse, error) {
	if req.Email == "" {
		err := fmt.Errorf("email is empty")
		return nil, err
	}

	userId, err := s.AuthRepo.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	userToken := token.CreateToken()
	err = s.AuthRepo.SaveToken(userId, userToken)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &auth.UserResponse{
		UserId: uint32(userId),
		Token:  userToken,
		Ok:     true,
	}, nil
}

func (s *GRPCServer) CheckSession(ctx context.Context, req *auth.UserToken) (*auth.UserResponse, error) {
	userId, err := s.AuthRepo.CheckSession(req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.UserResponse{
		UserId: uint32(userId),
		Ok:     true,
	}, nil
}

func (s *GRPCServer) Logout(ctx context.Context, req *auth.UserToken) (*auth.Response, error) {
	err := s.AuthRepo.DeleteToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.Response{
		Ok: true,
	}, nil
}

func (s *GRPCServer) ChangeUser(ctx context.Context, req *auth.UserChangeInfo) (*auth.Response, error) {
	user := dataStruct.User{Id: uint(req.UserId), Email: req.Email, Password: req.Password, BirthDay: req.Birthday}
	if user.Password != "" {
		hashedPass, err := CreatePass(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPass
	}
	err := s.AuthRepo.ChangeUser(user)
	if err != nil {
		return nil, err
	}
	return &auth.Response{
		Ok: true,
	}, nil
}

func CreatePass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
