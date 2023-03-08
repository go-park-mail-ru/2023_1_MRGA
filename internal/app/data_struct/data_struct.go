package data_struct

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
)

type User struct {
	UserId      uint          `json:"userId"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Age         int           `json:"age"`
	Sex         constform.Sex `json:"sex"`
	City        string        `json:"city"`
	Description string        `json:"description"`
	Avatar      string        `json:"avatar"`
}

type City struct {
	CityId uint   `json:"cityId"`
	Name   string `json:"name"`
}
