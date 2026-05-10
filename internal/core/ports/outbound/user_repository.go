package outbound

import "userAuth/internal/core/domain/user"

type UserRepository interface {
	Create (user *user.User) error
	Update (user *user.User) error

	GetByID (id string) (*user.User, error)
	GetByEmail (email string) (*user.User, error)
	GetByUsername (username string) (*user.User, error)

	ExistsByEmail (email string) (bool, error)
	ExistsByUsername (username string) (bool, error)
}