package handler

import (
	// "bytes"
	// "io"
	// "slices"

	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

var AllowedFileExtensions = []string{"audio/mp3", "audio/wav", "audio/ogg"}

type QuestionHandler struct {
	questionService interfaces.QuestionService
}

func NewQuestionHandler(questionService interfaces.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (h *QuestionHandler) AddQuestion(c *fiber.Ctx) error {
	var req *request.AddQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Type != model.Objective && req.Type != model.Subjective {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question type",
		})
	}

	if req.QuestionText == "" || req.XP == 0 || req.SectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	if req.Type == model.Objective && (req.CorrectOption < 0 || req.CorrectOption >= len(req.Options) || len(req.Options) < 4 || len(req.Explanation) == 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid objective question",
		})
	}

	if (req.Type == model.Subjective || req.Type == model.Test || req.Type == model.Voice) && (len(req.BestAnswer) == 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subjective question",
		})
	}

	resp, err := h.questionService.AddQuestion(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *QuestionHandler) EvaluateObjectiveAnswer(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *request.EvaluateObjectiveAnswerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.QuestionId == "" || req.SelectedOption<0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.questionService.EvaluateObjectiveAnswer(userId, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *QuestionHandler) EvaluateSubjectiveAnswer(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *request.EvaluateSubjectiveAnswerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.QuestionId == "" || len(req.UserAnswer)==0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.questionService.EvaluateSubjectiveAnswer(userId, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *QuestionHandler) GetHint(c *fiber.Ctx) error {
	questionId := c.Params("questionId")
	if questionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing questionId",
		})
	}

	resp, err := h.questionService.GetHint(questionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// func (h *QuestionHandler) EvaluateVoiceAnswer(c *fiber.Ctx) error {
// 	userId, ok := c.Locals("userId").(string)
// 	if !ok {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid userId",
// 		})
// 	}

// 	var req *request.EvaluateVoiceAnswerReq
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid request body",
// 		})
// 	}

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid file",
// 		})
// 	}

// 	var buf = &bytes.Buffer{}
// 	fileReader, err := file.Open()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to read file",
// 		})
// 	}
// 	defer fileReader.Close()

// 	if _, err := io.Copy(buf, fileReader); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to read file",
// 		})
// 	}

// 	if req.QuestionId == "" || !(slices.Contains(AllowedFileExtensions, file.Header.Get("Content-Type"))) {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Invalid Fields",
// 		})
// 	}

// 	resp, err := h.questionService.EvaluateVoiceAnswer(userId, c, req, buf)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(resp)
// }