package request

type Event string

const (
	Image Event = "image"
	Text  Event = "Text"
	Retry Event = "retry"
)

type TestData struct {
	QuestionId string `json:"questionId"`
	SessionId  string `json:"sessionId"`
	ImageUrl   string `json:"imageUrl,omitempty"`
	Answer     string `json:"answer,omitempty"`
}
