package auth

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
