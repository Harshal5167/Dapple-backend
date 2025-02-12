package model

type UserAnswerEvalutaion struct {
	Evaluation []Evaluation `json:"evaluation"`
	XPGained   int          `json:"xpGained"`
}

type Evaluation struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
