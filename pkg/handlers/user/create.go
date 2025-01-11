package user

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"net/http"
)

func (h *handlerImpl) CreateUser(ctx *request.ApiRequestContext) {
	request.GetRequestBody(ctx, models.User{}, h.validateUserFields, request.FailedToLoadBody)
}

func (h *handlerImpl) validateUserFields(ctx *request.ApiRequestContext, user models.User) {
	if err := validator.ValidateUserFields(user); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.UsersRepository().FindUserByUsernameOrEmail(user.Username, user.Email)

	if err == nil {
		ctx.Error("Username or email already exists", http.StatusConflict)
		return
	}

	hashedPassword, salt, err := sec.HashPassword(user.Password)

	if err != nil {
		ctx.Error("Error encrypting password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	user.Salt = salt
	user.Status = "ACTIVE"

	if err := h.UsersRepository().Save(user); err != nil {
		ctx.Error("Error saving user to the database", http.StatusInternalServerError)
		return
	}

	ctx.Created(map[string]interface{}{"message": "User created successfully"})
}
