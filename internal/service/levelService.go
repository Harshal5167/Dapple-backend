package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LevelService struct {
	levelRepo      interfaces.LevelRepository
	sectionService interfaces.SectionService
}

func NewLevelService(levelRepo interfaces.LevelRepository, sectionService interfaces.SectionService) *LevelService {
	return &LevelService{
		levelRepo:      levelRepo,
		sectionService: sectionService,
	}
}

func (s *LevelService) AddLevel(req *request.AddLevelRequest) (*response.AddLevelResponse, error) {
	levelId, err := s.levelRepo.AddLevel(model.Level{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &response.AddLevelResponse{
		LevelId: levelId,
	}, nil
}

func (s *LevelService) AddCompleteLevel(req *request.AddCompleteLevelRequest) (*response.AddLevelResponse, error) {
	levelId, err := s.levelRepo.AddLevel(model.Level{
		Name:        req.Level.Name,
		Description: req.Level.Description,
		ImageUrl:    req.Level.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	for _, section := range req.Sections {
		err := s.sectionService.AddCompleteSection(&request.SectionData{
			Name:      section.Name,
			TotalXP:   section.TotalXP,
			Questions: section.Questions,
			Lessons:   section.Lessons,
		}, levelId)
		if err != nil {
			return nil, err
		}
	}
	return &response.AddLevelResponse{
		LevelId: levelId,
	}, nil
}
