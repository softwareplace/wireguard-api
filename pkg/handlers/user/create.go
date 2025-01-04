package user

import (
	"encoding/json"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"net/http"
)

func (h *handlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		shared.MakeErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validateUserFields(user); err != nil {
		shared.MakeErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.UsersRepository().FindUserByUsernameOrEmail(user.Username, user.Email)

	if err == nil {
		shared.MakeErrorResponse(w, "Username or email already exists", http.StatusConflict)
		return
	}

	hashedPassword, salt, err := hashPassword(user.Password)

	if err != nil {
		shared.MakeErrorResponse(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword
	user.Salt = salt
	user.Status = "ACTIVE"

	if err := h.UsersRepository().Save(user); err != nil {
		shared.MakeErrorResponse(w, "Error saving user to the database", http.StatusInternalServerError)
		return
	}

	shared.MakeResponse(w, map[string]interface{}{"message": "User created successfully"}, http.StatusCreated)
}
