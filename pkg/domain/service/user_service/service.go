package user_service

import (
	repo "github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
	"github.com/softwareplace/wireguard-api/pkg/utils/file"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"log"
	"sync"
)

type Service interface {
	Init()
}

type serviceImpl struct {
	appEnv     env.ApplicationEnv
	repository repo.UsersRepository
}

var (
	serviceOnce     sync.Once
	serviceInstance Service
)

func GetService() Service {
	serviceOnce.Do(func() {
		serviceInstance = &serviceImpl{
			appEnv:     env.AppEnv(),
			repository: repo.Repository(),
		}
	})
	return serviceInstance
}

type userInit struct {
	User models.User `yaml:"user"`
}

func (s *serviceImpl) Init() {
	if s.appEnv.InitFilePath != "" {
		init, err := file.FromYaml(s.appEnv.InitFilePath, userInit{})
		if err != nil {
			log.Fatalf("Failed to read init file: %v", err)
		}

		initUserData := init.User

		if err := validator.ValidateUserFields(initUserData); err != nil {
			log.Fatalf("Failed to validate init user data: %v", err)
		}

		_, err = s.repository.FindUserByUsernameOrEmail(initUserData.Username, initUserData.Email)

		if err == nil {
			return
		}

		hashedPassword, salt, err := sec.HashPassword(initUserData.Password)

		if err != nil {
			log.Fatalf("Failed to encrypt init user password: %v", err)
		}

		initUserData.Password = hashedPassword
		initUserData.Salt = salt
		initUserData.Status = "ACTIVE"

		err = s.repository.Save(initUserData)

		if err != nil {
			log.Fatalf("Failed to save init user data: %v", err)
		}
	}
}
