package recommendation

type Recommendation struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Age         int    `json:"age"`
	Sex         uint   `json:"sex"`
	Description string `json:"description"`
	City        string `json:"city"`
}
