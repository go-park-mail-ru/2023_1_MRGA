package app

type IRepository interface {
	UploadFile(string, uint) (uint, error)
	GetFile(uint) (string, error)
}
