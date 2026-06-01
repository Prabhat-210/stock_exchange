package repository

import (
	"context"
	"userAuth/internal/core/ports/outbound"
	"userAuth/internal/core/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) outbound.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByEmail(email string) (*user.User, error) {
	query := `
		SELECT
			id,
			email,
			user_name,
			password_hash,
			is_email_verified,
			is_active,
			last_login_at,
			created_at,
			updated_at
		FROM auth.users
		WHERE email = $1
	`

	var u user.User
	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.PasswordHash,
		&u.IsEmailVerified,
		&u.IsActive,
		&u.LastLoginAt,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
