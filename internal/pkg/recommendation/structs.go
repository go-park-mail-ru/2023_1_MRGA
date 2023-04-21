package recommendation

import (
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
)

type Recommendation struct {
	Id          uint          `json:"userId" structs:"userId"`
	Name        string        `json:"name" structs:"name"`
	Photos      []uint        `json:"photos" structs:"photos"`
	Age         int           `json:"age" structs:"age"`
	Sex         constform.Sex `json:"sex" structs:"sex"`
	Description string        `json:"description"`
	City        string        `json:"city"`
	Hashtags    []string      `json:"hashtags"`
	Zodiac      string        `json:"zodiac"`
	Job         string        `json:"job"`
	Education   string        `json:"education"`
}

type DBRecommendation struct {
	Id          uint          `json:"userId" structs:"userId"`
	Name        string        `json:"name" structs:"name"`
	BirthDay    string        `json:"birthDay" structs:"birthDay" sql:"type:date" gorm:"type:date"`
	Sex         constform.Sex `json:"sex" structs:"sex"`
	Description string        `json:"description"`
	City        string        `json:"city"`
	Zodiac      string        `json:"zodiac"`
	Job         string        `json:"job"`
	Education   string        `json:"education"`
}

type UserRecommend struct {
	UserId uint `json:"userId" structs:"userId"`
}
