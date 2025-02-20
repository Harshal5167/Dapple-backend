package clients

import (
	"encoding/json"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/gofiber/fiber/v2"
)

func VoiceEvaluation(buf []byte) (*response.VoiceEvaluation, error) {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)

	agent.FileData(&fiber.FormFile{
		Fieldname: "file",
		Name:      "audio.wav",
		Content:   buf,
	}).MultipartForm(nil)

	api := config.VoiceModelAPI
	agent.Request().SetRequestURI(fmt.Sprintf("%s/analyze", api))

	if err := agent.Parse(); err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	statusCode, body, err := agent.Bytes()
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err[0])
	}

	if statusCode != fiber.StatusOK {
		return nil, fmt.Errorf("request failed")
	}

	var voiceEvaluation response.VoiceEvaluation
	if err := json.Unmarshal(body, &voiceEvaluation); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if voiceEvaluation.Status != "success" {
		return nil, fmt.Errorf("voice evaluation failed")
	}

	fmt.Println(voiceEvaluation.AudioFeatures)
	fmt.Println(voiceEvaluation.Status)
	fmt.Println(voiceEvaluation.Top3Emotions)
	fmt.Println(voiceEvaluation.AudioFeatures.VolumeMean)
	fmt.Println(voiceEvaluation.AudioFeatures.Tempo)

	return &voiceEvaluation, nil
}
