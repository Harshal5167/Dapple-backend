package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type QuestionService struct {
	questionRepo interfaces.QuestionRepository
	sectionRepo  interfaces.SectionRepository
}

func NewQuestionService(questionRepo interfaces.QuestionRepository, sectionRepo interfaces.SectionRepository) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
		sectionRepo:  sectionRepo,
	}
}

func (s *QuestionService) AddQuestion(req *dto.AddQuestionRequest) (*dto.AddQuestionResponse, error) {
	questionId, err := s.questionRepo.AddQuestion(model.Question{
		QuestionText:  req.QuestionText,
		XP:            req.XP,
		Type:          req.Type,
		Options:       req.Options,
		ImageUrl:      req.ImageUrl,
		CorrectOption: req.CorrectOption,
		BestAnswer:    req.BestAnswer,
		SectionId:     req.SectionId,
		Explanation:   req.Explanation,
	})
	if err != nil {
		return nil, err
	}

	err = s.sectionRepo.AddQuestionToSection(req.SectionId, questionId)
	if err != nil {
		return nil, err
	}

	return &dto.AddQuestionResponse{
		QuestionId: questionId,
	}, nil
}
