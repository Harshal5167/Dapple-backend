package videoEvaluation

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/gofiber/fiber/v2"
)

func EndSession(sessionId string) error {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)

	api := config.ImageModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/end_session?session_id=%s", api, sessionId))

	if err := agent.Parse(); err != nil {
		return fmt.Errorf("failed to parse request: %w", err)
	}

	statusCode, _, err := agent.Bytes()
	if err != nil {
		return fmt.Errorf("request failed: %w", err[0])
	}

	if statusCode != fiber.StatusOK {
		return fmt.Errorf("request failed")
	}

	// var response map[string]string
	// if err := json.Unmarshal(body, &response); err != nil {
	// 	return fmt.Errorf("failed to unmarshal response: %w", err)
	// }

	// if response["status"] != "success" {
	// 	return fmt.Errorf("failed to clear session")
	// }

	return nil
}
