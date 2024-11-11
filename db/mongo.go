package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

var Client *mongo.Client
var EscrowCollection *mongo.Collection

func Init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client 
	EscrowCollection = Client.Database("your_database_name").Collection("escrows") // Replace with your database name
}

// GetCollection provides access to a specific collection
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("your_database_name").Collection(collectionName)
}
