package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type ExpertService struct {
	expertRepo interfaces.ExpertRepository
}

func NewExpertService(repo interfaces.ExpertRepository) *ExpertService {
	return &ExpertService{expertRepo: repo}
}

func (s *ExpertService) AddExpert(req *request.AddExpertRequest) (*response.AddExpertResponse, error) {
	expert := &model.Expert{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		Bio:        req.Bio,
		XpRequired: req.XpRequired,
		Rating:     req.Rating,
	}

	expertId, err := s.expertRepo.AddExpert(expert)
	if err != nil {
		return nil, err
	}

	var schedules []model.Schedule
	for _, schedule := range req.Schedule {
		schedules = append(schedules, model.Schedule{
			Date: schedule.Date,
		})
		var timeSlotIds []string
		for _, timeSlot := range schedule.TimeSlots {
			ts := &model.TimeSlot{
				ExpertId:  expertId,
				Date:      schedule.Date,
				StartTime: timeSlot.StartTime,
				EndTime:   timeSlot.EndTime,
				Available: timeSlot.Available,
			}
			timeSlotId, err := s.expertRepo.AddTimeSlot(ts)
			if err != nil {
				return nil, err
			}
			timeSlotIds = append(timeSlotIds, timeSlotId)
		}
		schedules[len(schedules)-1].TimeSlotIds = timeSlotIds
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

	var schedules []request.Schedule
	for _, schedule := range expert.Schedule {
		schedules = append(schedules, request.Schedule{
			Date: schedule.Date,
		})
		for _, timeSlotId := range schedule.TimeSlotIds {
			timeSlot, err := s.expertRepo.GetTimeSlotById(timeSlotId)
			if err != nil {
				return nil, err
			}
			schedules[len(schedules)-1].TimeSlots = append(schedules[len(schedules)-1].TimeSlots, request.TimeSlot{
				ExpertId:   timeSlot.ExpertId,
				TimeSlotId: timeSlotId,
				StartTime:  timeSlot.StartTime,
				EndTime:    timeSlot.EndTime,
				Available:  timeSlot.Available,
			})
		}
	}

	return &response.GetExpertResponse{
		Name:       expert.Name,
		ImageURL:   expert.ImageURL,
		Bio:        expert.Bio,
		Schedule:   schedules,
		XpRequired: expert.XpRequired,
		Rating:     expert.Rating,
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
			ExpertId:   key,
			Name:       expert.Name,
			ImageURL:   expert.ImageURL,
			Bio:        expert.Bio,
			XpRequired: expert.XpRequired,
			Rating:     expert.Rating,
		})
	}

	return resp, nil
}
