package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&dataStruct.User{})
	if err != nil {
		log.Println("cant migrate db user")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserInfo{})
	if err != nil {
		log.Println("cant migrate db userInfo")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Sex{})
	if err != nil {
		log.Println("cant migrate db sex")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserPhoto{})
	if err != nil {
		log.Println("cant migrate db userPhoto")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.City{})
	if err != nil {
		log.Println("cant migrate db city")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Zodiac{})
	if err != nil {
		log.Println("cant migrate db zodiac")
		os.Exit(2)
	}
	err = db.AutoMigrate(&dataStruct.Job{})
	if err != nil {
		log.Println("cant migrate db job")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Education{})
	if err != nil {
		log.Println("cant migrate db education")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserFilter{})
	if err != nil {
		log.Println("cant migrate db userFilter")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserReason{})
	if err != nil {
		log.Println("cant migrate db userReason")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Reason{})
	if err != nil {
		log.Println("cant migrate db reason")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserHashtag{})
	if err != nil {
		log.Println("cant migrate db userHashtag")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Hashtag{})
	if err != nil {
		log.Println("cant migrate db hashtag")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserReaction{})
	if err != nil {
		log.Println("cant migrate db userReaction")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Reaction{})
	if err != nil {
		log.Println("cant migrate db reaction")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.UserHistory{})
	if err != nil {
		log.Println("cant migrate db userHistory")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Match{})
	if err != nil {
		log.Println("cant migrate db match")
		os.Exit(2)
	}

	err = db.AutoMigrate(&dataStruct.Status{})
	if err != nil {
		log.Println("cant migrate db status")
		os.Exit(2)
	}

}
