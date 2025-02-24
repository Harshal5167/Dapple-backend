package videoEvaluation

import (
	"encoding/json"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(reqBody *request.UploadImage) error {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)

	args := fiber.AcquireArgs()
	args.Set("sessionId", reqBody.SessionId)
	args.Set("questionId", reqBody.QuestionId)
	agent.FileData(
		&fiber.FormFile{
			Fieldname: "file",
			Name:      "image.png",
			Content:   reqBody.Image,
		},
	).MultipartForm(args)

	api := config.ImageModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/upload_frame", api))

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

	var response = map[string]interface{}{}

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response["status"] != "success" {
		return fmt.Errorf("request failed")
	}

	return nil
}
