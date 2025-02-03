package repository

import (
	"context"
	"errors"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthRepository struct {
	FirebaseApp *firebase.App
}

func NewAuthRepository(firebase *firebase.App) *AuthRepository {
	return &AuthRepository{FirebaseApp: firebase}
}

func (c *AuthRepository) CreateNewUser(user model.User) (string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	var ref *db.Ref
	ref, err = client.NewRef("users").Push(ctx, user)
	userId := ref.Key

	if err != nil {
		fmt.Printf("Failed adding user: %v", err)
		return "", err
	}
	return userId, nil
}

func (c *AuthRepository) GetUserDetailsFromEmail(email string) (model.User, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref := client.NewRef("users")
	var users map[string]model.User

	err = ref.OrderByChild("email").EqualTo(email).LimitToFirst(1).Get(ctx, &users)
	if err != nil {
		return model.User{}, fmt.Errorf("error querying database: %v", err)
	}

	for user := range users {
		userDetails := users[user]
		userDetails.UserId = user
		return userDetails, nil
	}
	return model.User{}, errors.New("unknown error occurred")
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
