package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/config"
)

type LessonService struct {
	lessonRepo  interfaces.LessonRepository
	sectionRepo interfaces.SectionRepository
}

func NewLessonService(lessonRepo interfaces.LessonRepository, sectionRepo interfaces.SectionRepository) *LessonService {
	return &LessonService{
		lessonRepo:  lessonRepo,
		sectionRepo: sectionRepo,
	}
}

func (s *LessonService) AddLesson(req *request.AddLessonRequest) (*response.AddLessonResponse, error) {
	noOfLessons, err := s.sectionRepo.GetNoOfItems(req.SectionId, "lessons")
	if err != nil {
		return nil, err
	}

	if noOfLessons >= config.MaxNoOfLessons {
		return nil, fmt.Errorf("cannot add more lessons to a section")
	}

	lessonId, err := s.lessonRepo.AddLesson(model.Lesson{
		Title:     req.Title,
		Content:   req.Content,
		SectionId: req.SectionId,
		ImageUrl:  req.ImageUrl,
		XP:        req.XP,
	})
	if err != nil {
		return nil, err
	}

	err = s.sectionRepo.AddLessonToSection(req.SectionId, lessonId)
	if err != nil {
		return nil, err
	}

	return &response.AddLessonResponse{
		LessonId: lessonId,
	}, nil
}
