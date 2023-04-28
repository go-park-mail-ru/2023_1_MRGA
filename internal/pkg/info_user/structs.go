package info_user

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"

type InfoStruct struct {
	Name        string `json:"name" structs:"name"`
	City        string `json:"city" structs:"city"`
	Email       string `json:"email" structs:"email"`
	Sex         uint   `json:"sex" structs:"sex"`
	Description string `json:"description" structs:"description"`
	Zodiac      string `json:"zodiac" structs:"zodiac"`
	Job         string `json:"job" structs:"job"`
	Education   string `json:"education" structs:"education"`
}
type UserRestTemp struct {
	Name string `json:"name"`
}

type UserRes struct {
	Name   string         `json:"name" structs:"name"`
	Age    int            `json:"age" structs:"age"`
	Avatar uint           `json:"avatarId" structs:"avatarId"`
	Step   constform.Step `json:"step" structs:"step"`
}

type InfoStructAnswer struct {
	Name        string `json:"name" structs:"name"`
	Email       string `json:"email" structs:"email"`
	City        string `json:"city" structs:"city"`
	Sex         uint   `json:"sex" structs:"sex"`
	Description string `json:"description" structs:"description"`
	Zodiac      string `json:"zodiac" structs:"zodiac"`
	Job         string `json:"job" structs:"job" structs:"job"`
	Education   string `json:"education" structs:"education"`
	Age         int    `json:"age" structs:"age"`
	Photos      []uint `json:"photos" structs:"photos"`
}

type InfoChange struct {
	Name        string `json:"name" structs:"name"`
	City        string `json:"city" structs:"city"`
	Sex         uint   `json:"sex" structs:"sex"`
	Description string `json:"description" structs:"description"`
	Zodiac      string `json:"zodiac" structs:"zodiac"`
	Job         string `json:"job" structs:"job"`
	Education   string `json:"education" structs:"education"`
}

type HashtagInp struct {
	Hashtag []string `json:"hashtag" structs:"hashtag"`
}
