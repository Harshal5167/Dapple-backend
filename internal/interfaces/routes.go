package interfaces

import "github.com/gofiber/fiber/v2"

type AuthRoutes interface {
	AuthRoutes(app *fiber.App)
}

type LevelRoutes interface {
	LevelRoutes(app *fiber.App)
}

type SectionRoutes interface {
	SectionRoutes(app *fiber.App)
}

type QuestionRoutes interface {
	QuestionRoutes(app *fiber.App)
}

type LessonRoutes interface {
	LessonRoutes(app *fiber.App)
}

type UserCourseRoutes interface {
	UserCourseRoutes(app *fiber.App)
}

type UserRoutes interface {
	UserRoutes(app *fiber.App)
}

type TestRoutes interface {
	TestRoutes(app *fiber.App)
}
