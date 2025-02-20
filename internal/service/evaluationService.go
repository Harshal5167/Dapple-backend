package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/clients"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
)

type EvaluationService struct {
	evaluationRepo    interfaces.EvaluationRepository
	questionRepo      interfaces.QuestionRepository
	geminiService     interfaces.GeminiService
	sectionRepo       interfaces.SectionRepository
	userCourseService interfaces.UserCourseService
	userRepo          interfaces.UserRepository
}

func NewEvaluationService(evaluationRepo interfaces.EvaluationRepository,
	questionRepo interfaces.QuestionRepository,
	geminiService interfaces.GeminiService,
	sectionRepo interfaces.SectionRepository,
	userCourseService interfaces.UserCourseService,
	userRepo interfaces.UserRepository,
) *EvaluationService {
	return &EvaluationService{
		evaluationRepo:    evaluationRepo,
		questionRepo:      questionRepo,
		geminiService:     geminiService,
		sectionRepo:       sectionRepo,
		userCourseService: userCourseService,
		userRepo:          userRepo,
	}
}

func (s *EvaluationService) EvaluateVoiceAnswer(userId string, req *request.EvaluateVoiceAnswerReq, buf []byte) (*response.EvaluateVoiceAnswerResponse, error) {
	voiceEvaluation, err := clients.VoiceEvaluation(buf)
	if err != nil {
		return nil, err
	}

	question, err := s.questionRepo.GetQuestionById(req.QuestionId)
	if err != nil {
		return nil, err
	}

	desiredVoiceEvaluation, err := s.evaluationRepo.GetVoiceEvaluationById(question.EvaluationId)
	if err != nil {
		return nil, err
	}

	xpGained := 0
	for _, emotion := range desiredVoiceEvaluation.Emotions {
		if voiceEvaluation.EmotionDistribution[emotion.Emotion] >= emotion.Confidence {
			xpGained += 40
			break
		}
	}
	if voiceEvaluation.AudioFeatures.SpeechRate >= desiredVoiceEvaluation.SpeechRate.Min && voiceEvaluation.AudioFeatures.SpeechRate <= desiredVoiceEvaluation.SpeechRate.Max {
		xpGained += 20
	}
	if voiceEvaluation.AudioFeatures.VolumeMean >= desiredVoiceEvaluation.VolumeMean.Min && voiceEvaluation.AudioFeatures.VolumeMean <= desiredVoiceEvaluation.VolumeMean.Max {
		xpGained += 20
	}
	if voiceEvaluation.AudioFeatures.SpectralCentroid >= desiredVoiceEvaluation.SpectralCentroid.Min && voiceEvaluation.AudioFeatures.SpectralCentroid <= desiredVoiceEvaluation.SpectralCentroid.Max {
		xpGained += 20
	}
	if voiceEvaluation.AudioFeatures.Tempo >= desiredVoiceEvaluation.Tempo.Min && voiceEvaluation.AudioFeatures.Tempo <= desiredVoiceEvaluation.Tempo.Max {
		xpGained += 20
	}

	formattedResponse, err := s.geminiService.FormatVoiceEvaluationResponse(voiceEvaluation, desiredVoiceEvaluation)
	if err != nil {
		return nil, err
	}

	progress, xp, err := s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, xpGained)
	if err != nil {
		return nil, err
	}

	if int(progress) >= config.MaxNoOfLessons+config.MaxNoOfQuestions {
		err = s.userCourseService.UpdateUserProgress(userId, question.SectionId, xp)
		if err != nil {
			return nil, err
		}
	}

	return &response.EvaluateVoiceAnswerResponse{
		Evaluation: formattedResponse.Evaluation,
		XP:         xpGained,
	}, nil
}

func (s *EvaluationService) EvaluateObjectiveAnswer(userId string, req *request.EvaluateObjectiveAnswerReq) (*response.EvaluateObjectiveAnswerResponse, error) {
	question, err := s.questionRepo.GetQuestionById(req.QuestionId)
	if err != nil {
		return nil, err
	}

	xp := 0
	if req.SelectedOption == (question.CorrectOption - 1) {
		xp = question.XP
	}
	progress, XP, err := s.sectionRepo.UpdateSectionProgress(userId, question.SectionId, xp)
	if err != nil {
		return nil, err
	}

	if int(progress) >= config.MaxNoOfLessons+config.MaxNoOfQuestions {
		err = s.userCourseService.UpdateUserProgress(userId, question.SectionId, XP)
		if err != nil {
			return nil, err
		}
		err = s.sectionRepo.DeleteSectionProgress(userId, question.SectionId)
		if err != nil {
			return nil, err
		}
	}

	return &response.EvaluateObjectiveAnswerResponse{
		CorrectOption: question.CorrectOption - 1,
		Explanation:   question.Explanation,
		XP:            xp,
	}, nil
}

func (s *EvaluationService) EvaluateSubjectiveAnswer(userId string, req *request.EvaluateSubjectiveAnswerReq) (*response.EvaluateSubjectiveAnswerResponse, error) {
	fmt.Println("1")
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

	if int(progress) >= config.MaxNoOfLessons+config.MaxNoOfQuestions {
		err = s.userCourseService.UpdateUserProgress(userId, question.SectionId, xp)
		if err != nil {
			return nil, err
		}
		err = s.sectionRepo.DeleteSectionProgress(userId, question.SectionId)
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
