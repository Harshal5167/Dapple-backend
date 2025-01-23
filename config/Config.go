package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseApp *firebase.App
	GeminiModel *genai.Client
}

func NewConfig(firebaseApp *firebase.App) *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// gemini key chaiye
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY is required in .env file")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		panic(fmt.Sprintf("Failed to create Gemini client: %v", err))
	}

	return &Config{
		FirebaseApp: firebaseApp,
		GeminiModel: client,
	}
}

func InitializeFirebaseApp() (*firebase.App, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	opt := option.WithCredentialsFile("config\\firebase-cred.json")
	config := &firebase.Config{
		ProjectID:   os.Getenv("PROJECT_ID"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
