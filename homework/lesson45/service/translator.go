package service

type TranslatingRequest struct {
	UzWord string `json:"uz_word,omitempty"`
}

type TranslatingResponse struct {
	EnWord string `json:"en_word,omitempty"`
}

type TranslatorService interface {
	Translator(in *TranslatingRequest) (*TranslatingResponse, error)
}
