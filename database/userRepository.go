package database

import (
	"JobTracker/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne(newUser models.User) error {
	// Get MongoDB collection reference
	userCollection := GetCollection("User")

	if userCollection == nil {
		return fmt.Errorf("failed to get MongoDB collection")
	}

	// Use a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert user into MongoDB
	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)       // Optional: Log the error
		return fmt.Errorf("error inserting user: %w", err) // Wrap original error
	}

	return nil
}

func FindOne(username string, password string) (models.User, error) {
	userCollection := GetCollection("User")
	var User models.User
	if userCollection == nil {
		return models.User{}, fmt.Errorf("failed to fetch from collections")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	log.Print("Username", username)
	err := userCollection.FindOne(ctx, filter).Decode(&User)
	if err != nil {
		return models.User{}, fmt.Errorf("%v", err)
	}

	if User.Password != password {
		return models.User{}, fmt.Errorf("invalid password")
	}
	return User, nil
}
