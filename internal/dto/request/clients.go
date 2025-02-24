package request

type UploadImage struct {
	Image      []byte `json:"image"`
	SessionId  string `json:"sessionId"`
	QuestionId string `json:"questionId"`
}
