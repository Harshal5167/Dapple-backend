package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddSectionResponse struct {
	SectionId string `json:"sectionId"`
}

type SectionData struct {
	Data            []map[string]interface{} `json:"data"`
	SessionId       string                   `json:"sessionId,omitempty"`
	SectionProgress model.SectionProgress    `json:"sectionProgress,omitempty"`
}