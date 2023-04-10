package usecase

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/token"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
)

type AuthUseCase struct {
	userRepo       auth.IRepositoryAuth
	hashSalt       string
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.IRepositoryAuth,
	hashSalt string,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) Register(user *dataStruct.User) (string, error) {
	hashedPass, err := CreatePass(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPass

	userId, err := a.userRepo.AddUser(user)
	if err != nil {
		return "", err
	}

	userToken := token.CreateToken()
	err = a.userRepo.SaveToken(userId, userToken)
	if err != nil {
		return "", err
	}

	return userToken, nil
}

//if user.Avatar == "" {
//user.Avatar = _default.DefaultAvatar
//}
func (a *AuthUseCase) Login(logInp auth.LoginInput) (string, error) {

	if logInp.Email == "" {
		err := fmt.Errorf("email is empty")
		return "", err
	}

	userId, err := a.userRepo.Login(logInp.Email, logInp.Password)
	if err != nil {
		return "", err
	}

	userToken := token.CreateToken()
	err = a.userRepo.SaveToken(userId, userToken)
	if err != nil {
		return "", err
	}

	return userToken, nil
}

func (u *AuthUseCase) GetUserById(userId uint) (user auth.UserRes, err error) {

	userTemp, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return
	}
	user.Name = userTemp.Name
	user.Email = userTemp.Email
	photos, err := u.userRepo.GetUserPhoto(userId)
	if err != nil {
		return
	}
	for _, p := range photos {
		photo := auth.Photo{PhotoId: p.Photo, Avatar: p.Avatar}
		user.Photos = append(user.Photos, photo)
	}
	return
}

func (u *AuthUseCase) ChangeUser(user dataStruct.User) error {
	if user.Password != "" {
		hashedPass, err := CreatePass(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPass
	}
	err := u.userRepo.ChangeUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUseCase) Logout(token string) error {

	err := a.userRepo.DeleteToken(token)
	return err
}

func CreatePass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
