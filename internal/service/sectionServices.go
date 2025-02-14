package service

import (
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
}

func NewSectionService(
	sectionRepo interfaces.SectionRepository,
	levelRepo interfaces.LevelRepository,
	questionRepo interfaces.QuestionRepository,
	lessonRepo interfaces.LessonRepository,
	questionService interfaces.QuestionService,
	lessonService interfaces.LessonService) *SectionService {
	return &SectionService{
		sectionRepo:     sectionRepo,
		levelRepo:       levelRepo,
		questionRepo:    questionRepo,
		lessonRepo:      lessonRepo,
		questionService: questionService,
		lessonService:   lessonService,
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
		questionList = append(questionList, map[string]interface{}{
			"questionId":   questionId,
			"type":         question.Type,
			"questionText": question.QuestionText,
			"imageUrl":     question.ImageUrl,
			"XP":           question.XP,
		})
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
			"imageUrl": lesson.ImageUrl,
			"XP":       lesson.XP,
		})
	}

	var data []map[string]interface{}
	for i := 0; i < min(MaxNoOfLessons/2, len(lessonList)); i++ {
		data = append(data, lessonList[i])
	}
	for i := 0; i < min(MaxNoOfQuestions/2, len(questionList)); i++ {
		data = append(data, questionList[i])
	}
	for i := MaxNoOfLessons / 2; i < len(lessonList); i++ {
		data = append(data, lessonList[i])
	}
	for i := MaxNoOfQuestions / 2; i < len(questionList); i++ {
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

func (s *SectionService) AddCompleteSection(req *model.SectionData, levelId string) error {
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
			QuestionText:  question.QuestionText,
			Options:       question.Options,
			Hint:          question.Hint,
			CorrectOption: question.CorrectOption,
			SectionId:     addSectionResponse.SectionId,
			ImageUrl:      question.ImageUrl,
			Type:          question.Type,
			BestAnswer:    question.BestAnswer,
			Explanation:   question.Explanation,
			XP:            question.XP,
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

	_, err = s.sectionRepo.UpdateSectionProgress(userId, sectionId, xp)
	if err != nil {
		return err
	}
	return nil
}
