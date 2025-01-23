package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/google/generative-ai-go/genai"
)

type geminiService struct {
	client *genai.Client
}

func NewGeminiService(client *genai.Client) interfaces.GeminiService {
	return &geminiService{
		client: client,
	}
}

func (s *geminiService) EvaluateAnswer(ctx context.Context, req *interfaces.EvaluationRequest) (*interfaces.EvaluationResponse, error) {
	model := s.client.GenerativeModel("gemini-1.0-pro")

	evalExample := "{\n    \"evaluation\": {\n"
	feedbackExample := "    \"feedback\": {\n"

	for i, criteria := range req.EvaluationCriteria {
		evalExample += fmt.Sprintf("        \"%s\": \"High/Moderate/Low\"", criteria)
		feedbackExample += fmt.Sprintf("        \"%s\": \"Your detailed feedback for %s\"", criteria, criteria)

		if i < len(req.EvaluationCriteria)-1 {
			evalExample += ",\n"
			feedbackExample += ",\n"
		} else {
			evalExample += "\n    },\n"
			feedbackExample += "\n    }\n}"
		}
	}

	prompt := fmt.Sprintf(`Evaluate this answer based on the given criteria.
    
Question: %s
User Answer: %s
Evaluation Criteria: %v

Respond only with a JSON object in this exact format:
%s%s`, req.Question, req.UserAnswer, req.EvaluationCriteria, evalExample, feedbackExample)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("generate content error: %v", err)
	}

	if resp == nil || len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("no response generated")
	}

	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response content")
	}

	responseText := ""
	for _, part := range candidate.Content.Parts {
		responseText += fmt.Sprintf("%v", part)
	}

	var response interfaces.EvaluationResponse
	err = json.Unmarshal([]byte(responseText), &response)
	if err != nil {
		return &interfaces.EvaluationResponse{
			Evaluation: make(map[string]string),
			Feedback:   make(map[string]string),
			Error:      fmt.Sprintf("parsing error: %v", err),
		}, nil
	}

	return &response, nil
}
