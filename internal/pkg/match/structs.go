package match

type UserRes struct {
	UserId uint   `json:"userId" structs:"userId"`
	Name   string `json:"name" structs:"name"`
	Age    int    `json:"age" structs:"age"`
	Photo  uint   `json:"avatar" structs:"avatar"`
	Shown  bool   `json:"shown" structs:"shown"`
}

type ReactionInp struct {
	EvaluatedUserId uint   `json:"evaluatedUserId"`
	Reaction        string `json:"reaction"`
}

const (
	MissedMatch = iota
	NewMatch
	FirstReaction
)

type ReactionResult struct {
	ResultCode int
}

type ChatAnswer struct {
	UserId uint   `json:"userId" structs:"userId"`
	Name   string `json:"name" structs:"name"`
	Photo  uint   `json:"avatar" structs:"avatar"`
}
