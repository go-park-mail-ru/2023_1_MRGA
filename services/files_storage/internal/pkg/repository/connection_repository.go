package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/env_getter"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConn() (db *gorm.DB, err error) {
	dialector, err := getDialector()
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(postgres.Open(dialector), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Не удалось подключиться к БД")
	}

	err = makeMigrate(db)
	if err != nil {
		return nil, err
	}

	log.Println("К БД успешно подключился микросервис файлов")
	return db, nil
}

func getDialector() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Не удалось найти файл окружения")
	}

	host := env_getter.GetHostFromEnv("STORAGE_HOST")

	port := os.Getenv("STORAGE_PORT")
	user := os.Getenv("STORAGE_USER")
	password := os.Getenv("STORAGE_PASSWORD")
	dbname := os.Getenv("STORAGE_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname), nil
}

func makeMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&File{})
	if err != nil {
		return errors.New("Не удалось обновить таблицу photos")
	}

	return nil
}
