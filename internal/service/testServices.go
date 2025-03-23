package service

import (
	"encoding/base64"
	"math"
	"time"

	"github.com/Harshal5167/Dapple-backend/internal/clients/videoEvaluation"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
)

type TestService struct {
	geminiService     interfaces.GeminiService
	testRepo          interfaces.TestRepository
	sectionRepo       interfaces.SectionRepository
	questionRepo      interfaces.QuestionRepository
	evaluationRepo    interfaces.EvaluationRepository
	userCourseService interfaces.UserCourseService
}

func NewTestService(geminiService interfaces.GeminiService,
	testRepo interfaces.TestRepository,
	sectionRepo interfaces.SectionRepository,
	questionRepo interfaces.QuestionRepository,
	evaluationRepo interfaces.EvaluationRepository,
	userCourseService interfaces.UserCourseService) *TestService {
	return &TestService{
		geminiService:     geminiService,
		testRepo:          testRepo,
		sectionRepo:       sectionRepo,
		questionRepo:      questionRepo,
		evaluationRepo:    evaluationRepo,
		userCourseService: userCourseService,
	}
}

func (s *TestService) EvaluateTestAnswer(message *request.TestData) (bool, error) {
	question, err := s.questionRepo.GetQuestionById(message.QuestionId)
	if err != nil {
		return false, err
	}

	evaluation, err := s.evaluationRepo.GetVideoEvaluationById(question.EvaluationId)
	if err != nil {
		return false, err
	}

	obtainedEvaluation, err := videoEvaluation.QuestionResult(message.QuestionId, message.SessionId)
	if err != nil {
		return false, err
	}

	testEval, err := s.geminiService.EvaluateTestAnswer(message.Answer, question, obtainedEvaluation, evaluation)
	if err != nil {
		return false, err
	}

	var totalXP int = 0
	if evaluation.Emotion == obtainedEvaluation.AverageEmotion {
		diff := math.Abs(evaluation.Confidence - obtainedEvaluation.AverageConfidence)
		totalXP += (int)(100 - diff)
		totalXP -= totalXP % 10
		totalXP = max(totalXP, 30)
	}

	testEval.UserAnswerXP += totalXP
	err = s.testRepo.StoreQuestionResult(message.SessionId, message.QuestionId, testEval)
	if err != nil {
		return false, err
	}

	questions, _, err := s.sectionRepo.GetQuestionsAndLessons(question.SectionId)
	if err != nil {
		return false, err
	}
	if questions[len(questions)-1] == message.QuestionId {
		err = videoEvaluation.EndSession(message.SessionId)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (s *TestService) EvaluateImageAnswer(message *request.TestData) error {
	image, err := base64.StdEncoding.DecodeString(message.ImageUrl)
	if err != nil {
		return nil
	}

	err = videoEvaluation.UploadImage(&request.UploadImage{
		Image:      image,
		SessionId:  message.SessionId,
		QuestionId: message.QuestionId,
	})
	if err != nil {
		return nil
	}
	return nil
}

func (s *TestService) GetTestResult(userId string, sessionId string, sectionId string) (*response.TestResultResponse, error) {
	testSession, err := s.testRepo.GetTestSession(sessionId, sectionId)
	if err != nil {
		return nil, err
	}

	var testResultResponse = &response.TestResultResponse{}
	testResultResponse.TotalTimeTaken = int(time.Now().Unix() - testSession.Timestamp)

	questionResult, err := s.testRepo.GetAllQuestionResults(sessionId, sectionId)
	if err != nil {
		return nil, err
	}
	testResultResponse.QuestionResult = questionResult

	for _, question := range questionResult {
		testResultResponse.TotalXP += question.UserAnswerXP
	}
	err = s.testRepo.ClearTestSession(sessionId, sectionId)
	if err != nil {
		return nil, err
	}

	err = s.userCourseService.UpdateUserProgress(userId, sectionId, testResultResponse.TotalXP)
	if err != nil {
		return nil, err
	}

	err = videoEvaluation.ClearSession(sessionId)
	if err != nil {
		return nil, err
	}

	return testResultResponse, nil
}

func ( s *TestService) RetryQuestion(sessionId string, questionId string) error {
	err := videoEvaluation.ClearQuestionFrames(sessionId, questionId)
	if err != nil {
		return err
	}
	return nil
}