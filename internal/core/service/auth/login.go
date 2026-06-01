package authService

import (
	"errors"
	"userAuth/internal/core/models"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *AuthService) Login(email, password string) (*models.AuthToken, error) {
	//TODO: get data from db verify and JWT
	u, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := verifyPassword(u.PasswordHash, password); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := generateToken(u.ID, u.Email)
	if err != nil {
		return nil, err
	}
	return token, nil
}
