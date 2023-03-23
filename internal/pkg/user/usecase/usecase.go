package usecase

import (
	"log"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/user"
)

type UserUseCase struct {
	userRepo user.IRepositoryUser
}

func NewUserUseCase(
	userRepo user.IRepositoryUser) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u *UserUseCase) GetUserByToken(token string) (user user.UserRes, errr error) {
	userId, err := u.userRepo.GetUserIdByToken(token)
	log.Println(userId)
	if err != nil {
		return
	}

	user, err = u.userRepo.GetUserById(userId)
	if err != nil {
		return
	}

	return
}
