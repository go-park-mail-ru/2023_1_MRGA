package photo

type AnswerPhoto struct {
	Status int
	Error string
	Body   struct {
		PhotoID uint
	}
}

type ResponseUploadFile struct {
	Status int
	Error string 
	Body   struct {
		PathToFile string
	}
}
