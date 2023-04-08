package recommendation

import "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"

type Recommendation struct {
	Username    string        `json:"username"`
	Avatar      string        `json:"avatar"`
	Age         int           `json:"age"`
	Sex         constform.Sex `json:"sex"`
	Description string        `json:"description"`
	City        string        `json:"city"`
}

type FilterInput struct {
	MinAge    int      `json:"minAge"`
	MaxAge    int      `json:"maxAge"`
	SearchSex uint     `json:"sex"`
	Reason    []string `json:"reason"`
}
