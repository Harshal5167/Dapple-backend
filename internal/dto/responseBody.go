package dto

type AuthResponse struct {
	Token     string `json:"token"`
	FirstName string `json:"firstName"`
	Level     int    `json:"level"`
	Section   int    `json:"section"`
}

type EvaluationResponse struct {
	Evaluation map[string]string `json:"evaluation"`
	Feedback   map[string]string `json:"feedback"`
	Error      string            `json:"error,omitempty"`
}