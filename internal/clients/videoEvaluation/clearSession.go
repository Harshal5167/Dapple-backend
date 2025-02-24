package videoEvaluation

import (
	"encoding/json"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/gofiber/fiber/v2"
)

func ClearSession(sessionId string) error {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)
	args := fiber.AcquireArgs()
	args.Set("sessionId", sessionId)

	api := config.ImageModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/clear_session", api))

	if err := agent.Parse(); err != nil {
		return fmt.Errorf("failed to parse request: %w", err)
	}

	statusCode, body, err := agent.Bytes()
	if err != nil {
		return fmt.Errorf("request failed: %w", err[0])
	}

	if statusCode != fiber.StatusOK {
		return fmt.Errorf("request failed")
	}

	var response map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response["status"] != "success" {
		return fmt.Errorf("failed to clear session")
	}

	return nil
}
