package response

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
)

type AddExpertResponse struct {
	ExpertId string `json:"expertId"`
}

type GetExpertResponse struct {
	Experience      string             `json:"experience"`
	PatientsTreated string             `json:"patientsTreated"`
	ExpertId        string             `json:"expertId,omitempty"`
	Name            string             `json:"name"`
	ImageURL        string             `json:"imageUrl,omitempty"`
	Bio             string             `json:"bio"`
	Schedule        []request.Schedule `json:"schedule,omitempty"`
	XpRequired      int                `json:"xpRequired"`
	Rating          float64            `json:"rating"`
}

type GetExpertScheduleResponse struct {
	Schedule []request.Schedule `json:"schedule,omitempty"`
}
