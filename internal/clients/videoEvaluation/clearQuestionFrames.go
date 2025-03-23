package videoEvaluation

import (
	"encoding/json"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/gofiber/fiber/v2"
)

func ClearQuestionFrames(sessionId, questionId string) error {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)

	args := fiber.AcquireArgs()
	args.Set("session_id", sessionId)
	args.Set("question_id", questionId)

	api := config.ImageModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/clear_question", api))

	if err := agent.Parse(); err != nil {
		return fmt.Errorf("failed to parse request: %w", err)
	}

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return fmt.Errorf("request failed: %w", errs[0])
	}

	if statusCode != fiber.StatusOK {
		return fmt.Errorf("request failed with status code %d", statusCode)
	}

	var response map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response["status"] != "success" {
		return fmt.Errorf("failed to clear question")
	}

	return nil
}
