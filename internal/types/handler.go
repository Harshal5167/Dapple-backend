package types

import "github.com/Harshal5167/Dapple-backend/internal/interfaces"

type Handler struct {
	AuthHandler       interfaces.AuthHandler
	LevelHandler      interfaces.LevelHandler
	SectionHandler    interfaces.SectionHandler
	QuestionHandler   interfaces.QuestionHandler
	LessonHandler     interfaces.LessonHandler
	UserCourseHandler interfaces.UserCourseHandler
}
