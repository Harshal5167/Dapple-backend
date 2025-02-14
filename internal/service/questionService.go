package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

var MaxNoOfQuestions int = 4

type QuestionService struct {
	questionRepo      interfaces.QuestionRepository
	sectionRepo       interfaces.SectionRepository
	geminiService     interfaces.GeminiService
	userRepo          interfaces.UserRepository
	UserCourseService interfaces.UserCourseService
}

func NewQuestionService(
	questionRepo interfaces.QuestionRepository,
	sectionRepo interfaces.SectionRepository,
	geminiService interfaces.GeminiService,
	userRepo interfaces.UserRepository,
	UserCourseService interfaces.UserCourseService) *QuestionService {
	return &QuestionService{
		questionRepo:      questionRepo,
		sectionRepo:       sectionRepo,
		geminiService:     geminiService,
		userRepo:          userRepo,
		UserCourseService: UserCourseService,
	}
}

func (s *QuestionService) AddQuestion(req *request.AddQuestionRequest) (*response.AddQuestionResponse, error) {
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
		Hint:          req.Hint,
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

	return &response.AddQuestionResponse{
		QuestionId: questionId,
	}, nil
}

func (s *QuestionService) EvaluateObjectiveAnswer(userId string, req *request.EvaluateObjectiveAnswerReq) (*response.EvaluateObjectiveAnswerResponse, error) {
	question, err := s.questionRepo.GetQuestionById(req.QuestionId)
	if err != nil {
		return nil, err
	}

	xp := 0
	if req.SelectedOption == question.CorrectOption {
		xp = question.XP
	}
	progress, _, err := s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, xp)
	if err != nil {
		return nil, err
	}

	if int(progress) >= MaxNoOfLessons+MaxNoOfQuestions {
		err = s.UserCourseService.UpdateUserProgress(userId, question.SectionId, xp)
		if err != nil {
			return nil, err
		}
	}

	return &response.EvaluateObjectiveAnswerResponse{
		CorrectOption: question.CorrectOption,
		Explanation:   question.Explanation,
		XP:            xp,
	}, nil
}

func (s *QuestionService) EvaluateSubjectiveAnswer(userId string, req *request.EvaluateSubjectiveAnswerReq) (*response.EvaluateSubjectiveAnswerResponse, error) {
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

	progress, xp, err := s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, userAnswerEvaluation.XPGained)
	if err != nil {
		return nil, err
	}

	if int(progress) >= MaxNoOfLessons+MaxNoOfQuestions {
		err = s.UserCourseService.UpdateUserProgress(userId, question.SectionId, xp)
		if err != nil {
			return nil, err
		}
	}

	return &response.EvaluateSubjectiveAnswerResponse{
		Evaluation: userAnswerEvaluation.Evaluation,
		BestAnswer: question.BestAnswer,
		UserAnswer: req.UserAnswer,
		XP:         userAnswerEvaluation.XPGained,
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
