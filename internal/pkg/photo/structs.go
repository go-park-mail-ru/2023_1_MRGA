package photo

type AnswerPhoto struct {
	Status int `json:"status"`
	Body   struct {
		PhotoID uint `json:"photoID"`
	} `json:"body"`
}
