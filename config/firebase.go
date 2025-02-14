package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func InitializeFirebaseClient() (*db.Client, *auth.Client) {
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
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Errorf("error initializing auth client: %v", err))
	}

	client, err := app.Database(context.Background())
	if err != nil {
		panic(fmt.Errorf("error initializing database client: %v", err))
	}
	return client, auth
}
