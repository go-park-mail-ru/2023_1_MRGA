package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func InitRepository() (Repository, error) {
	var err error
	var repo Repository
	repo.db, err = getConn()
	return repo, err
}
