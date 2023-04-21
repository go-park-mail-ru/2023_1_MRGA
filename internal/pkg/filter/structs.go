package filter

type FilterInput struct {
	MinAge    int      `json:"minAge" structs:"minAge"`
	MaxAge    int      `json:"maxAge" structs:"maxAge"`
	SearchSex uint     `json:"sex" structs:"searchSex"`
	Reason    []string `json:"reason" structs:"reason"`
}
