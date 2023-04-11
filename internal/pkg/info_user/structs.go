package info_user

type InfoStruct struct {
	Name        string `json:"name" structs:"name"`
	City        string `json:"city" structs:"city"`
	Sex         uint   `json:"sex" structs:"sex"`
	Description string `json:"description" structs:"description"`
	Zodiac      string `json:"zodiac" structs:"zodiac"`
	Job         string `json:"job" structs:"job"`
	Education   string `json:"education" structs:"education"`
}

type InfoStructAnswer struct {
	Name        string  `json:"name" structs:"name"`
	City        string  `json:"city" structs:"city"`
	Sex         uint    `json:"sex" structs:"sex"`
	Description string  `json:"description" structs:"description"`
	Zodiac      string  `json:"zodiac" structs:"zodiac"`
	Job         string  `json:"job" structs:"job" structs:"job"`
	Education   string  `json:"education" structs:"education"`
	Age         int     `json:"age" structs:"age"`
	Photos      []Photo `json:"photos" structs:"photos"`
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

type Photo struct {
	PhotoId uint `json:"photoId" structs:"photoId"`
	Avatar  bool `json:"avatar" structs:"avatar"`
}
