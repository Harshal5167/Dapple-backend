package dto

type RegisterRequest struct {
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	FirebaseToken string `json:"firebaseToken"`
}

type LoginRequest struct {
	Email         string `json:"email"`
	FirebaseToken string `json:"firebaseToken"`
}

type EvaluationRequest struct {
	Question           string   `json:"question"`
	UserAnswer         string   `json:"user_answer"`
	EvaluationCriteria []string `json:"evaluation_criteria"`
}

type AddLevelRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageUrl,omitempty"`
}
