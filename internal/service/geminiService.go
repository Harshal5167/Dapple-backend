package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
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

func (s *geminiService) GenerateUserCourse(user model.User, levelDetails []map[string]string) (*response.LevelsForUser, error) {
	ctx := context.Background()
	model := s.client.GenerativeModel(config.ModelName)

	promptTemplate := `You are a neurodiverse expert which focuses on overcoming social interactions anxiety and teach patients social cues.
	You have to design a course for user which will help them to overcome their social anxiety and improve their social skills.
	We are providing you with some available content which has been designed by our team. You have to select the most suitable content for the user based on their profile.
	Please analyze the following user profile and available levels to select the 1 most suitable levels for this user.

	User Profile:
	- FirstName: %s
	- LastName: %s
	- Age: %d
	- Gender: %s
	- Profession : %s
	- SocialChallenges: %s
	- StrugglingSocialSetting: %s

	Analyze this user info and indentify the problems it has then provide him the solution by suggesting 1 levels from the available levels which will be best for him to overcome his fears and problems.
	
	Available Levels: %s
	
	Please select exactly 1 levels that would be most appropriate for this user based on their profile. Consider the following factors:
	1. User's age 
	2. Alignment with his profession and gender.
	3. Help him overcome his social challenges.
	4. Help him perform better in his struggling social settings.
	
	Return your response in the following json format:
	` + "`" + "`" + "`" + `json{
		"selectedLevelIds": [
			"levelId"
		]
	}` + "`" + "`" + "`" + `
	Note that you have to select exactly 1 levels so the selectedLevels should contains 1 levelIds.
	Give the response like json format with the response wrapped in json and backticks (just like standard json format).
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
	start := strings.Index(responseText.String(), "```json") + 7
	end := strings.LastIndex(responseText.String(), "```")
	if start == -1 || end == -1 || start >= end {
		return nil, fmt.Errorf("invalid format")
	}
	jsonStr := responseText.String()[start:end]

	var response response.LevelsForUser
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	if len(response.SelectedLevelIds) != 1 {
		return nil, fmt.Errorf("error generating levels for user")
	}

	return &response, nil
}

func (s *geminiService) EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error) {
	ctx := context.Background()
	gemini := s.client.GenerativeModel(config.ModelName)

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
		Don't add his name anywhere in the response that was given to you just for reference.
		Return your response in the following json format:
		` + "`" + "`" + "`" + `json{
			"Evaluation":[
				{
					"title": "Key Concepts to Focus",
					"content": "Evaluate how well the user's response aligns with social norms, emotional intelligence, and effective communication strategies. Focus on aspects like active listening, conversational flow, confidence, and non-verbal cues. Ensure the user understands how to engage appropriately in different social settings based on their specific challenges. this field should not contain more than 25 words"
				},
				{
					"title": "Key Points to Answer Better",
					"content": "Provide actionable improvements tailored to the user's social struggles. Offer guidance on structuring responses more effectively, using open-ended engagement, showing empathy, and maintaining clarity. Highlight specific techniques such as mirroring, tone modulation, or asking follow-up questions to improve interaction quality. Compare the response with the best answer and suggest personalized tweaks for better alignment. this field should not contain more than 25 words"
				}
			],
			"xpGained": "based on the user answer give him a xp which you think he should get out of the total xp of that question. also the xp should be in multiple of 10 like 10,20,30,... give the xp generously and dont be stingy in giving it and should not exceed the total xp of the question and not 0." (int)
		}` + "`" + "`" + "`" + `
		Note that you have to write the two things very nicely in a good way as per the user profile and as per the user answer. you can always refer to the best answer.
		THE CONTENT FIELD SHOULD NOT CONTAIN ANY SPECIAL CHARACTERS LIKE SINGLE OR DOUBLE QUOTES OR NUMBERS IT SHOULD BE PLAIN TEXT NOT IN README FORM NEITHER WITH \N AND \T TYPE THINGS. REMEMBER THIS WHILE GENERATING THE RESPONSE.
		REMEMBER NO DOUBLE QUOTES OR SINGLE QUOTES OR NUMBERS IN THE CONTENT FIELD.
		Give the response like json format with the response wrapped in json and backticks (just like standard json format).
		`

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

	fmt.Println(responseText)
	start := strings.Index(responseText, "```json") + 7
	end := strings.LastIndex(responseText, "```")
	if start == -1 || end == -1 || start >= end {
		fmt.Println("Invalid format")
		return nil, fmt.Errorf("invalid format")
	}
	jsonStr := responseText[start:end]
	fmt.Println(jsonStr)

	response := &model.UserAnswerEvalutaion{}
	err = json.Unmarshal([]byte(jsonStr), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}

	return response, nil
}

