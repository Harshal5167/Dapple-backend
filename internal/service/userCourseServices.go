package service

import (
	"github.com/Harshal5167/Dapple-backend/data"
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseService struct {
	userCourseRepo interfaces.UserCourseRepository
	levelRepo      interfaces.LevelRepository
	geminiService  interfaces.GeminiService
}

func NewUserCourseService(repo interfaces.UserCourseRepository, levelRepo interfaces.LevelRepository, geminiService interfaces.GeminiService) interfaces.UserCourseService {
	return &UserCourseService{
		userCourseRepo: repo,
		levelRepo:      levelRepo,
		geminiService:  geminiService,
	}
}

func (s *UserCourseService) TailorUserCourse(userId string, user model.User) error {
	levels, err := s.levelRepo.GetAllLevels()
	if err != nil {
		return err
	}

	var levelDetails []map[string]string
	for _, level := range levels {
		for key, value := range level {
			levelDetails = append(levelDetails, map[string]string{
				"levelId":     key,
				"levelName":   value.Name,
				"desctiption": value.Description,
			})
		}
	}

	levelsForUser, err := s.geminiService.GenerateUserCourse(user, levelDetails)
	if err != nil {
		return err
	}

	if err := s.userCourseRepo.AddUserCourse(userId, levelsForUser); err != nil {
		return err
	}
	return nil
}

func (s *UserCourseService) GetUserCourse(userId string) (*dto.UserCourseResponse, error) {
	userCourse, err := s.userCourseRepo.GetUserCourse(userId)
	if err != nil {
		return nil, err
	}

	var levels []model.Level
	for _, levelId := range userCourse.Levels {
		level, err := s.levelRepo.GetLevelById(levelId)
		if err != nil {
			return nil, err
		}
		levels = append(levels, *level)
	}
	
	return &dto.UserCourseResponse{
		Levels: levels,
		SectionData: data.SectionData,
		UserProgess: userCourse.UserProgress,
	}, nil
}