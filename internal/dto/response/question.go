package response

type AddQuestionResponse struct {
	QuestionId string `json:"questionId"`
}

type GetHintResponse struct {
	Hint string `json:"hint"`
}
