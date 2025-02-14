package repository

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

type AuthRepository struct {
	firebaseAuth *auth.Client
}

func NewAuthRepository(auth *auth.Client) *AuthRepository {
	return &AuthRepository{
		firebaseAuth: auth,
	}
}

func (c *AuthRepository) VerifyFirebaseToken(token string) (bool, string, error) {
	verifiedToken, err := c.firebaseAuth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return false, "", err
	}

	return true, verifiedToken.Claims["email"].(string), nil
}
