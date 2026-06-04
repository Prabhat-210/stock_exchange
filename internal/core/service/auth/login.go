package authService

import (
	"errors"
	requestDTO "userAuth/internal/adapters/inbound/http/request_DTO"
	"userAuth/internal/core/domain/user"
	"userAuth/internal/core/models"

	"github.com/google/uuid"
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

func (s *AuthService) SignUp(req *requestDTO.SignUpRequest) (*models.AuthToken, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userId := uuid.New().String();
	user := &user.User{
		ID: 		  userId,
		Email:        req.Email,
		Username:     req.UserName,
		PasswordHash: hashedPassword,
	}
	err = s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}
	token, err := generateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return token, nil
}
