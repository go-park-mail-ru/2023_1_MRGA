package app

type IRepository interface {
	UploadFileV1(string, uint) (uint, error)
	UploadFile(string, uint) (error)
	GetFile(uint) (string, error)
}
