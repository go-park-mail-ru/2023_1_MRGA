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

func (r *AuthRepository) ChangeUser(user dataStruct.User) error {
	userDb := &dataStruct.User{}
	err := r.db.First(userDb, "id= ?", user.Id).Error
	if err != nil {
		return err
	}
	if user.Email != "" {
		userDb.Email = user.Email
	}
	if user.BirthDay != "" {
		userDb.BirthDay = user.BirthDay
	}
	if user.Password != "" {
		userDb.Password = user.Password
	}

	err = r.db.Save(&userDb).Error
	return err
}

func (r *AuthRepository) GetUserById(userId uint) (userRes auth.UserRestTemp, err error) {

	user := auth.UserRestTemp{}
	err = r.db.Table("users").Select("user_infos.name").
		Where("users.id =?", userId).
		Joins("Join user_infos on users.id = user_infos.user_id").
		Find(&user).Error
	if err != nil {
		return
	}

	return user, nil
}

func (r *AuthRepository) GetUserPhoto(userId uint) (photos []dataStruct.UserPhoto, err error) {
	err = r.db.Find(&photos, "user_id = ?", userId).Error
	return
}

func (r *AuthRepository) GetAvatarId(userId uint) (avatar uint, err error) {
	err = r.db.Table("user_photos up").Select("up.photo").Where("user_id = ? AND avatar = ?", userId, true).Find(&avatar).Error
	return
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
	err = r.client.Set(token, userIdStr, 120*time.Hour).Err()
	return
}

func (r *AuthRepository) GetAge(userId uint) (int, error) {
	var user dataStruct.User
	err := r.db.First(&user, "id=?", userId).Error
	if err != nil {
		return 0, err
	}
	age, err := CalculateAge(user.BirthDay)
	if err != nil {
		return 0, err
	}
	return age, nil
}
