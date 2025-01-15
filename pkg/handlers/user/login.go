package user

import (
	"errors"
	"github.com/softwareplace/http-utils/server"

	"github.com/softwareplace/wireguard-api/pkg/models"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"github.com/softwareplace/wireguard-api/pkg/utils/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (h *handlerImpl) Login(ctx *server.ApiRequestContext) {
	server.GetRequestBody(ctx, models.User{}, h.checkUserCredentials, server.FailedToLoadBody)
}

func (h *handlerImpl) checkUserCredentials(ctx *server.ApiRequestContext, userInput models.User) {
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
		log.Printf("[%s]:: find user by username or email failed: %v", ctx.GetSessionId(), err)

		return
	}

	if !validator.CheckPassword(userInput.Password, userResponse.Password, userResponse.Salt) {
		ctx.Unauthorized()
		return
	}

	// Generate JWT and respond
	tokenData, err := h.ApiSecurityService().GenerateJWT(*userResponse)
	if err != nil {
		log.Printf("[%s]:: generating new jwt failed: %v", ctx.GetSessionId(), err)
		ctx.InternalServerError("Error generating token")
		return
	}

	ctx.Ok(tokenData)
}
