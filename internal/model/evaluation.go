package model

type UserAnswerEvalutaion struct {
	Evaluation []Evaluation `json:"evaluation"`
	XPGained   int          `json:"xpGained,omitempty"`
}

type Evaluation struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