func (s *geminiService) FormatVoiceEvaluationResponse(obtainedVoiceEvaluation *response.VoiceEvaluation, desiredVoiceEvaluation *model.VoiceEvaluation) (*model.UserAnswerEvalutaion, error) {
	ctx := context.Background()
	gemini := s.client.GenerativeModel(config.ModelName)

	promptTemplate := `I'm providing you the desired voice evaluation for a question and also the obtained evaluation from the user's voice. You have to format the obtained evaluation and the desired evaluation in a user friendly layman language so that he can understand it easily and can work on it.
	Explain both the things in a very simple language and in a very easy way so that he can understand it easily.

	Obtianed Evaluation
	- Top Emotions identified from his voice: %s
	- Spectral Centroid of voice: %f
	- Tempo of voice: %f 
	- Volume Mean: %f
	- Speech Rate of the voice: %f

	Desired Evaluation output
	- Desired Top Emotions: %s
	- Desired Spectral Centroid: %s
	- Desired Tempo of voice: %s 
	- Desired Volume Mean: %s
	- Desired Speech Rate of the voice: %s

	Dont show him numbers for this things or neither the complex names of this fields, you change and explain him in the most easy way so that he can understand it clearly.
	Return your response in the following json format:
	` + "`" + "`" + "`" + `json{
		"Evaluation":[
			{
				"title": "Your Voice Evaluation",
				"content": "here explain him in easy terms for the obtained evaluation. this field should not contain more than 50 words"
			},
			{
				"title": "Key Points to Answer Better",
				"content": "here explain him in easy terms for the desired evaluation. This field should not contain more than 50 words"
			}
		]
	}` + "`" + "`" + "`" + `
	Note that you have to write the two things very nicely in a good way.
	THE CONTENT FIELD SHOULD NOT CONTAIN ANY SPECIAL CHARACTERS LIKE SINGLE OR DOUBLE QUOTES OR NUMBERS IT SHOULD BE PLAIN TEXT NOT IN README FORM NEITHER WITH \N AND \T TYPE THINGS. REMEMBER THIS WHILE GENERATING THE RESPONSE.
	REMEMBER NO DOUBLE QUOTES OR SINGLE QUOTES OR NUMBERS IN THE CONTENT FIELD.
	Give the response like json format with the response wrapped in json and backticks (just like standard json format).
	The response should be in valid json format look closely for closing brackets.
	`

	obtainedEmotions := utils.BuildStringForEmotions(obtainedVoiceEvaluation.Top3Emotions)
	desiredEmotions := utils.BuildStringForEmotions(desiredVoiceEvaluation.Emotions)

	prompt := fmt.Sprintf(
		promptTemplate,
		obtainedEmotions,
		obtainedVoiceEvaluation.AudioFeatures.SpectralCentroid,
		obtainedVoiceEvaluation.AudioFeatures.Tempo,
		obtainedVoiceEvaluation.AudioFeatures.VolumeMean,
		obtainedVoiceEvaluation.AudioFeatures.SpeechRate,
		desiredEmotions,
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.SpectralCentroid.Max, desiredVoiceEvaluation.SpectralCentroid.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.Tempo.Max, desiredVoiceEvaluation.Tempo.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.VolumeMean.Max, desiredVoiceEvaluation.VolumeMean.Min),
		fmt.Sprintf("max: %f, min: %f", desiredVoiceEvaluation.SpeechRate.Max, desiredVoiceEvaluation.SpeechRate.Min))

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

	fmt.Println(responseText)
	start := strings.Index(responseText, "```json") + 7
	end := strings.LastIndex(responseText, "```")
	if start == -1 || end == -1 || start >= end {
		fmt.Println("Invalid format")
		return nil, fmt.Errorf("invalid format")
	}
	jsonStr := responseText[start:end]

	response := &model.UserAnswerEvalutaion{}
	err = json.Unmarshal([]byte(jsonStr), response)
	if err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return response, nil
}
