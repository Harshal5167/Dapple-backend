package service

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/data"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseService struct {
	userCourseRepo interfaces.UserCourseRepository
	levelRepo      interfaces.LevelRepository
	geminiService  interfaces.GeminiService
	sectionRepo    interfaces.SectionRepository
}

func NewUserCourseService(
	userCourseRepo interfaces.UserCourseRepository,
	levelRepo interfaces.LevelRepository,
	geminiService interfaces.GeminiService,
	sectionRepo interfaces.SectionRepository) interfaces.UserCourseService {
	return &UserCourseService{
		userCourseRepo: userCourseRepo,
		levelRepo:      levelRepo,
		geminiService:  geminiService,
		sectionRepo:    sectionRepo,
	}
}

func (s *UserCourseService) TailorUserCourse(userId string, user model.User) error {
	levels, err := s.levelRepo.GetAllLevels()
	if err != nil {
		return err
	}

	var levelDetails []map[string]string
	fmt.Println(levels)
	for key, level := range levels {
		levelDetails = append(levelDetails, map[string]string{
			"levelId":     key,
			"levelName":   level.Name,
			"description": level.Description,
		})
	}

	fmt.Println(levelDetails)
	levelsForUser, err := s.geminiService.GenerateUserCourse(user, levelDetails)
	if err != nil {
		return err
	}
	fmt.Println(levelsForUser)

	if err := s.userCourseRepo.AddUserCourse(userId, levelsForUser); err != nil {
		return err
	}
	return nil
}

func (s *UserCourseService) GetUserCourse(userId string) (*response.UserCourseResponse, error) {
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

	return &response.UserCourseResponse{
		Levels:      levels,
		SectionData: data.SectionData,
		UserProgess: userCourse.UserProgress,
	}, nil
}

func (s *UserCourseService) UpdateUserProgress(userId string, sectionId string) error {
	nextSectionId, err := s.sectionRepo.GetNextSectionId(sectionId)
	if err != nil {
		return err
	}
	if nextSectionId == "" {
		if err := s.userCourseRepo.UpdateUserProgress(userId, true); err != nil {
			return err
		}
	} else {
		if err := s.userCourseRepo.UpdateUserProgress(userId, false); err != nil {
			return err
		}
	}
	return nil
}
