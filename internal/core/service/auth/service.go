package authService

import "github.com/rs/zerolog"

type AuthService struct {
	log zerolog.Logger
}

func NewAuthService(logger zerolog.Logger) *AuthService {
	return &AuthService{
		log: logger,
	}
}
