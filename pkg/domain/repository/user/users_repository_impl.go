package user

import (
	"context"
	"fmt"
	"github.com/eliasmeireles/wireguard-api/pkg/domain/db"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
	"github.com/eliasmeireles/wireguard-api/pkg/utils/date"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (impl *usersRepositoryImpl) Save(user models.User) error {
	collection := db.GetDB().Collection("users")

	nowToString := date.NowToString()
	user.CreatedAt = nowToString
	user.UpdatedAt = nowToString

	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (impl *usersRepositoryImpl) Update(user models.User) error {
	collection := db.GetDB().Collection("users")
	user.UpdatedAt = date.NowToString()

	filter := bson.M{"_id": user.Id}
	_, err := collection.ReplaceOne(context.Background(), filter, user)

	return err
}

func (impl *usersRepositoryImpl) FindUserBySalt(salt string) (*models.User, error) {
	collection := db.GetDB().Collection("users")
	var currentUser models.User

	err := collection.FindOne(context.Background(), map[string]interface{}{
		"salt": salt,
	}).Decode(&currentUser)

	if err != nil {
		log.Printf("Error finding user by salt: %v", err)
		return nil, err
	}
	return &currentUser, nil
}

func (impl *usersRepositoryImpl) FindUserByEmail(email string) (*models.User, error) {
	collection := db.GetDB().Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), map[string]interface{}{
		"email": email,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (impl *usersRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	collection := db.GetDB().Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), map[string]interface{}{
		"username": username,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (impl *usersRepositoryImpl) FindUserByUsernameOrEmail(username string, email string) (*models.User, error) {
	if username != "" {
		return impl.FindUserByUsername(username)

	}

	if email != "" {
		return impl.FindUserByEmail(email)
	}

	return nil, fmt.Errorf("username or email must be provided")
}
