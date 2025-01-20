package config

import (
	"fmt"
	"context"
  
	firebase "firebase.google.com/go/v4"
	// "firebase.google.com/go/v4/auth"
  
	"google.golang.org/api/option"
  )
  
func InitializeFirebaseApp() (*firebase.App, error) {
  opt := option.WithCredentialsFile("firebase-cred.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
	return nil, fmt.Errorf("error initializing app: %v", err)
  }
  fmt.Println("Firebase app initialized")
  return app, nil
}
  