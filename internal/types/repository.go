package types

import "github.com/Harshal5167/Dapple-backend/internal/interfaces"

type Repositories struct {
	AuthRepo       interfaces.AuthRepository
	LevelRepo      interfaces.LevelRepository
	SectionRepo    interfaces.SectionRepository
	QuestionRepo   interfaces.QuestionRepository
	LessonRepo     interfaces.LessonRepository
	UserCourseRepo interfaces.UserCourseRepository
}
