package interfaces

import "fmt"

// AuthProvider implements Authenticator
type AuthProvider struct {
	authorized map[string]string
}

func (ap *AuthProvider) Authenticate(username, password string) bool {
	return ap.authorized[username] == password
}

// GOOD
func NewAuthenticator(authorized map[string]string) *AuthProvider {
	return &AuthProvider{authorized: authorized}
}

// // BAD
// func NewAuthenticator(authorized map[string]string) Authenticator {
// 	return &AuthProvider{authorized: authorized}
// }

// somewhere in another package ...
type Authenticator interface {
	Authenticate(username, password string) bool
}

func DoSomeAuthentication(provider Authenticator) error {
	if !provider.Authenticate("felix", "top_secret_password") {
		return fmt.Errorf("no authentication possible")
	}

	return fmt.Errorf("auth successful")
}
