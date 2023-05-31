//go:generate easyjson -all structs.go
package photo

type AnswerPhoto struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Body   struct {
		PhotoID uint `json:"photoID"`
	} `json:"body"`
}

type ResponseUploadFile struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Body   struct {
		PathToFile string `json:"pathToFile"`
	} `json:"body"`
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
