package user

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"net/http"
)

func (h *handlerImpl) UpdateUser(ctx *request.ApiRequestContext) {
	request.GetRequestBody(ctx, models.UserUpdate{}, h.useUpdateValidation, request.FailedToLoadBody)
}

func (h *handlerImpl) useUpdateValidation(ctx *request.ApiRequestContext, updatedUser models.UserUpdate) {
	currentUser, err := h.UsersRepository().FindUserBySalt(ctx.AccessContext.AccessId)

	if err != nil {
		ctx.Error("Error finding user in the database", http.StatusInternalServerError)
		return
	}

	// If a new password is provided, hash it and generate a new salt
	if updatedUser.Password != "" {
		pass, salt, err := sec.HashPassword(updatedUser.Password)
		currentUser.Password = pass
		currentUser.Salt = salt
		if err != nil {
			ctx.Error("Error encrypting password", http.StatusInternalServerError)
			return
		}
	}

	// Save updated user to database
	err = h.UsersRepository().Update(*currentUser)

	if err != nil {
		ctx.Error("Error updating user in the database", http.StatusInternalServerError)
		return
	}
	ctx.Ok(map[string]interface{}{"message": "User updated successfully"})
}
