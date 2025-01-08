package user

import (
	"context"
	"fmt"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/date"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (r *usersRepositoryImpl) Save(user models.User) error {
	collection := r.collection()

	nowToString := date.NowToString()
	user.CreatedAt = nowToString
	user.UpdatedAt = nowToString

	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (r *usersRepositoryImpl) Update(user models.User) error {
	collection := r.collection()
	user.UpdatedAt = date.NowToString()

	filter := bson.M{"_id": user.Id}
	_, err := collection.ReplaceOne(context.Background(), filter, user)

	return err
}

func (r *usersRepositoryImpl) FindUserBySalt(salt string) (*models.User, error) {
	collection := r.collection()
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

func (r *usersRepositoryImpl) FindUserByEmail(email string) (*models.User, error) {
	collection := r.collection()
	var user models.User
	err := collection.FindOne(context.Background(), map[string]interface{}{
		"email": email,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	collection := r.collection()
	var user models.User
	err := collection.FindOne(context.Background(), map[string]interface{}{
		"username": username,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepositoryImpl) FindUserByUsernameOrEmail(username string, email string) (*models.User, error) {
	if username != "" {
		return r.FindUserByUsername(username)

	}

	if email != "" {
		return r.FindUserByEmail(email)
	}

	return nil, fmt.Errorf("username or email must be provided")
}
