package user

import (
	"encoding/json"
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/request"
	"github.com/eliasmeireles/wireguard-api/pkg/handlers/shared"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
	"net/http"
)

func (h *handlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		shared.MakeErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	requestContext := request.Build(w, r)

	currentUser, success := h.ApiSecurityService().Validation(&requestContext, isCurrentUser(h))

	if !success {
		return
	}

	// Validate each field; if empty, use the current value from the database

	if updatedUser.Role == "" {
		updatedUser.Role = currentUser.Role
	}

	// If a new password is provided, hash it and generate a new salt
	if updatedUser.Password != "" {
		pass, salt, err := hashPassword(updatedUser.Password)
		currentUser.Password = pass
		currentUser.Salt = salt
		if err != nil {
			shared.MakeErrorResponse(w, "Error encrypting password", http.StatusInternalServerError)
			return
		}
	}

	// Update user fields with the new data
	currentUser.Role = updatedUser.Role

	// Save updated user to database
	err = h.UsersRepository().Update(*currentUser)

	if err != nil {
		shared.MakeErrorResponse(w, "Error updating user in the database", http.StatusInternalServerError)
		return
	}
	shared.MakeResponse(w, map[string]interface{}{"message": "User updated successfully"}, http.StatusOK)
}

func isCurrentUser(h *handlerImpl) func(ctx *request.ApiRequestContext) (*models.User, bool) {
	return func(ctx *request.ApiRequestContext) (*models.User, bool) {
		accessContext := ctx.GetAccessContext()
		userData, err := h.UsersRepository().FindUserBySalt(accessContext.AccessId)
		if err != nil {
			shared.MakeErrorResponse(ctx.Writer, "Authorization failed", http.StatusForbidden)
			return nil, false
		}
		ctx.SetUser(userData)
		return userData, true
	}
}
