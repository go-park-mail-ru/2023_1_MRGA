package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func getDialector() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Не удалось найти файл окружения")
	}

	host := os.Getenv("CHAT_HOST")
	port := os.Getenv("CHAT_PORT")
	user := os.Getenv("CHAT_USER")
	password := os.Getenv("CHAT_PASSWORD")
	dbname := os.Getenv("CHAT_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname), nil
}

func getConn() (db *gorm.DB, err error) {
	dialector, err := getDialector()
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(postgres.Open(dialector), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Не удалось подключиться к БД чатов")
	}

	log.Println("К БД успешно подключился микросервис чатов")
	return db, nil
}

func (repo *Repository) makeMigrate() error {
	err := repo.db.AutoMigrate(&Message{})
	if err != nil {
		return errors.New("Не удалось обновить таблицу Message")
	}

	return nil
}

func InitRepository() (repo Repository, err error) {
	repo.db, err = getConn()

	repo.makeMigrate()

	return repo, err
}
