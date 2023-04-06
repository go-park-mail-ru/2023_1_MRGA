package repository

import (
	"strconv"
	"time"

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

func NewAuthRepo(db *gorm.DB, client *redis.Client) *AuthRepository {

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
	err = r.db.Table("users").Select("users.id, email", "users.id = ?", userId).
		Joins("JOIN user_photos on users.id=user_photos.id").
		Joins("Join user_infos on users.id = user_infos.id").
		Find(&user).Error
	if err != nil {
		return
	}

	return user, nil

	//return userRes, fmt.Errorf("user are not found")
}

func (r *AuthRepository) GetUserIdByToken(token string) (uint, error) {
	user, err := r.client.Get(token).Result()
	if err != nil {
		return 0, err
	}
	userId, err := strconv.Atoi(user)
	if err != nil {
		return 0, err
	}
	return uint(userId), nil
}

func (r *AuthRepository) DeleteToken(token string) error {
	err := r.client.Del(token).Err()
	return err
}

func (r *AuthRepository) SaveToken(userId uint, token string) (err error) {
	userIdStr := strconv.Itoa(int(userId))
	err = r.client.Set(token, userIdStr, 200*time.Second).Err()
	return
}
