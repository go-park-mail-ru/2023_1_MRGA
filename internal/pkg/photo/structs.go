//go:generate easyjson -all structs.go
package photo

type AnswerPhoto struct {
	Status int
	Error  string
	Body   struct {
		PhotoID uint
	}
}

type ResponseUploadFile struct {
	Status int
	Error  string
	Body   struct {
		PathToFile string
	}
}

type VoiceResponse struct {
	Qid    string `json:"qid"`
	Result struct {
		Texts []struct {
			Text           string  `json:"text"`
			Confidence     float32 `json:"confidence"`
			PunctuatedText string  `json:"punctuated_text"`
		} `json:"texts"`
		PhraseId string `json:"phrase_id"`
	} `json:"result"`
}
