package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
)

type AuthService interface {
	Login(reqBody *dto.LoginRequest) (*dto.AuthResponse, error)
	Register(reqBody *dto.RegisterRequest) (*dto.AuthResponse, error)
}

type LevelService interface {
	AddLevel(req *dto.AddLevelRequest) (*dto.AddLevelResponse, error)
}

type SectionService interface {
	AddSection(req *dto.AddSectionRequest) (*dto.AddSectionResponse, error)
}

type QuestionService interface {
	AddQuestion(req *dto.AddQuestionRequest) (*dto.AddQuestionResponse, error)
}

type LessonService interface {
	AddLesson(req *dto.AddLessonRequest) (*dto.AddLessonResponse, error)
}