package data_struct

type Complaint struct {
	ID     uint `sql:"unique;type:uuid;primary_key;default:" gorm:"primaryKey;unique"`
	UserId uint `gorm:"unique"`
	Count  int
}
