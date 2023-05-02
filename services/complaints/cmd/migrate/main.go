package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/app/dsn"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/pkg/data_struct"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&dataStruct.Complaint{})
	if err != nil {
		log.Println("cant migrate db user")
		os.Exit(2)
	}
}
