package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type AppointmentService struct {
	appointmentRepository interfaces.AppointmentRepository
	userRepo              interfaces.UserRepository
	expertRepo            interfaces.ExpertRepository
}

func NewAppointmentService(appointmentRepository interfaces.AppointmentRepository, userRepo interfaces.UserRepository, expertRepo interfaces.ExpertRepository) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
		userRepo:              userRepo,
		expertRepo:            expertRepo,
	}
}

func (s *AppointmentService) CreateAppointment(timeSlotId string, userId string) (*response.CreateAppointmentResponse, error) {
	timeSlot, err := s.appointmentRepository.GetTimeSlotById(timeSlotId)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	expert, err := s.expertRepo.GetExpertById(timeSlot.ExpertId)
	if err != nil {
		return nil, err
	}

	if user.XP < expert.XpRequired {
		return nil, fmt.Errorf("user does not have enough XP to book this appointment")
	}
	if !timeSlot.Available {
		return nil, fmt.Errorf("time slot is not available")
	}
	err = s.userRepo.UpdateUserXP(userId, (expert.XpRequired * (-1)))
	if err != nil {
		return nil, err
	}
	err = s.appointmentRepository.UpdateTimeSlotAvailability(timeSlotId, false)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile("config\\google-cloud-oauth2-credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}
	client := utils.GetClient(config)

	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
		return nil, err
	}

	event := &calendar.Event{
		Summary:     "Expert Session",
		Description: "Your personalized expert session with one-to-one interaction.",
		Start: &calendar.EventDateTime{
			DateTime: timeSlot.StartTime.Format(time.RFC3339),
			TimeZone: "Asia/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: timeSlot.EndTime.Format(time.RFC3339),
			TimeZone: "Asia/Kolkata",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: expert.Email},
			{Email: user.Email},
		},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: fmt.Sprintf("req-%d", time.Now().UnixNano()),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		},
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).ConferenceDataVersion(1).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
		return nil, err
	}

	appointmentKey, err := s.appointmentRepository.CreateAppointment(&model.Appointment{
		TimeSlotId:         timeSlotId,
		UserId:             userId,
		MeetingLink:        event.HangoutLink,
		ExpertID:           timeSlot.ExpertId,
		GoogleCalendarLink: event.HtmlLink,
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateAppointmentResponse{
		AppointmentId:      appointmentKey,
		GoogleCalendarLink: event.HtmlLink,
	}, nil
}

func (s *AppointmentService) GetAllAppointments(userId string) ([]response.GetAllAppointmentsResponse, error) {
	appointments, err := s.appointmentRepository.GetAllAppointments(userId)
	if err != nil {
		return nil, err
	}

	var appointmentResponses []response.GetAllAppointmentsResponse
	for i, appointment := range appointments {
		expert, err := s.expertRepo.GetExpertById(appointment.ExpertID)
		if err != nil {
			return nil, err
		}

		timeSlot, err := s.appointmentRepository.GetTimeSlotById(appointment.TimeSlotId)
		if err != nil {
			return nil, err
		}

		appointmentResponses = append(appointmentResponses, response.GetAllAppointmentsResponse{
			AppointmentId: i,
			ExpertName:    expert.Name,
			ImageUrl:      expert.ImageURL,
			Rating:        expert.Rating,
			Date:          timeSlot.Date,
			StartTime:     timeSlot.StartTime,
			EndTime:       timeSlot.EndTime,
		})
	}

	return appointmentResponses, nil
}

func (s *AppointmentService) GetAppointmentById(appointmentId string) (*response.GetAppointmentByIdResponse, error) {
	appointment, err := s.appointmentRepository.GetAppointmentById(appointmentId)
	if err != nil {
		return nil, err
	}

	expert, err := s.expertRepo.GetExpertById(appointment.ExpertID)
	if err != nil {
		return nil, err
	}

	timslot, err := s.appointmentRepository.GetTimeSlotById(appointment.TimeSlotId)
	if err != nil {
		return nil, err
	}

	return &response.GetAppointmentByIdResponse{
		ExpertName:         expert.Name,
		ImageUrl:           expert.ImageURL,
		Rating:             expert.Rating,
		PatientsTreated:    expert.PatientsTreated,
		Experience:         expert.Experience,
		Bio:                expert.Bio,
		Date:               timslot.Date,
		StartTime:          timslot.StartTime,
		EndTime:            timslot.EndTime,
		GoogleCalendarLink: appointment.GoogleCalendarLink,
		MeetingLink:        appointment.MeetingLink,
	}, nil
}
