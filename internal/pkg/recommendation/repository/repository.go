package repository

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/recommendation"
	"github.com/go-redis/redis"
)

type RecRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func NewRecRepo(db *gorm.DB, client *redis.Client) *RecRepository {
	r := RecRepository{db, client}

	return &r
}

func (r *RecRepository) GetRecommendation(userId uint) (recommendations []recommendation.Recommendation, err error) {

	// Get filtered users
	var filteredUsers []struct {
		name        string
		photo       string
		sex         uint
		birthDay    string
		description string
		city        string
	}
	err = r.db.Table("user_infos").
		Select("user_infos.name, user_photos.photo, user_infos.sex, users.birth_day, user_infos.description, cities.name").
		Joins("JOIN user_photos ON user_photos.user_id = user_infos.user_id").
		Joins("JOIN user_filters ON user_filters.user_id = user_infos.user_id").
		Joins("JOIN cities ON cities.id = user_infos.city_id").
		Joins("JOIN users ON users.id = user_infos.user_id").
		Where("user_infos.user_id != ?", userId).
		Where("user_photos.avatar = true").
		Where("user_infos.sex = user_filters.search_sex").
		Where("user_infos.birth_day BETWEEN ? AND ?", calculateBirthYear(10), calculateBirthYear(10)).
		Order("random()").
		Limit(10).
		Find(&filteredUsers).
		Error

	if err != nil {
		return recommendations, err
	}

	// Create recommendations from filtered users
	for _, user := range filteredUsers {
		recommendation := recommendation.Recommendation{
			Name:        user.name,
			Avatar:      user.photo,
			Age:         calculateAge(user.birthDay),
			Sex:         user.sex,
			Description: user.description,
			City:        user.city,
		}
		recommendations = append(recommendations, recommendation)
	}

	return recommendations, nil
}

func (r *RecRepository) GetUserIdByToken(token string) (uint, error) {
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

// Helper function to calculate birth year from age
func calculateBirthYear(age int) string {
	return fmt.Sprintf("%d-01-01", time.Now().Year()-age)
}

// Helper function to calculate age from birth date
func calculateAge(birthDay string) int {
	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, birthDay)
	age := time.Since(birthDate).Hours() / 24 / 365
	return int(age)
}
