package request

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddCompleteLevelRequest struct {
	Level    model.Level         `json:"level"`
	Sections []model.SectionData `json:"sections"`
}

type AddLevelRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl,omitempty"`
}