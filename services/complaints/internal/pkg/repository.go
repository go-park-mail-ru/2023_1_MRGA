package compRepo

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/pkg/data_struct"

type CompRepo interface {
	SaveComplaint(complaint dataStruct.Complaint) error
	IncrementComplaint(userId uint) error
	CheckCountComplaint(userId uint) (int, error)
}
