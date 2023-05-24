package filter

//easyjson:json
type FilterInput struct {
	MinAge    int      `json:"minAge" structs:"minAge"`
	MaxAge    int      `json:"maxAge" structs:"maxAge"`
	SearchSex uint     `json:"sexSearch" structs:"sexSearch"`
	Reason    []string `json:"reason" structs:"reason"`
}
