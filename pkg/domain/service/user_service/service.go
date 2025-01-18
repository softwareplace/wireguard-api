package user_service

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/security"
	repo "github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
	"github.com/softwareplace/wireguard-api/pkg/utils/file"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"log"
	"net/http"
)

type Service interface {
	Init()
	LoadUserRoles(ctx api_context.ApiRequestContext[*request.ApiContext]) []string
}

type serviceImpl struct {
	appEnv          env.ApplicationEnv
	repository      repo.UsersRepository
	securityService security.ApiSecurityService[*request.ApiContext]
}

func (s *serviceImpl) LoadUserRoles(ctx api_context.ApiRequestContext[*request.ApiContext]) []string {
	user, err := s.repository.FindUserBySalt(ctx.RequestData.Salt())
	if err != nil {
		log.Printf("[%s]:: error finding user: %v", ctx.GetSessionId(), err)
		ctx.Error("Error finding user in the database", http.StatusInternalServerError)
		return nil
	}
	ctx.RequestData.User = user
	return user.Roles
}

func GetService() Service {
	return &serviceImpl{
		appEnv:          env.AppEnv(),
		repository:      repo.Repository(),
		securityService: security.GetApiSecurityService[*request.ApiContext](env.AppEnv().ApiSecretAuthorization),
	}
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
