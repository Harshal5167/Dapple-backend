package model

type TestAnswerEval struct {
	UserAnswerXP  int             `json:"userAnswerXP"`
	Evaluation    []Evaluation    `json:"evaluation"`
	AnswerSummary []AnswerSummary `json:"answerSummary"`
}

type AnswerSummary struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserScore int    `json:"userScore"`
	MaxScore  int    `json:"maxScore"`
}

type TestSession struct {
	TotalXP   int   `json:"totalXP" redis:"totalXP"`
	Timestamp int64 `json:"timestamp" redis:"timestamp"`
}
