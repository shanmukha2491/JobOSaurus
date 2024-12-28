package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase *mongo.Database

func ConnectMongo() {

	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		log.Fatalf("MONGO_URL environment variable not set")
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increased timeout
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Select the database
	database := client.Database("JobOSaurus")
	MongoDatabase = database

	log.Println("Connected to MongoDB successfully")
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := MongoDatabase.Collection(collectionName)
	return collection
}
