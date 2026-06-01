package inboundPort

import "userAuth/internal/core/models"

type AuthService interface {
	Login(email, password string) (*models.AuthToken, error)
}
