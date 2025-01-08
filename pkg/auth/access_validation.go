package auth

import (
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

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

			if !hasResourceAccess(w, r, ctx) {
				return
			}
		}

		next.ServeHTTP(w, ctx.Request)
	})
}

func hasResourceAccess(w http.ResponseWriter, r *http.Request, ctx request.ApiRequestContext) bool {
	userRoles, err := ctx.GetRoles()

	if err != nil {
		shared.MakeErrorResponse(w, "You are not allowed to access this resource", http.StatusUnauthorized)
		return false
	}

	accessRoles, hasRoles := GetRolesForPath(r)

	if !hasRoles {
		shared.MakeErrorResponse(w, "You are not allowed to access this resource", http.StatusUnauthorized)
		return false
	}

	hasAccess := false

	for _, userRole := range userRoles {
		for _, accessRole := range accessRoles {
			if userRole == accessRole {
				hasAccess = true
				break
			}
		}
		if hasAccess {
			break
		}
	}

	if !hasAccess {
		shared.MakeErrorResponse(w, "You are not allowed to access this resource", http.StatusUnauthorized)
		return false
	}
	return true
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
