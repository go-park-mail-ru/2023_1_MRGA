package info

type InfoStruct struct {
	Name        string   `json:"name"`
	City        string   `json:"city"`
	Sex         uint     `json:"sex"`
	Description string   `json:"description"`
	Zodiac      string   `json:"zodiac"`
	Job         string   `json:"job"`
	Education   string   `json:"education"`
	Avatar      string   `json:"avatar"`
	Photo       []string `json:"photo"`
}

type InfoChange struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Sex         uint   `json:"sex"`
	Description string `json:"description"`
	Zodiac      string `json:"zodiac"`
	Job         string `json:"job"`
	Education   string `json:"education"`
}
