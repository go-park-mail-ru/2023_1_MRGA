package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *AuthRepository {

	r := AuthRepository{db}
	return &r
}

func (r *AuthRepository) Login(input string, passwordInp string) (uint, error) {
	var userDB *dataStruct.User
	err := r.db.Model(&dataStruct.User{}).Where("username = ? or email = ?", input, input).Take(&userDB).Error
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(passwordInp))
	if err != nil {
		return 0, err
	}
	if passwordInp == userDB.Password {

	}
	return userDB.Id, nil
}

func (r *AuthRepository) AddUser(user *dataStruct.User) (uint, error) {

	if err := r.CheckUsername(user.Username); err != nil {
		return 0, err
	}

	if err := r.CheckEmail(user.Email); err != nil {
		return 0, err
	}

	if err := CheckAge(user.Age); err != nil {
		return 0, err
	}

	err := r.db.Create(user).Error
	return user.Id, err
}

//
//func (r *AuthRepository) DeleteToken(token string) error {
//	var userId uint
//	flagFound := false
//	for indexUser, tokenDS := range r.UserTokens {
//		if tokenDS == token {
//			userId = indexUser
//			flagFound = true
//			break
//		}
//	}
//
//	if !flagFound {
//		return fmt.Errorf("UnAuthorised")
//	}
//
//	delete(r.UserTokens, userId)
//	return nil
//}
//
//func (r *AuthRepository) SaveToken(userId uint, token string) {
//	tokenUser := r.UserTokens
//	tokenUser[userId] = token
//	r.UserTokens = tokenUser
//}
