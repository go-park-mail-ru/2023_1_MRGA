package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/pkg/data_struct"
)

type ComplainRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *ComplainRepository {

	r := ComplainRepository{db}
	return &r
}

func (r *ComplainRepository) SaveComplaint(complaint dataStruct.Complaint) error {
	err := r.db.Create(&complaint).Error
	return err
}

func (r *ComplainRepository) IncrementComplaint(userId uint) error {
	var complaint dataStruct.Complaint
	err := r.db.Table("complaints").Find(&complaint, "user_id = ?", userId).Error
	if err != nil {
		return err
	}

	complaint.Count += 1
	err = r.db.Save(&complaint).Error
	return err
}

func (r *ComplainRepository) CheckCountComplaint(userId uint) (int, error) {
	var complaint dataStruct.Complaint
	err := r.db.Table("complaints").Find(&complaint, "user_id = ?", userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}

	return complaint.Count, err
}
