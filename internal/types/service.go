package types

import "github.com/Harshal5167/Dapple-backend/internal/interfaces"

type Services struct {
	AuthService       interfaces.AuthService
	GeminiService     interfaces.GeminiService
	LevelService      interfaces.LevelService
	SectionService    interfaces.SectionService
	QuestionService   interfaces.QuestionService
	LessonService     interfaces.LessonService
	UserCourseService interfaces.UserCourseService
}
