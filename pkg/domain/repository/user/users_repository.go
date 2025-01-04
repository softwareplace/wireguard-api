package user

import (
	"github.com/softwareplace/wireguard-api/pkg/models"
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
}

func Repository() UsersRepository {
	return &usersRepositoryImpl{}
}
