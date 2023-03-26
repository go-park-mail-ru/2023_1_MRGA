package dataStruct

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
)

type User struct {
	Id          uint          `sql:"unique;type:uuid;primary_key;default:" json:"userId" gorm:"primarykey;unique"`
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
