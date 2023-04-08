package info_user

type InfoStruct struct {
	Name        string   `json:"name" structs:"name"`
	City        string   `json:"city" structs:"city"`
	Sex         uint     `json:"sex" structs:"sex"`
	Description string   `json:"description" structs:"description"`
	Zodiac      string   `json:"zodiac" structs:"zodiac"`
	Job         string   `json:"job" structs:"job"`
	Education   string   `json:"education" structs:"education"`
	Avatar      string   `json:"avatar" structs:"avatar"`
	Photo       []string `json:"photo" structs:"photo"`
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
