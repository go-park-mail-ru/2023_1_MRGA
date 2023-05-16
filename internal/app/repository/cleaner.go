package repository

import (
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

func CleanCount(db *gorm.DB) error {
	err := db.Model(&dataStruct.User{}).Where("count > ?", 0).Update("count", 0).Error
	return err
}
