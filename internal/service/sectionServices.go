package service

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/clients/videoEvaluation"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type SectionService struct {
	sectionRepo     interfaces.SectionRepository
	levelRepo       interfaces.LevelRepository
	questionRepo    interfaces.QuestionRepository
	lessonRepo      interfaces.LessonRepository
	questionService interfaces.QuestionService
	lessonService   interfaces.LessonService
	testRepo        interfaces.TestRepository
}

func NewSectionService(
	sectionRepo interfaces.SectionRepository,
	levelRepo interfaces.LevelRepository,
	questionRepo interfaces.QuestionRepository,
	lessonRepo interfaces.LessonRepository,
	questionService interfaces.QuestionService,
	lessonService interfaces.LessonService,
	testRepo interfaces.TestRepository) *SectionService {
	return &SectionService{
		sectionRepo:     sectionRepo,
		levelRepo:       levelRepo,
		questionRepo:    questionRepo,
		lessonRepo:      lessonRepo,
		questionService: questionService,
		lessonService:   lessonService,
		testRepo:        testRepo,
	}
}

func (s *SectionService) AddSection(req *request.AddSectionRequest) (*response.AddSectionResponse, error) {
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

	return &response.AddSectionResponse{
		SectionId: sectionId,
	}, nil
}

func (s *SectionService) GetSectionData(userId string, sectionId string) (*response.SectionData, error) {
	nextSectionId, err := s.sectionRepo.GetNextSectionId(sectionId)
	if err != nil {
		return nil, err
	}

	if nextSectionId == "" {
		return s.GetTestData(sectionId)
	}
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
		if question.Type == model.Objective {
			questionList = append(questionList, map[string]interface{}{
				"questionId": questionId,
				"type":       question.Type,
				"question":   question.QuestionText,
				"options":    question.Options,
				"xp":         question.XP,
				"imageUrl": func() interface{} {
					if question.ImageUrl != "" {
						return question.ImageUrl
					}
					return nil
				}(),
			})
		} else if question.Type == model.Subjective || question.Type == model.Voice || question.Type == model.Test {
			questionList = append(questionList, map[string]interface{}{
				"questionId": questionId,
				"type":       question.Type,
				"question":   question.QuestionText,
				"xp":         question.XP,
				"imageUrl": func() interface{} {
					if question.ImageUrl != "" {
						return question.ImageUrl
					}
					return nil
				}(),
			})
		}
	}

	var lessonList []map[string]interface{}
	for _, lessonId := range lessons {
		lesson, err := s.lessonRepo.GetLessonById(lessonId)
		if err != nil {
			return nil, err
		}
		lessonList = append(lessonList, map[string]interface{}{
			"lessonId": lessonId,
			"title":    lesson.Title,
			"content":  lesson.Content,
			"imageUrl": func() interface{} {
				if lesson.ImageUrl != "" {
					return lesson.ImageUrl
				}
				return nil
			}(),
			"xp": lesson.XP,
		})
	}

	var data []map[string]interface{}
	for i := 0; i < min(config.MaxNoOfLessons/2, len(lessonList)); i++ {
		data = append(data, lessonList[i])
	}
	for i := 0; i < min(config.MaxNoOfQuestions/2, len(questionList)); i++ {
		data = append(data, questionList[i])
	}
	for i := config.MaxNoOfLessons / 2; i < len(lessonList); i++ {
		data = append(data, lessonList[i])
	}
	for i := config.MaxNoOfQuestions / 2; i < len(questionList); i++ {
		data = append(data, questionList[i])
	}

	progress, err := s.sectionRepo.StoreSectionProgress(userId, sectionId)
	if err != nil {
		return nil, err
	}

	return &response.SectionData{
		Data:            data,
		SectionProgress: (*progress),
	}, nil
}

func (s *SectionService) AddCompleteSection(req *request.SectionData, levelId string) error {
	addSectionResponse, err := s.AddSection(&request.AddSectionRequest{
		Name:    req.Name,
		LevelId: levelId,
		TotalXP: req.TotalXP,
	})
	if err != nil {
		return err
	}

	for _, question := range req.Questions {
		_, err := s.questionService.AddQuestion(&request.AddQuestionRequest{
			QuestionText:    question.QuestionText,
			Options:         question.Options,
			Hint:            question.Hint,
			CorrectOption:   question.CorrectOption,
			SectionId:       addSectionResponse.SectionId,
			ImageUrl:        question.ImageUrl,
			Type:            question.Type,
			BestAnswer:      question.BestAnswer,
			Explanation:     question.Explanation,
			XP:              question.XP,
			VoiceEvaluation: question.VoiceEvaluation,
			VideoEvaluation: question.VideoEvaluation,
		})
		if err != nil {
			return err
		}
	}

	for _, lesson := range req.Lessons {
		_, err := s.lessonService.AddLesson(&request.AddLessonRequest{
			Title:     lesson.Title,
			Content:   lesson.Content,
			SectionId: addSectionResponse.SectionId,
			XP:        lesson.XP,
			ImageUrl:  lesson.ImageUrl,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SectionService) UpdateSectionProgress(userId string, lessonId string) error {
	lesson, err := s.lessonRepo.GetLessonById(lessonId)
	if err != nil {
		return err
	}

	sectionId := lesson.SectionId
	xp := lesson.XP

	_, _, err = s.sectionRepo.UpdateSectionProgress(userId, sectionId, xp)
	if err != nil {
		return err
	}
	return nil
}

func (s *SectionService) GetTestData(sectionId string) (*response.SectionData, error) {
	questions, _, err := s.sectionRepo.GetQuestionsAndLessons(sectionId)
	if err != nil {
		return nil, err
	}

	var sessionId string
	sessionId, err = videoEvaluation.StartSession()
	if err != nil {
		return nil, err
	}

	var questionList []map[string]interface{}
	for _, questionId := range questions {
		question, err := s.questionRepo.GetQuestionById(questionId)
		if err != nil {
			return nil, err
		}
		questionList = append(questionList, map[string]interface{}{
			"questionId": questionId,
			"type":       question.Type,
			"question":   question.QuestionText,
			"xp":         question.XP,
			"imageUrl": func() interface{} {
				if question.ImageUrl != "" {
					return question.ImageUrl
				}
				return nil
			}(),
		})
	}

	err = s.testRepo.StoreTestSession(sessionId, sectionId)
	if err != nil {
		return nil, err
	}

	return &response.SectionData{
		Data:      questionList,
		SessionId: sessionId,
	}, nil
}
