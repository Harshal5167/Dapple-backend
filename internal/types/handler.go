package types

import "github.com/Harshal5167/Dapple-backend/internal/interfaces"

type Handler struct {
	AuthHandler     interfaces.AuthHandler
	GeminiHandler   interfaces.GeminiHandler
	LevelHandler    interfaces.LevelHandler
	SectionHandler  interfaces.SectionHandler
	QuestionHandler interfaces.QuestionHandler
	LessonHandler   interfaces.LessonHandler
	UserCourseHandler interfaces.UserCourseHandler
}
