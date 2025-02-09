package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"
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

func (s *geminiService) EvaluateAnswer(req *dto.EvaluationRequest) (*dto.EvaluationResponse, error) {
	ctx := context.Background()
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

	var response dto.EvaluationResponse
	err = json.Unmarshal([]byte(responseText), &response)
	if err != nil {
		return &dto.EvaluationResponse{
			Evaluation: make(map[string]string),
			Feedback:   make(map[string]string),
			Error:      fmt.Sprintf("parsing error: %v", err),
		}, nil
	}

	return &response, nil
}

func (s *geminiService) GenerateUserCourse(user model.User, levelDetails []map[string]string) (*dto.LevelsForUser, error) {
	ctx := context.Background()
	model := s.client.GenerativeModel("gemini-1.0-pro")

	promptTemplate := `You are a neurodiverse expert which focuses on overcoming social interactions anxiety and teach patients social cues.
	You have to design a course for user which will help them to overcome their social anxiety and improve their social skills.
	We are providing you with some available content which has been designed by our team. You have to select the most suitable content for the user based on their profile.
	Please analyze the following user profile and available levels to select the 10 most suitable levels for this user.

	User Profile:
	- FirstName: %s
	- LastName: %s
	- Age: %d
	- Gender: %s
	- Profession : %s
	- SocialChallenges: %s
	- StrugglingSocialSetting: %s

	Analyze this user info and indentify the problems it has then provide him the solution by suggesting 10 levels from the available levels which will be best for him to overcome his fears and problems.
	
	Available Levels: %s
	
	Please select exactly 10 levels that would be most appropriate for this user based on their profile. Consider the following factors:
	1. User's age 
	2. Alignment with his profession and gender.
	3. Help him overcome his social challenges.
	4. Help him perform better in his struggling social settings.
	
	Return your response in the following JSON format:
	{
		"selectedLevelIds": [
			"levelId"
		]
	}
	Note that you have to select exactly 10 levels so the selectedLevels should contains 10 levelIds.
	`

	levelDetailsString := utils.BuildStringForLevels(levelDetails)
	prompt := fmt.Sprintf(
		promptTemplate,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		user.Profession,
		strings.Join(user.SocialChallenges, ", "),
		strings.Join(user.StrugglingSocialSetting, ", "),
		levelDetailsString)

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

	var responseText strings.Builder
	for _, part := range candidate.Content.Parts {
		responseText.WriteString(fmt.Sprintf("%v", part))
	}

	var response dto.LevelsForUser
	err = json.Unmarshal([]byte(responseText.String()), &response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	if len(response.SelectedLevelIds) != 10 {
		return nil, fmt.Errorf("error generating levels for user")
	}

	return &response, nil
}
