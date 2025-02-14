package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type LevelHandler interface {
	AddLevel(c *fiber.Ctx) error
	AddCompleteLevel(c *fiber.Ctx) error
}

type SectionHandler interface {
	AddSection(c *fiber.Ctx) error
	GetSection(c *fiber.Ctx) error
	UpdateSectionProgress(c *fiber.Ctx) error
}

type QuestionHandler interface {
	AddQuestion(c *fiber.Ctx) error
	EvaluateObjectiveAnswer(c *fiber.Ctx) error
	EvaluateSubjectiveAnswer(c *fiber.Ctx) error
	GetHint(c *fiber.Ctx) error
}

type LessonHandler interface {
	AddLesson(c *fiber.Ctx) error
}

type UserCourseHandler interface {
	GetUserCourse(c *fiber.Ctx) error
}
