package photo

//go:generate mockgen -source=usecase.go -destination=mocks/usecase.go -package=mock
type UseCase interface {
	SavePhoto(userId uint, photoId uint, avatar bool) error
	DeletePhoto(userId uint, photoId int) error
	ChangePhoto(num int, photoId uint, userId uint) error

	GetAllPhotos(userId uint) ([]uint, error)
	GetAvatar(userId uint) (uint, error)
}
