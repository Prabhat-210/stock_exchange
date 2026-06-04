package inboundPort

import (
	requestDTO "userAuth/internal/adapters/inbound/http/request_DTO"
	"userAuth/internal/core/models"
)

type AuthService interface {
	Login(email, password string) (*models.AuthToken, error)
	SignUp(req *requestDTO.SignUpRequest) (*models.AuthToken, error)
}
