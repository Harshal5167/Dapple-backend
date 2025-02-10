package model

type Level struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageUrl,omitempty"`
	Sections    []string `json:"sections,omitempty"`
}
