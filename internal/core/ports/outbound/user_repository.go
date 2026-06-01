package outbound

import "userAuth/internal/core/domain/user"

type UserRepository interface {
	GetByEmail(email string) (*user.User, error)
}