package auth

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/user"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/security"
	"net/http"
	"sync"
)

var (
	roles              = make(map[string][]string)
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

func AddRoles(path string, requiredRoles ...string) {
	if len(requiredRoles) > 0 {
		roles[path] = requiredRoles
	}
}

func GetRolesForPath(r *http.Request) ([]string, bool) {
	path := r.Method + "::" + r.URL.Path
	if roles[path] != nil {
		return roles[path], true
	}
	return nil, false
}
