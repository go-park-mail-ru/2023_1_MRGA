package recommendation

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"

type Recommendation struct {
	Name        string        `json:"username" structs:"name"`
	Photo       string        `json:"avatar" structs:"name"`
	Age         int           `json:"age" structs:"name"`
	Sex         constform.Sex `json:"sex" structs:"name"`
	Description string        `json:"description"`
	City        string        `json:"city"`
	Hashtags    []string      `json:"hashtags"`
	Zodiac      string        `json:"zodiac"`
	Job         string        `json:"job"`
	Education   string        `json:"education"`
}

type UserRecommend struct {
	UserId uint `json:"userId" structs:"userId"`
}

type FilterInput struct {
	MinAge    int      `json:"minAge" structs:"minAge"`
	MaxAge    int      `json:"maxAge" structs:"maxAge"`
	SearchSex uint     `json:"sex" structs:"searchSex"`
	Reason    []string `json:"reason" structs:"reason"`
}
