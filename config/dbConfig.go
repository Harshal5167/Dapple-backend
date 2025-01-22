package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	// "firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseApp *firebase.App
}

func NewConfig(firebaseApp *firebase.App) *Config {
	return &Config{FirebaseApp: firebaseApp}
}

func InitializeFirebaseApp() (*firebase.App, error) {
	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

	opt := option.WithCredentialsFile("config\\firebase-cred.json")
	config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
