package app

type IRepository interface {
	UploadPhoto(string, uint) (uint, error)
	UploadFile(string, uint) error
	GetFile(uint) (string, error)
}
