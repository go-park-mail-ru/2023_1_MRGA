package auth

type LoginInput struct {
	Input    string `json:"input"`
	Password string `json:"password"`
}
