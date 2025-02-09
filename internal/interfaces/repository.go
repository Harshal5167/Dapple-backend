package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthRepository interface {
	CreateNewUser(user model.User) (string, error)
	GetUserDetailsFromEmail(email string) (model.User, error)
	VerifyFirebaseToken(token string) (bool, string, error)
}

type LevelRepository interface {
	AddLevel(level model.Level) (string, error)
	AddSectionToLevel(levelId string, sectionId string) error
	GetAllLevels() ([]map[string]model.Level, error)
	GetLevelById(levelId string) (*model.Level, error)
}

type SectionRepository interface {
	AddSection(section model.Section) (string, error)
	AddQuestionToSection(sectionId string, questionId string) error
	AddLessonToSection(sectionId string, lessonId string) error
}

type QuestionRepository interface {
	AddQuestion(question model.Question) (string, error)
}

type LessonRepository interface {
	AddLesson(lesson model.Lesson) (string, error)
}

type UserCourseRepository interface {
	AddUserCourse(userId string, levelsForUser *dto.LevelsForUser) error
	GetUserCourse(userId string) (*model.UserCourse, error)
}
