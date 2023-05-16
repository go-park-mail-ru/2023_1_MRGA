package dsn

import (
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/utils/env_getter"
)

func FromEnv() string {

	host := env_getter.GetHostFromEnv("DB_HOST")
	if host == "" {
		return ""
	}

	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}
