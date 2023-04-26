package token

import "github.com/google/uuid"

func CreateToken() string {
	newToken := uuid.NewString()
	return newToken
}
