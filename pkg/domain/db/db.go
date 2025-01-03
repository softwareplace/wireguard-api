package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func InitMongoDB() {
	connectionChecker()
	GetDB()
}

func GetDB() *mongo.Database {
	return GetDBClient().Database(os.Getenv("MONGO_DATABASE"))
}

func GetDBClient() *mongo.Client {

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	uri := os.Getenv("MONGO_URI")
	if username == "" || password == "" || uri == "" {
		log.Panicf("MongoDB environment variables are missing or incomplete")
	}
	credentials := options.Credential{
		Username: username,
		Password: password,
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %v", err)
	}
	return client
}

func connectionChecker() {
	log.Println("Checking database connection...")
	err := GetDBClient().Ping(context.Background(), nil)
	if err != nil {
		log.Panicf("MongoDB connection check failed: %v", err)
	}

	log.Println("Database connected successfully")
}
