package service

import (
	"time"

	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type ExpertService struct {
	expertRepo      interfaces.ExpertRepository
	appointmentRepo interfaces.AppointmentRepository
}

func NewExpertService(expertRepo interfaces.ExpertRepository, appointmentRepo interfaces.AppointmentRepository) *ExpertService {
	return &ExpertService{
		expertRepo:      expertRepo,
		appointmentRepo: appointmentRepo,
	}
}

func (s *ExpertService) AddExpert(req *request.AddExpertRequest) (*response.AddExpertResponse, error) {
	expert := &model.Expert{
		Name:            req.Name,
		ImageURL:        req.ImageURL,
		Bio:             req.Bio,
		XpRequired:      req.XpRequired,
		Rating:          req.Rating,
		Email:           req.Email,
		Experience:      req.Experience,
		PatientsTreated: req.PatientsTreated,
	}

	expertId, err := s.expertRepo.AddExpert(expert)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Kolkata")
	today := time.Now().In(loc)
	var schedules []model.Schedule

	slotStartHours := []int{10, 11, 12, 13, 14, 15, 16, 17}
	slotStartMinutes := []int{30, 30, 30, 30, 30, 30, 30, 30}

	for i := 0; i < 30; i++ {
		currentDate := today.AddDate(0, 0, i)
		dateOnly := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, loc)

		var timeSlotIds []string
		for j := 0; j < 8; j++ {
			start := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), slotStartHours[j], slotStartMinutes[j], 0, 0, loc)
			end := start.Add(1 * time.Hour)

			ts := &model.TimeSlot{
				ExpertId:  expertId,
				Date:      dateOnly,
				StartTime: start,
				EndTime:   end,
				Available: true,
			}

			timeSlotId, err := s.appointmentRepo.AddTimeSlot(ts)
			if err != nil {
				return nil, err
			}
			timeSlotIds = append(timeSlotIds, timeSlotId)
		}

		schedules = append(schedules, model.Schedule{
			Date:        dateOnly,
			TimeSlotIds: timeSlotIds,
		})
	}

	err = s.expertRepo.UpdateExpert(expertId, schedules)
	if err != nil {
		return nil, err
	}

	return &response.AddExpertResponse{
		ExpertId: expertId,
	}, nil
}

func (s *ExpertService) GetExpertById(expertId string) (*response.GetExpertResponse, error) {
	expert, err := s.expertRepo.GetExpertById(expertId)
	if err != nil {
		return nil, err
	}

	timeSlots, err := s.appointmentRepo.GetTimeSlotsByExpertId(expertId)
	if err != nil {
		return nil, err
	}

	loc := time.Now().Location()
	now := time.Now().In(loc)
	todayDateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// Create a lookup map for timeSlots
	slotMap := make(map[string]*model.TimeSlot)
	for _, ts := range timeSlots {
		slotMap[ts.Id] = ts
	}

	var schedules []request.Schedule

	for _, schedule := range expert.Schedule {
		// Skip past dates
		scheduleDate := schedule.Date.In(loc)
		scheduleDateOnly := time.Date(scheduleDate.Year(), scheduleDate.Month(), scheduleDate.Day(), 0, 0, 0, 0, loc)
		if scheduleDateOnly.Before(todayDateOnly) {
			continue
		}

		s := request.Schedule{
			Date: schedule.Date,
		}

		for _, tsId := range schedule.TimeSlotIds {
			ts, ok := slotMap[tsId]
			if !ok {
				continue
			}

			available := ts.Available

			// Mark time slots as unavailable if they are in the past today
			if scheduleDateOnly.Equal(todayDateOnly) && ts.StartTime.Before(now) {
				available = false
			}

			s.TimeSlots = append(s.TimeSlots, request.TimeSlot{
				TimeSlotId: tsId,
				StartTime:  ts.StartTime,
				EndTime:    ts.EndTime,
				Available:  available,
			})
		}

		// Only add if there are valid timeslots remaining
		if len(s.TimeSlots) > 0 {
			schedules = append(schedules, s)
		}
	}

	return &response.GetExpertResponse{
		ExpertId:        expertId,
		Name:            expert.Name,
		ImageURL:        expert.ImageURL,
		Bio:             expert.Bio,
		XpRequired:      expert.XpRequired,
		Rating:          expert.Rating,
		Experience:      expert.Experience,
		PatientsTreated: expert.PatientsTreated,
		Schedule:        schedules,
	}, nil
}

func (s *ExpertService) GetAllExperts() ([]*response.GetExpertResponse, error) {
	experts, err := s.expertRepo.GetAllExperts()
	if err != nil {
		return nil, err
	}

	var resp []*response.GetExpertResponse
	for key, expert := range experts {
		resp = append(resp, &response.GetExpertResponse{
			ExpertId:        key,
			Name:            expert.Name,
			ImageURL:        expert.ImageURL,
			Bio:             expert.Bio,
			XpRequired:      expert.XpRequired,
			Rating:          expert.Rating,
			Experience:      expert.Experience,
			PatientsTreated: expert.PatientsTreated,
		})
	}

	return resp, nil
}

func (s *ExpertService) GetExpertSchedule(expertId string) (*response.GetExpertScheduleResponse, error) {
	expert, err := s.expertRepo.GetExpertById(expertId)
	if err != nil {
		return nil, err
	}

	timeSlots, err := s.appointmentRepo.GetTimeSlotsByExpertId(expertId)
	if err != nil {
		return nil, err
	}

	loc := time.Now().Location()
	now := time.Now().In(loc)
	todayDateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// Create a lookup map for timeSlots
	slotMap := make(map[string]*model.TimeSlot)
	for _, ts := range timeSlots {
		slotMap[ts.Id] = ts
	}

	var schedules []request.Schedule

	for _, schedule := range expert.Schedule {
		// Skip past dates
		scheduleDate := schedule.Date.In(loc)
		scheduleDateOnly := time.Date(scheduleDate.Year(), scheduleDate.Month(), scheduleDate.Day(), 0, 0, 0, 0, loc)
		if scheduleDateOnly.Before(todayDateOnly) {
			continue
		}

		s := request.Schedule{
			Date: schedule.Date,
		}

		for _, tsId := range schedule.TimeSlotIds {
			ts, ok := slotMap[tsId]
			if !ok {
				continue
			}

			available := ts.Available

			// Mark time slots as unavailable if they are in the past today
			if scheduleDateOnly.Equal(todayDateOnly) && ts.StartTime.Before(now) {
				available = false
			}

			s.TimeSlots = append(s.TimeSlots, request.TimeSlot{
				TimeSlotId: tsId,
				StartTime:  ts.StartTime,
				EndTime:    ts.EndTime,
				Available:  available,
			})
		}

		// Only add if there are valid timeslots remaining
		if len(s.TimeSlots) > 0 {
			schedules = append(schedules, s)
		}
	}

	return &response.GetExpertScheduleResponse{
		Schedule: schedules,
	}, nil
}
