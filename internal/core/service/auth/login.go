package authService

import (
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *AuthService) Login(email, password string) (string, error) {
	//TODO: get data from db verify and JWT
	return "login successfully", nil
}
