package repository

import (
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/auth"
)

type AuthRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func NewRepo(db *gorm.DB, client *redis.Client) *AuthRepository {

	r := AuthRepository{db, client}
	return &r
}

func (r *AuthRepository) Login(email string, passwordInp string) (uint, error) {
	var userDB *dataStruct.User
	err := r.db.Model(&dataStruct.User{}).Where("email = ?", email).Take(&userDB).Error
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
	if err := r.CheckBirthDay(user.BirthDay); err != nil {
		return 0, err
	}

	err := r.db.Create(user).Error
	return user.Id, err
}

func (r *AuthRepository) GetUserById(userId uint) (userRes auth.UserRes, err error) {

	user := auth.UserRes{}
	err = r.db.Table("users").Select("users.id, email").
		Joins("JOIN user_photos on users.id=user_photos.id").
		Joins("Join user_infos on users.id = user_infos.id").
		Find(&user).Error
	if err != nil {
		return
	}

	return user, nil

	//return userRes, fmt.Errorf("user are not found")
}

func (r *AuthRepository) GetUserIdByToken(InpToken string) (uint, error) {
	//for userId, userToken := range r.UserTokens {
	//	if userToken == InpToken {
	//		return userId, nil
	//	}
	//}
	return 1, nil
	//return 1, fmt.Errorf("user are not found")
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
