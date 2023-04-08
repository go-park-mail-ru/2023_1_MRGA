package recommendation

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"

type Recommendation struct {
	Username    string        `json:"username" structs:"name"`
	Avatar      string        `json:"avatar" structs:"name"`
	Age         int           `json:"age" structs:"name"`
	Sex         constform.Sex `json:"sex" structs:"name"`
	Description string        `json:"description"`
	City        string        `json:"city"`
}

type FilterInput struct {
	MinAge    int      `json:"minAge" structs:"minAge"`
	MaxAge    int      `json:"maxAge" structs:"maxAge"`
	SearchSex uint     `json:"sex" structs:"searchSex"`
	Reason    []string `json:"reason" structs:"reason"`
}
