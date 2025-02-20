package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type QuestionService struct {
	questionRepo      interfaces.QuestionRepository
	sectionRepo       interfaces.SectionRepository
	geminiService     interfaces.GeminiService
	userRepo          interfaces.UserRepository
	UserCourseService interfaces.UserCourseService
	evaluationRepo    interfaces.EvaluationRepository
}

func NewQuestionService(
	questionRepo interfaces.QuestionRepository,
	sectionRepo interfaces.SectionRepository,
	geminiService interfaces.GeminiService,
	userRepo interfaces.UserRepository,
	UserCourseService interfaces.UserCourseService,
	evaluationRepo interfaces.EvaluationRepository) *QuestionService {
	return &QuestionService{
		questionRepo:      questionRepo,
		sectionRepo:       sectionRepo,
		geminiService:     geminiService,
		userRepo:          userRepo,
		UserCourseService: UserCourseService,
		evaluationRepo:    evaluationRepo,
	}
}

func (s *QuestionService) AddQuestion(req *request.AddQuestionRequest) (*response.AddQuestionResponse, error) {
	noOfQuestions, err := s.sectionRepo.GetNoOfItems(req.SectionId, "questions")
	if err != nil {
		return nil, err
	}
	if noOfQuestions >= config.MaxNoOfQuestions {
		return nil, fmt.Errorf("cannot add more questions to a section")
	}

	var evaluationId string
	if req.Type == "voice" {
		evaluationId, err = s.evaluationRepo.AddVoiceEvaluation(req.VoiceEvaluation)
		if err != nil {
			return nil, err
		}
	}

	question := model.Question{
		QuestionText:  req.QuestionText,
		XP:            req.XP,
		Type:          req.Type,
		Hint:          req.Hint,
		Options:       req.Options,
		ImageUrl:      req.ImageUrl,
		CorrectOption: req.CorrectOption,
		BestAnswer:    req.BestAnswer,
		SectionId:     req.SectionId,
		Explanation:   req.Explanation,
	}

	if req.Type == "voice" {
		question.EvaluationId = evaluationId
	}

	questionId, err := s.questionRepo.AddQuestion(question)
	if err != nil {
		return nil, err
	}

	err = s.sectionRepo.AddQuestionToSection(req.SectionId, questionId)
	if err != nil {
		return nil, err
	}

	return &response.AddQuestionResponse{
		QuestionId: questionId,
	}, nil
}

func (s *QuestionService) GetHint(questionId string) (*response.GetHintResponse, error) {
	hint, err := s.questionRepo.GetHint(questionId)
	if err != nil {
		return nil, err
	}

	return &response.GetHintResponse{
		Hint: hint,
	}, nil
}
