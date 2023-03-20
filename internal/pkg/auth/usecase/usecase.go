package usecase

import (
	"crypto/sha1"
	"net/http"
	"time"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
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

func (a *AuthUseCase) Register(user *dataStruct.User) error {
	hashedPass := CreatePass(user.Password)
	user.Password = hashedPass

	if user.Avatar == "" {
		user.Avatar = DefaultAvatar
	}

	userId, err := a.userRepo.AddUser(*user)
	if err != nil {
		return err
	}

	userToken := token.CreateToken()
	a.userRepo.SaveToken(userId, userToken)

	http.SetCookie(w, &http.Cookie{
		Name:     SessionTokenCookieName,
		Value:    userToken,
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
	})

	return nil
}

func CreatePass(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte("123456789"))

	return string(bs)
}
