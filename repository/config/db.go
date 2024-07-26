package repository

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func InitMongoDB() *mongo.Client {

	err := godotenv.Load(".env") // inget lokasi file .env nya	// Kesalahan berfikir
	if err != nil{
		log.Fatalf("Error loading .env file: %v", err)
	}


	// retrieve MongoDB URI from environment variables (.env file)
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == ""{
		log.Fatal("MONGODB_URI is not set in .env file")
	}


	// URI untuk menghubungkan dengan database MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI) 
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}