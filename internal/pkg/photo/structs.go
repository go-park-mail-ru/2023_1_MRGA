package photo

type AnswerPhoto struct {
	Status int `json:"status"`
	Body   struct {
		PhotoID uint `json:"photoID"`
	} `json:"body"`
}

type ResponseUploadFile struct {
	Status int `json:"status"`
	Body   struct {
		PathToFile string `json:"pathToFile"`
	} `json:"body"`
}
