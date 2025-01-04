package db

import (
	"context"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
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
		log.Panicf("Database connecting to %s failed: %v", uri, err)
	}
	return client
}

func connectionChecker() {
	log.Printf("Checking %s database connection...", dbEnv.Uri)
	err := GetDBClient().Ping(context.Background(), nil)
	if err != nil {
		log.Panicf("MongoDB %s connection check failed: %v", dbEnv.Uri, err)
	}

	log.Printf("Database %s connected successfully\n", dbEnv.Uri)
}
