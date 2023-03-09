package token

import "github.com/google/uuid"

var UserTokens = map[string]string{}

func CreateToken() string {
	newToken := uuid.NewString()
	return newToken
}
