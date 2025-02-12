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

func (s *geminiService) GenerateUserCourse(user model.User, levelDetails []map[string]string) (*dto.LevelsForUser, error) {
	ctx := context.Background()
	model := s.client.GenerativeModel("gemini-1.0-pro")

	promptTemplate := `You are a neurodiverse expert which focuses on overcoming social interactions anxiety and teach patients social cues.
	You have to design a course for user which will help them to overcome their social anxiety and improve their social skills.
	We are providing you with some available content which has been designed by our team. You have to select the most suitable content for the user based on their profile.
	Please analyze the following user profile and available levels to select the 3 most suitable levels for this user.

	User Profile:
	- FirstName: %s
	- LastName: %s
	- Age: %d
	- Gender: %s
	- Profession : %s
	- SocialChallenges: %s
	- StrugglingSocialSetting: %s

	Analyze this user info and indentify the problems it has then provide him the solution by suggesting 3 levels from the available levels which will be best for him to overcome his fears and problems.
	
	Available Levels: %s
	
	Please select exactly 3 levels that would be most appropriate for this user based on their profile. Consider the following factors:
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
	Note that you have to select exactly 3 levels so the selectedLevels should contains 3 levelIds.
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

	fmt.Println(responseText.String())

	var response dto.LevelsForUser
	err = json.Unmarshal([]byte(responseText.String()), &response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	if len(response.SelectedLevelIds) != 3 {
		return nil, fmt.Errorf("error generating levels for user")
	}

	return &response, nil
}

func (s *geminiService) EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error) {
	ctx := context.Background()
	gemini := s.client.GenerativeModel("gemini-1.0-pro")

	promptTemplate := `You are a neurodiverse expert which focuses on overcoming social interactions anxiety and teach patients social cues.
		I am providing you a user profile and his/her details and a question which he/she has answered. You have to evaluate the user's answer based on the given criteria and provide some feedback to it in the specific format.
	
		User Profile:
		- FirstName: %s
		- LastName: %s
		- Age: %d
		- Gender: %s
		- Profession : %s
		- SocialChallenges: %s
		- StrugglingSocialSetting: %s
	
		Analyze this user info.
		
		Question: %s
		User Answer: %s
		Best Answer: %s
		XP: %d
	
		Evaluation Task:
		- Assess the user's response based on their social challenges and struggling social settings.
		- Compare it to the best answer and identify gaps or areas of improvement.
		- Provide personalized feedback that is constructive, actionable, and encouraging
	
		Analyze the user answer based on his profile and question and create a personalized feedback for him in the below given format.
		You have to provide feedback like key concepts which he/she should focus while answering such questions and key points to answer better after evaluating his response.
		Return your response in the following JSON format:
		{
			"Evaluation":[
				{
					"title": "key concepts to focus while answering",
					"content": "Evaluate how well the user's response aligns with social norms, emotional intelligence, and effective communication strategies. Focus on aspects like active listening, conversational flow, confidence, and non-verbal cues. Ensure the user understands how to engage appropriately in different social settings based on their specific challenges." (string)
				},
				{
					"title": "key points to answer better",
					"content": "Provide actionable improvements tailored to the user's social struggles. Offer guidance on structuring responses more effectively, using open-ended engagement, showing empathy, and maintaining clarity. Highlight specific techniques such as mirroring, tone modulation, or asking follow-up questions to improve interaction quality. Compare the response with the best answer and suggest personalized tweaks for better alignment." (string)
				},
			],
			"xpGained": "based on the user answer give him a xp which you think he should get out of the total xp of that question. also the xp should be in multiple of 10 like 10,20,30,... and should not exceed the total xp of the question and not 0." (int)
		}
		Note that you have to write the two things very nicely in a good way as per the user profile and as per the user answer. you can always refer to the best answer.`

	prompt := fmt.Sprintf(
		promptTemplate,
		user.FirstName,
		user.LastName,
		user.Age,
		user.Gender,
		user.Profession,
		strings.Join(user.SocialChallenges, ", "),
		strings.Join(user.StrugglingSocialSetting, ", "),
		question.QuestionText,
		strings.Join(userAnswer, ", "),
		strings.Join(question.BestAnswer, ", "),
		question.XP)

	resp, err := gemini.GenerateContent(ctx, genai.Text(prompt))
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

	var response *model.UserAnswerEvalutaion
	err = json.Unmarshal([]byte(responseText), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return response, nil
}
