package user

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"

type UserRes struct {
	Username    string        `structs:"username"`
	Email       string        `structs:"email"`
	Age         int           `structs:"age"`
	Sex         constform.Sex `structs:"sex"`
	City        string        `structs:"city"`
	Description string        `structs:"description"`
	Avatar      string        `structs:"avatar"`
}
