package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"
	"github.com/google/generative-ai-go/genai"
)

type geminiService struct {
	client *genai.Client
}

func NewGeminiService(client *genai.Client) interfaces.GeminiService {
	return &geminiService{
		client: client,
	}
}

func (s *geminiService) GenerateUserCourse(user model.User, levelDetails []map[string]string) (*response.LevelsForUser, error) {
	model := s.client.GenerativeModel(config.ModelName)

	levelDetailsString := utils.BuildStringForLevels(levelDetails)
	prompt := fmt.Sprintf(
		config.GenerateUserCoursePrompt,
		user.Age,
		user.Gender,
		user.Profession,
		strings.Join(user.SocialChallenges, ", "),
		strings.Join(user.StrugglingSocialSetting, ", "),
		levelDetailsString)

	jsonStr, err := utils.CallGemini(context.Background(), model, prompt)
	if err != nil {
		return nil, fmt.Errorf("error generating levels for user: %v", err)
	}

	var response response.LevelsForUser
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	return &response, nil
}

func (s *geminiService) EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error) {
	gemini := s.client.GenerativeModel(config.ModelName)

	prompt := fmt.Sprintf(
		config.EvaluateUserAnswerPrompt,
		user.Age,
		user.Gender,
		user.Profession,
		strings.Join(user.SocialChallenges, ", "),
		strings.Join(user.StrugglingSocialSetting, ", "),
		question.QuestionText,
		strings.Join(userAnswer, ", "),
		strings.Join(question.BestAnswer, ", "),
		question.XP)

	jsonStr, err := utils.CallGemini(context.Background(), gemini, prompt)
	if err != nil {
		return nil, fmt.Errorf("generate content error: %v", err)
	}

	response := &model.UserAnswerEvalutaion{}
	err = json.Unmarshal([]byte(jsonStr), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	return response, nil
}

func (s *geminiService) FormatVoiceEvaluationResponse(obtainedVoiceEvaluation *response.VoiceEvaluation, desiredVoiceEvaluation *model.VoiceEvaluation) (*model.UserAnswerEvalutaion, error) {
	gemini := s.client.GenerativeModel(config.ModelName)

	obtainedEmotions := utils.BuildStringForEmotions(obtainedVoiceEvaluation.Top3Emotions)
	desiredEmotions := utils.BuildStringForEmotions(desiredVoiceEvaluation.Emotions)

	prompt := fmt.Sprintf(
		config.FormatVoiceEvaluationResponsePrompt,
		obtainedEmotions,
		obtainedVoiceEvaluation.AudioFeatures.SpectralCentroid,
		obtainedVoiceEvaluation.AudioFeatures.Tempo,
		obtainedVoiceEvaluation.AudioFeatures.VolumeMean,
		obtainedVoiceEvaluation.AudioFeatures.SpeechRate,
		desiredEmotions,
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.SpectralCentroid.Max, desiredVoiceEvaluation.SpectralCentroid.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.Tempo.Max, desiredVoiceEvaluation.Tempo.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.VolumeMean.Max, desiredVoiceEvaluation.VolumeMean.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.SpeechRate.Max, desiredVoiceEvaluation.SpeechRate.Min))

	jsonStr, err := utils.CallGemini(context.Background(), gemini, prompt)
	if err != nil {
		return nil, fmt.Errorf("generate content error: %v", err)
	}

	response := &model.UserAnswerEvalutaion{}
	err = json.Unmarshal([]byte(jsonStr), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return response, nil
}

func (s *geminiService) EvaluateTestAnswer(Answer string, question *model.Question, obtainedVideoEvaluation *response.QuestionResult, desiredVideoEvaluation *model.Emotion) (*model.TestAnswerEval, error) {
	gemini := s.client.GenerativeModel(config.ModelName)

	prompt := fmt.Sprintf(
		config.EvaluateTestAnswerPrompt,
		question.QuestionText,
		Answer,
		strings.Join(question.BestAnswer, ", "),
		question.XP/2,
		obtainedVideoEvaluation.ResultSummary.MostCommonEmotion,
		obtainedVideoEvaluation.ResultSummary.EmotionVariability,
		obtainedVideoEvaluation.ResultSummary.OverallTrend,
		strings.Join(obtainedVideoEvaluation.ResultSummary.NotableObservations, ", "),
		desiredVideoEvaluation.Emotion,
		desiredVideoEvaluation.Confidence,
	)

	jsonStr, err := utils.CallGemini(context.Background(), gemini, prompt)
	if err != nil {
		return nil, fmt.Errorf("generate content error: %v", err)
	}

	fmt.Println("Gemini response:", jsonStr)

	response := &model.TestAnswerEval{}
	err = json.Unmarshal([]byte(jsonStr), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return response, nil
}
