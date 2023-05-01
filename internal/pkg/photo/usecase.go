package photo

type UseCase interface {
	SavePhoto(userId uint, photoId uint, avatar bool) error
	DeletePhoto(userId uint, photoId uint) error
	ChangePhoto(num int, photoId uint, userId uint) error

	GetAllPhotos(userId uint) ([]uint, error)
	GetAvatar(userId uint) (uint, error)
}
