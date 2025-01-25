package auth

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
	var users map[string]map[string]interface{}

	err = ref.OrderByChild("email").EqualTo(email).Get(ctx, &users)
	if err != nil {
		return model.User{}, fmt.Errorf("error querying database: %v", err)
	}
	for key, userData := range users {
		user := model.User{
			UserId:    key,
			Email:     userData["email"].(string),
			FirstName: userData["firstName"].(string),
			LastName:  userData["lastName"].(string),
			Age:       userData["age"].(int),
			Role:      userData["role"].(string),
			XP:        userData["XP"].(int),
			Badges:    userData["badges"].([]int),
			Level:     userData["levels"].(int),
			Section:   userData["section"].(int),
		}
		return user, nil
	}
	return model.User{}, errors.New("unknown error occurred")
}

func (c *AuthRepository) VerifyFirebaseToken(token string) (bool, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Auth(ctx)
	if err != nil {
		return false, err
	}

	_, err = client.VerifyIDToken(ctx, token)
	if err != nil {
		return false, err
	}
	return true, nil
}
