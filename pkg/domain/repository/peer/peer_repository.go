package peer

import (
	"context"
	"errors"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/date"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (r *repositoryImpl) SaveAll(peers []models.Peer) error {
	database := r.collection()

	for _, peer := range peers {
		filter := bson.M{"fileName": peer.FileName}
		nowToString := date.NowToString()
		peer.UpdatedAt = nowToString

		// Check if the peer already exists in the database
		var existingPeer models.Peer
		err := database.FindOne(context.Background(), filter).Decode(&existingPeer)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				peer.CreatedAt = nowToString // Set CreatedAt only if not already present
			} else {
				log.Println("Error checking if peer exists in the database", err)
				return err
			}
		} else {
			peer.Status = existingPeer.Status
			peer.CreatedAt = existingPeer.CreatedAt // Preserve original CreatedAt if peer exists
		}

		update := bson.M{"$set": peer}

		_, err = database.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
		if err != nil {
			log.Println("Error saving or updating peer to the database", err)
			return err
		}
	}

	return nil
}

func (r *repositoryImpl) Save(peer models.Peer) error {
	database := r.collection()
	nowToString := date.NowToString()
	peer.UpdatedAt = nowToString
	peer.CreatedAt = nowToString

	_, err := database.InsertOne(context.Background(), peer)
	if err != nil {
		log.Println("Error saving peer to the database", err)
		return err
	}

	return nil
}

func (r *repositoryImpl) GetAvailablePeer() (*models.Peer, error) {
	database := r.collection()

	var peer models.Peer
	err := database.FindOne(context.Background(), bson.M{}).Decode(&peer)
	if err != nil {
		log.Println("Error retrieving the first peer from the database", err)
		return nil, err
	}

	return &peer, nil
}
