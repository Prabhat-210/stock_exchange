package user

import "time"

type User struct {
	ID string
	Email string
	UserName string
	PasswordHash string
	IsEmailVerified bool
	IsActive bool
	LastLoginAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CanLogin() bool  {
	return u.IsActive && u.IsEmailVerified
}

func (u *User) MarkEmailIsVerified() {
	u.IsEmailVerified = true
}

func (u *User) Activate() {
	u.IsActive = true
}

func (u *User) Deactivate() {
	u.IsActive = false
}

func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
}