package request

import "time"

type AddExpertRequest struct {
	Name            string  `json:"name"`
	ImageURL        string  `json:"imageUrl,omitempty"`
	Bio             string  `json:"bio"`
	XpRequired      int     `json:"xpRequired"`
	Rating          float64 `json:"rating"`
	Experience      string  `json:"experience"`
	PatientsTreated string  `json:"patientsTreated"`
	Email           string  `json:"email"`
}

type Schedule struct {
	Date      time.Time  `json:"date"`
	TimeSlots []TimeSlot `json:"timeSlots"`
}

type TimeSlot struct {
	ExpertId   string    `json:"expertId"`
	TimeSlotId string    `json:"timeSlotId"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Available  bool      `json:"available"`
}
