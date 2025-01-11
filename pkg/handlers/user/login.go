package user

import (
	"errors"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (h *handlerImpl) Login(ctx *request.ApiRequestContext) {
	request.GetRequestBody(ctx, models.User{}, h.checkUserCredentials, request.FailedToLoadBody)
}

func (h *handlerImpl) checkUserCredentials(ctx *request.ApiRequestContext, userInput models.User) {
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
			ctx.Unauthorized()
			return
		}
		ctx.InternalServerError("Internal Server Error")

		return
	}

	if !validator.CheckPassword(userInput.Password, userResponse.Password, userResponse.Salt) {
		ctx.Unauthorized()
		return
	}

	// Generate JWT and respond
	tokenData, err := h.ApiSecurityService().GenerateJWT(*userResponse)
	if err != nil {
		ctx.InternalServerError("Error generating token")
		return
	}

	ctx.Ok(tokenData)
}
