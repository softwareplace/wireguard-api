package user

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type UsersRepository interface {
	Save(user models.User) error
	Update(user models.User) error
	FindUserBySalt(salt string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserByUsernameOrEmail(username string, email string) (*models.User, error)
}

type usersRepositoryImpl struct {
	database *mongo.Database
}

func (r *usersRepositoryImpl) collection() *mongo.Collection {
	return r.database.Collection("users")
}

var (
	repository UsersRepository
	once       sync.Once
)

func Repository() *UsersRepository {
	once.Do(func() {
		repository = &usersRepositoryImpl{
			database: db.GetDB(),
		}
	})
	return &repository
}
