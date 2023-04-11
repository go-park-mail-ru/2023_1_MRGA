package photo

type UseCase interface {
	SavePhoto(userId uint, photoId uint, avatar bool) error
	DeletePhoto(userId uint, photoId uint) error
}
