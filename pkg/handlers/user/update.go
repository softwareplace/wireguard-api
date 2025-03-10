package user

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"log"
	"net/http"
)

func (h *handlerImpl) UpdateUser(ctx *api_context.ApiRequestContext[*request.ApiContext]) {
	server.GetRequestBody(ctx, models.UserUpdate{}, h.useUpdateValidation, server.FailedToLoadBody)
}

func (h *handlerImpl) useUpdateValidation(ctx *api_context.ApiRequestContext[*request.ApiContext], updatedUser models.UserUpdate) {
	currentUser, err := h.UsersRepository().FindUserBySalt(ctx.AccessId)

	if err != nil {
		log.Printf("[%s]:: find user by salt failed: %v", ctx.GetSessionId(), err)

		ctx.Error("Error finding user in the database", http.StatusInternalServerError)
		return
	}

	// If a new password is provided, hash it and generate a new salt
	if updatedUser.Password != "" {
		pass, salt, err := sec.HashPassword(updatedUser.Password)
		currentUser.Password = pass
		currentUser.Salt = salt
		if err != nil {
			log.Printf("[%s]:: encrypting password failed: %v", ctx.GetSessionId(), err)
			ctx.Error("Error encrypting password", http.StatusInternalServerError)
			return
		}
	}

	// Save updated user to database
	err = h.UsersRepository().Update(*currentUser)

	if err != nil {
		log.Printf("[%s]:: updateing user failed: %v", ctx.GetSessionId(), err)
		ctx.Error("Error updating user in the database", http.StatusInternalServerError)
		return
	}
	log.Printf("[%s]:: user successfully updated", ctx.GetSessionId())
	ctx.Ok(map[string]interface{}{"message": "User updated successfully"})
}
