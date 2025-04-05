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
	args.Set("session_id", reqBody.SessionId)
	args.Set("question_id", reqBody.QuestionId)
	agent.FileData(
		&fiber.FormFile{
			Fieldname: "image",
			Name:      "image.jpeg",
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
		fmt.Println("Status code:", statusCode)
		return fmt.Errorf("request failed")
	}

	var response = map[string]interface{}{}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if response["status"] != "success" {
		fmt.Println("Response status:", response["status"])
		return fmt.Errorf("request failed")
	}

	fmt.Println("success:")
	return nil
}
