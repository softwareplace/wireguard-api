package user

import (
	"encoding/json"
	"errors"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func (h *handlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		shared.MakeErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}
	decrypt, err := sec.Decrypt(userInput.Password, []byte(sec.SampleEncryptKey))

	if err != nil {
		log.Printf("Failed to decrypt password: %v", err)
	} else {
		userInput.Password = decrypt
	}

	// Validate userResponse credentials
	userResponse, err := h.UsersRepository().FindUserByUsernameOrEmail(userInput.Username, userInput.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			shared.MakeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		shared.MakeErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !checkPassword(userInput.Password, userResponse.Password, userResponse.Salt) {
		shared.MakeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Generate JWT and respond
	tokenData, err := h.ApiSecurityService().GenerateJWT(*userResponse)
	if err != nil {
		shared.MakeErrorResponse(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tokenData)
	if err != nil {
		shared.MakeErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
