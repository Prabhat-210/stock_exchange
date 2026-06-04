package repository

import (
	"context"
	"userAuth/internal/core/domain/user"
	"userAuth/internal/core/ports/outbound"

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

func (r *userRepository) Save(req *user.User) error {
	query := `
		INSERT INTO auth.users (
    		id,
    		email,
   			user_name,
    		password_hash
		)
		VALUES (
    		$1,
   			$2,
   			$3,
    		$4
		)
		RETURNING id;
	`

	// var user user.User
	var userID string

	err := r.db.QueryRow(
		context.Background(),
		query,
		req.ID,
		req.Email,
		req.Username,
		req.PasswordHash,
	).Scan(&userID)

	if err != nil {
		return err
	}
	return nil
}
