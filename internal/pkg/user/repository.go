package user

type IRepositoryUser interface {
	GetUserIdByToken(string) (uint, error)
	GetUserById(uint) (UserRes, error)
}
