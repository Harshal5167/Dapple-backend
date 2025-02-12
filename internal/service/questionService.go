package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

var MaxNoOfQuestions int = 4

type QuestionService struct {
	questionRepo  interfaces.QuestionRepository
	sectionRepo   interfaces.SectionRepository
	geminiService interfaces.GeminiService
	userRepo      interfaces.UserRepository
}

func NewQuestionService(
	questionRepo interfaces.QuestionRepository,
	sectionRepo interfaces.SectionRepository,
	geminiService interfaces.GeminiService,
	userRepo interfaces.UserRepository) *QuestionService {
	return &QuestionService{
		questionRepo:  questionRepo,
		sectionRepo:   sectionRepo,
		geminiService: geminiService,
		userRepo:      userRepo,
	}
}

func (s *QuestionService) AddQuestion(req *dto.AddQuestionRequest) (*dto.AddQuestionResponse, error) {
	noOfQuestions, err := s.sectionRepo.GetNoOfItems(req.SectionId, "questions")
	if err != nil {
		return nil, err
	}
	if noOfQuestions >= MaxNoOfQuestions {
		return nil, fmt.Errorf("cannot add more questions to a section")
	}

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

func (s *QuestionService) EvaluateObjectiveAnswer(userId string, req *dto.EvaluateObjectiveAnswerReq) (*dto.EvaluateObjectiveAnswerResponse, error) {
	question, err := s.questionRepo.GetQuestionById(req.QuestionId)
	if err != nil {
		return nil, err
	}

	xp := 0
	if req.SelectedOption == question.CorrectOption {
		xp = question.XP
	}
	err = s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, xp)
	if err != nil {
		return nil, err
	}

	return &dto.EvaluateObjectiveAnswerResponse{
		CorrectOption: question.CorrectOption,
		Explanation:   question.Explanation,
		XP:            xp,
	}, nil
}

func (s *QuestionService) EvaluateSubjectiveAnswer(userId string, req *dto.EvaluateSubjectiveAnswerReq) (*dto.EvaluateSubjectiveAnswerResponse, error) {
	question, err := s.questionRepo.GetQuestionById(req.QuestionId)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	userAnswerEvaluation, err := s.geminiService.EvaluateUserAnswer(user, question, req.UserAnswer)
	if err != nil {
		return nil, err
	}

	err = s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, userAnswerEvaluation.XPGained)
	if err != nil {
		return nil, err
	}

	return &dto.EvaluateSubjectiveAnswerResponse{
		Evaluation: userAnswerEvaluation.Evaluation,
		BestAnswer: question.BestAnswer,
		XP:         userAnswerEvaluation.XPGained,
	}, nil
}
