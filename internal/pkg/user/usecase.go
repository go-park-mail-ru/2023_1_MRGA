package user

type UseCase interface {
	GetUserByToken(string) (UserRes, error)
}
