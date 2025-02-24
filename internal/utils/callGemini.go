package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func CallGemini(ctx context.Context, gemini *genai.GenerativeModel, prompt string) (string, error) {
	resp, err := gemini.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("generate content error: %v", err)
	}

	if resp == nil || len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return "", fmt.Errorf("empty response content")
	}

	responseText := ""
	for _, part := range candidate.Content.Parts {
		responseText += fmt.Sprintf("%v", part)
	}

	start := strings.Index(responseText, "```json") + 7
	end := strings.LastIndex(responseText, "```")
	if start == -1 || end == -1 || start >= end {
		return "", fmt.Errorf("invalid format")
	}
	return responseText[start:end], nil
}
