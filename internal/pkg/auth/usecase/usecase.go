package usecase

import (
	"crypto/sha1"
	"fmt"
	"time"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
	_default "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/default"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth/delivery"
)

type AuthUseCase struct {
	userRepo       auth.IRepositoryAuth
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.IRepositoryAuth,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) Register(user *dataStruct.User) (string, error) {
	hashedPass := CreatePass(user.Password)
	user.Password = hashedPass

	if user.Avatar == "" {
		user.Avatar = _default.DefaultAvatar
	}

	userId, err := a.userRepo.AddUser(*user)
	if err != nil {
		return "", err
	}

	userToken := token.CreateToken()
	a.userRepo.SaveToken(userId, userToken)

	return userToken, nil
}

func (a *AuthUseCase) Login(logInp delivery.LoginInput) (string, error) {
	hashPass := CreatePass(logInp.Password)

	if logInp.Input == "" {
		err := fmt.Errorf("email and username are empty")
		return "", err
	}

	userId, err := a.userRepo.Login(logInp.Input, hashPass)
	if err != nil {
		return "", err
	}

	userToken := token.CreateToken()
	a.userRepo.SaveToken(userId, userToken)

	return userToken, nil
}

func (a *AuthUseCase) Logout(token string) error {
	err := a.userRepo.DeleteToken(token)
	return err
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
