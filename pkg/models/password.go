package models

import (
	"errors"
	"fmt"

	"github.com/trevino-676/passmanager/pkg/encrypt"
)

// Password struct
// This represent a registry in the password's document.
// Platform: Plataform or webpage to which the password corresponds.
// User: user or email for the password.
// Password: password of the platform.
type Password struct {
	Platform string
	User     string
	Password string
}

// NewPassword
// This function create a instance of password struct.
// Params:
// platform (string): Name of the platform
// user (string): user for the password.
// password (string): password.
// Returns:
// An instance of password struct
func NewPassword(platform, user, password string) *Password {
	return &Password{platform, user, password}
}

// Validate
// Validate that exists the information in the struct
// Returns
// A boolean. True if all the information is correct, false if not.
func (p *Password) Validate() bool {
	if p.Platform != "" && p.User != "" && p.Password != "" {
		return true
	}
	return false
}

// ToString
// Convert the struct into a string
// Returns
// string with the struct data.
func (p *Password) ToString() string {
	return fmt.Sprintf("%s|%s|%s", p.Platform, p.User, p.Password)
}

// Encrypt
// Encrypt the Password registry
// Params:
// secretKey (string): secretKey with 16 lenght characters
// Returns:
// string encrypted and error.
func (p *Password) Encrypt(secretKey string) (string, error) {
	if len(secretKey) != 16 {
		return "", errors.New("the secretKey doesn't have 16 characters")
	}
	return encrypt.EncryptRegistry(p.ToString(), secretKey)
}
