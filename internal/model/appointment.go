package model

type Appointment struct {
	ExpertID string   `json:"expertId"`
	TimeSlot TimeSlot `json:"timeSlot"`
}
