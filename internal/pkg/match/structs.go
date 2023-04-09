package match

type UserRes struct {
	Email  string `json:"email" structs:"email"`
	Name   string `json:"name" structs:"name"`
	Avatar string `json:"avatar" structs:"avatar"`
	Shown  bool   `json:"shown" structs:"shown"`
}

type ReactionInp struct {
	Email    string `json:"email"`
	Reaction string `json:"reaction"`
}

type ChatAnswer struct {
	Email  string `json:"email" structs:"email"`
	Name   string `json:"name" structs:"name"`
	Avatar string `json:"avatar" structs:"avatar"`
}
