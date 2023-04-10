package repository

type File struct {
	ID   uint   `gorm:"primary_key"`
	Path string `gorm:"not null"`
}
