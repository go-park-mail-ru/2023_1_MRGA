package user

type UseCase interface {
	GetCurrentUser() (UserRes, error)
}
