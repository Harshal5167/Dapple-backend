package model

import "time"

type Expert struct {
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Experience      string     `json:"experience"`
	PatientsTreated string     `json:"patientsTreated"`
	ImageURL        string     `json:"imageUrl,omitempty"`
	Bio             string     `json:"bio"`
	Schedule        []Schedule `json:"schedule,omitempty"`
	XpRequired      int        `json:"xpRequired"`
	Rating          float64    `json:"rating"`
}

type Schedule struct {
	Date        time.Time `json:"date"`
	TimeSlotIds []string  `json:"timeSlotIds"`
}
