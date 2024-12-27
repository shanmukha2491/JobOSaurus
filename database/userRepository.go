package database

import (
	"JobTracker/models"
	"context"
	"fmt"
	"log"
	"time"
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
