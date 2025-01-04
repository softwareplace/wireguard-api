package security

import "github.com/softwareplace/wireguard-api/pkg/utils/sec"

func (a *apiSecurityServiceImpl) Encrypt(value string) (string, error) {
	secret := a.Secret()
	return sec.Encrypt(value, secret)
}

func (a *apiSecurityServiceImpl) Decrypt(encrypted string) (string, error) {
	secret := a.Secret()
	return sec.Decrypt(encrypted, secret)
}
