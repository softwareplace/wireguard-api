package user

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"log"
	"net/http"
)

func (h *handlerImpl) CreateUser(ctx *api_context.ApiRequestContext[*request.ApiContext]) {
	server.GetRequestBody(ctx, models.User{}, h.validateUserFields, server.FailedToLoadBody)
}

func (h *handlerImpl) validateUserFields(ctx *api_context.ApiRequestContext[*request.ApiContext], user models.User) {
	if err := validator.ValidateUserFields(user); err != nil {
		log.Printf("[%s]:: validation failed with error: %v", ctx.GetSessionId(), err)
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.UsersRepository().FindUserByUsernameOrEmail(user.Username, user.Email)

	if err == nil {
		log.Printf("[%s]:: username or email already exists: %v", ctx.GetSessionId(), err)
		ctx.Error("Username or email already exists", http.StatusConflict)
		return
	}

	hashedPassword, salt, err := sec.HashPassword(user.Password)

	if err != nil {
		log.Printf("[%s]:: error encrypting password: %v", ctx.GetSessionId(), err)
		ctx.Error("Error encrypting password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	user.Salt = salt
	user.Status = "ACTIVE"

	if err := h.UsersRepository().Save(user); err != nil {
		log.Printf("[%s]:: error saving user_service to the database: %v", ctx.GetSessionId(), err)
		ctx.Error("Error saving user_service to the database", http.StatusInternalServerError)
		return
	}

	log.Printf("[%s]:: user_service created successfully", ctx.GetSessionId())
	ctx.Created(map[string]interface{}{"message": "User created successfully"})
}
