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
	"time"
)

func (r *repositoryImpl) SaveAll(peers []models.Peer) error {
	database := r.collection()

	for _, peer := range peers {
		filter := bson.M{"publicKey": peer.PublicKey}
		nowToString := date.NowToString()
		peer.UpdatedAt = nowToString
		peer.Status = "AVAILABLE"

		if peer.LastHandshake != "" && peer.LastHandshake != "0" {

			lastHandshakeTime, err := time.Parse(time.RFC3339, peer.LastHandshake)
			if err != nil {
				log.Println("Error parsing last handshake time:", err)
			} else {
				// Check if the last handshake was less than five minutes ago
				if time.Since(lastHandshakeTime) < 5*time.Minute {
					peer.Status = "UNAVAILABLE"
				}
			}
		}

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

func (r *repositoryImpl) Update(peer models.Peer) error {
	database := r.collection()
	nowToString := date.NowToString()
	peer.UpdatedAt = nowToString

	filter := bson.M{"_id": peer.Id} // Identify the peer by its ID
	update := bson.M{"$set": peer}   // Update all fields in the peer object

	_, err := database.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error updating peer by ID in the database", err)
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

func (r *repositoryImpl) FindByPublicKey(pubKey string) (*models.Peer, error) {
	database := r.collection()
	var peer models.Peer

	err := database.FindOne(context.Background(), bson.M{"publicKey": pubKey}).Decode(&peer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No peer found with the given public key:", pubKey)
			return nil, nil
		}
		log.Println("Error retrieving the peer by public key from the database", err)
		return nil, err
	}

	return &peer, nil
}
