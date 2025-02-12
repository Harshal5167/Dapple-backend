package repository

import (
	"context"
	firebase "firebase.google.com/go/v4"
)

type AuthRepository struct {
	FirebaseApp *firebase.App
}

func NewAuthRepository(firebase *firebase.App) *AuthRepository {
	return &AuthRepository{FirebaseApp: firebase}
}

func (c *AuthRepository) VerifyFirebaseToken(token string) (bool, string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Auth(ctx)
	if err != nil {
		return false, "", err
	}

	verifiedToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return false, "", err
	}

	return true, verifiedToken.Claims["email"].(string), nil
}
