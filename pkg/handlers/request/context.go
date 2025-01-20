package request

type UserPrincipal struct {
	Username string
	Email    string
	Salt     string
	Roles    []string
	Status   string
}

type ApiContext struct {
	User                *UserPrincipal
	AccessId            string
	ApiKeyId            string
	AuthorizationClaims map[string]interface{}
	ApiKeyClaims        map[string]interface{}
}

func NewApiContext(user *UserPrincipal) *ApiContext {
	return &ApiContext{
		User: user,
	}
}

func (ctx *ApiContext) GetSalt() string {
	return ctx.User.Salt
}

func (ctx *ApiContext) GetRoles() []string {
	return ctx.User.Roles
}
