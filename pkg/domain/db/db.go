package db

import (
	"context"
	"github.com/eliasmeireles/wireguard-api/pkg/utils/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var dbEnv = env.AppEnv().DBEnv

func InitMongoDB() {
	connectionChecker()
	GetDB()
}

func GetDB() *mongo.Database {
	return GetDBClient().Database(dbEnv.DatabaseName)
}

func GetDBClient() *mongo.Client {

	username := dbEnv.Username
	password := dbEnv.Password
	uri := dbEnv.Uri
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
