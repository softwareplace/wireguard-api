package auth

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/security"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
	"sync"
)

var (
	openPath           []string
	openPathLock       sync.RWMutex
	apiSecurityService = security.GetApiSecurityService()
	usersRepo          = user.Repository()
)

func AddOpenPath(path string) {
	openPathLock.Lock()
	defer openPathLock.Unlock()
	openPath = append(openPath, path)
}

func AccessValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		openPathLock.RLock()
		defer openPathLock.RUnlock()
		matchFound := false
		for _, path := range openPath {
			if path == r.Method+"::"+r.URL.Path {
				matchFound = true
				break
			}
		}

		ctx := request.Build(w, r)

		if !matchFound {
			_, success := apiSecurityService.Validation(&ctx, _nextValidation)
			if !success {
				return
			}
		}

		next.ServeHTTP(w, ctx.Request)
	})
}

func _nextValidation(ctx *request.ApiRequestContext) (*models.User, bool) {
	accessContext := ctx.GetAccessContext()
	userData, err := usersRepo.FindUserBySalt(accessContext.AccessId)
	if err != nil {
		log.Printf("Failed to valiaded user %v access: %v", ctx, err)
		return nil, false
	}
	ctx.SetUser(userData)
	return userData, true
}
