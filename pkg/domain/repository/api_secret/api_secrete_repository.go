package api_secret

import (
	"context"
	"errors"
	"github.com/eliasmeireles/wireguard-api/pkg/domain/db"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
	"github.com/eliasmeireles/wireguard-api/pkg/utils/date"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (r *repositoryImpl) Save(apiSecret models.ApiSecret) (*string, error) {
	database := db.GetDB().Collection("api_secret")

	apiSecret.CreatedAt = date.NowToString()
	apiSecret.UpdatedAt = date.NowToString()

	result, err := database.InsertOne(context.Background(), apiSecret)
	if err != nil {
		log.Println("Error saving peer to the database", err)
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

func (r *repositoryImpl) GetById(id string) (models.ApiSecret, error) {
	// Assuming 'db.GetDB()' returns a *mongo.Database
	database := db.GetDB().Collection("api_secret")

	// Convert string id to ObjectID (if needed)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.ApiSecret{}, errors.New("invalid ID format")
	}

	// Define a filter to match the document with the specified ID
	filter := bson.M{"_id": objectID}

	// Define a variable to hold the result
	var apiSecret models.ApiSecret

	// Find the document
	err = database.FindOne(context.Background(), filter).Decode(&apiSecret)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.ApiSecret{}, errors.New("document not found")
		}
		log.Printf("Error finding document: %v", err)
		return models.ApiSecret{}, err
	}

	// Return the result
	return apiSecret, nil
}
