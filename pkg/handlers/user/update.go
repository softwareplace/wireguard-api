package user

import (
	"encoding/json"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"net/http"
)

func (h *handlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.UserUpdate
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		shared.MakeErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	build := request.Build(w, r)
	accessContext := build.GetAccessContext()

	currentUser, err := h.UsersRepository().FindUserBySalt(accessContext.AccessId)

	if err != nil {
		shared.MakeErrorResponse(w, "Error finding user in the database", http.StatusInternalServerError)
		return
	}

	// If a new password is provided, hash it and generate a new salt
	if updatedUser.Password != "" {
		pass, salt, err := sec.HashPassword(updatedUser.Password)
		currentUser.Password = pass
		currentUser.Salt = salt
		if err != nil {
			shared.MakeErrorResponse(w, "Error encrypting password", http.StatusInternalServerError)
			return
		}
	}

	// Save updated user to database
	err = h.UsersRepository().Update(*currentUser)

	if err != nil {
		shared.MakeErrorResponse(w, "Error updating user in the database", http.StatusInternalServerError)
		return
	}
	shared.MakeResponse(w, map[string]interface{}{"message": "User updated successfully"}, http.StatusOK)
}
