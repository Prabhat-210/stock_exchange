package authService

import (
	"userAuth/internal/core/ports/outbound"

	"github.com/rs/zerolog"
)

type AuthService struct {
	log zerolog.Logger
	userRepo outbound.UserRepository
}

func NewAuthService(logger zerolog.Logger, userRepo outbound.UserRepository) *AuthService {
	return &AuthService{
		log: logger,
		userRepo: userRepo,
	}
}
