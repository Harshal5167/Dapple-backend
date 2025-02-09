package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type SectionService struct {
	sectionRepo  interfaces.SectionRepository
	levelRepo    interfaces.LevelRepository
	questionRepo interfaces.QuestionRepository
	lessonRepo   interfaces.LessonRepository
}

// GetSection implements interfaces.SectionService.
func (s *SectionService) GetSection(sectionId string) (*dto.SectionData, error) {
	panic("unimplemented")
}

func NewSectionService(sectionRepo interfaces.SectionRepository, levelRepo interfaces.LevelRepository, questionRepo interfaces.QuestionRepository, lessonRepo interfaces.LessonRepository) *SectionService {
	return &SectionService{
		sectionRepo:  sectionRepo,
		levelRepo:    levelRepo,
		questionRepo: questionRepo,
		lessonRepo:   lessonRepo,
	}
}

func (s *SectionService) AddSection(req *dto.AddSectionRequest) (*dto.AddSectionResponse, error) {
	sectionId, err := s.sectionRepo.AddSection(model.Section{
		Name:    req.Name,
		LevelId: req.LevelId,
		TotalXP: req.TotalXP,
	})
	if err != nil {
		return nil, err
	}

	err = s.levelRepo.AddSectionToLevel(req.LevelId, sectionId)
	if err != nil {
		return nil, err
	}

	return &dto.AddSectionResponse{
		SectionId: sectionId,
	}, nil
}

func (s *SectionService) GetSectionData(sectionId string) (*dto.SectionData, error) {
	questions, lessons, err := s.sectionRepo.GetQuestionsAndLessons(sectionId)
	if err != nil {
		return nil, err
	}

	var questionList []map[string]interface{}
	for _, questionId := range questions {
		question, err := s.questionRepo.GetQuestionById(questionId)
		if err != nil {
			return nil, err
		}
		questionList = append(questionList, *question)
	}

	var lessonList []map[string]interface{}
	for _, lessonId := range lessons {
		lesson, err := s.lessonRepo.GetLessonById(lessonId)
		if err != nil {
			return nil, err
		}
		lessonList = append(lessonList, *lesson)
	}

	var data []map[string]interface{}
	for i := 0; i < min(2, len(lessonList)); i++ {
		data = append(data, lessonList[i])
	}
	for i := 0; i < min(2, len(questionList)); i++ {
		data = append(data, questionList[i])
	}
	for i := 2; i < len(lessonList); i++ {
		data = append(data, lessonList[i])
	}
	for i := 2; i < len(questionList); i++ {
		data = append(data, questionList[i])
	}

	return &dto.SectionData{
		Data: data,
	}, nil
}
