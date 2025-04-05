package videoEvaluation

import (
	"encoding/json"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/gofiber/fiber/v2"
)

func QuestionResult(questionId string, sessionId string) (*response.QuestionResult, error) {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)

	api := config.ImageModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/get_question_results?session_id=%s&question_id=%s", api, sessionId, questionId))

	fmt.Println("Request URL:", req.RequestURI())
	if err := agent.Parse(); err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	statusCode, body, err := agent.Bytes()
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err[0])
	}

	fmt.Println("Response Status Code:", statusCode)

	if statusCode != fiber.StatusOK {
		return nil, fmt.Errorf("request failed")
	}

	var response response.QuestionResult
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	fmt.Println("Response:", response)

	return &response, nil
}
