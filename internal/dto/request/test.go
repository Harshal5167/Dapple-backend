package request

type Event string

const (
	Image Event = "image"
	Text  Event = "Text"
)

type TestData struct {
	QuestionId string `json:"questionId"`
	SessionId  string `json:"sessionId"`
	Event      string `json:"event"`
	ImageUrl   string `json:"imageUrl"`
	Answer     string `json:"answer"`
}
