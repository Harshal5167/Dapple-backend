package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

type AuthRepository struct {
	FirebaseApp *firebase.App
}

func NewAuthRepository(firebase *firebase.App) *AuthRepository {
	return &AuthRepository{FirebaseApp: firebase}
}

func (c *AuthRepository) CheckExistingUser(email string, username string) (bool, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Firestore(ctx)
	if err != nil {
		fmt.Printf("Failed to get Firestore client: %v", err)
		return false, err
	}
	defer client.Close()

	query := client.Collection("users").Where("email", "==", email).Where("username", "==", username)
	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		if err.Error() == "iterator: no more items" {
			return false, nil
		}
		fmt.Printf("Failed to iterate documents: %v", err)
		return false, err
	}

	if doc.Exists() {
		return true, nil
	}

	return false, nil
}
