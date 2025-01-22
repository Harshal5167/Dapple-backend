package model

type EvaluationRequest struct {
    Question           string   `json:"question"`
    UserAnswer         string   `json:"user_answer"`
    EvaluationCriteria []string `json:"evaluation_criteria"`
}

type CriteriaEvaluation struct {
    Level    string `json:"level"`
    Feedback string `json:"feedback"`
}

type EvaluationResponse struct {
    Evaluation map[string]string            `json:"evaluation"`
    Feedback   map[string]string            `json:"feedback"`
    Error      string                       `json:"error,omitempty"`
}