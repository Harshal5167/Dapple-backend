package response

import "time"

type CreateAppointmentResponse struct {
	GoogleCalendarLink string `json:"googleCalendarLink"`
	AppointmentId      string `json:"appointmentId"`
}

type GetAllAppointmentsResponse struct {
	AppointmentId string    `json:"appointmentId"`
	ExpertName    string    `json:"expertName"`
	ImageUrl      string    `json:"imageUrl,omitempty"`
	Rating        float64   `json:"rating"`
	Date          time.Time `json:"date"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

type GetAppointmentByIdResponse struct {
	ExpertName         string    `json:"expertName"`
	ImageUrl           string    `json:"imageUrl,omitempty"`
	Rating             float64   `json:"rating"`
	PatientsTreated    string    `json:"patientsTreated"`
	Experience         string    `json:"experience"`
	Bio                string    `json:"bio"`
	Date               time.Time `json:"date"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	GoogleCalendarLink string    `json:"googleCalendarLink"`
	MeetingLink        string    `json:"meetingLink"`
}
