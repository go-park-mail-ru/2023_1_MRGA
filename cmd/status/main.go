package main

import (
	"flag"
	"fmt"
	"log"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/dsn"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	all   = flag.Bool("all", false, "Обновить статус для всех пользователей")
	email = flag.String("email", "", "Email пользователя")
	s     = flag.Uint("s", 0, "Статус пользователя")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Не удалось подключиться к env файлу: " + err.Error())
	}

	flag.Parse()

	if *s == 0 {
		log.Fatalf("Не передан флаг s")
	}

	fmt.Println("Скрипт запущен")

	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: " + err.Error())
	}

	if *all {
		err = db.Table("users").Where("status <> ?", *s).Updates(dataStruct.User{Status: *s}).Error
		if err != nil {
			log.Fatalf("Не удалось обновить статусы всех пользователей: " + err.Error())
		}
		fmt.Println("Статус всех пользователей обновлен.")
		return
	}

	if *email != "" {
		err = db.Table("users").Where("email = ?", *email).Updates(dataStruct.User{Status: *s}).Error
		if err != nil {
			log.Fatalf("Не удалось обновить статус пользователя: " + err.Error())
		}
		fmt.Printf("Статус пользователя с email %s обновлен.\n", *email)
		return
	}

	fmt.Println("Вы должны указать либо флаг -all, либо флаг -email.")
}
