package repository

import "gorm.io/gorm"

var repository struct {
	db *gorm.DB
}

func InitRepository() error {
	var err error
	repository.db, err = getConn()
	return err
}
