package security

import (
	"time"

	"userAuth/internal/core/models"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key") // TODO: move to config/env

func GenerateToken(
	userID string,
	email string,
) (*models.AuthToken, error) {

	expiry := time.Now().Add(15 * time.Minute).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     expiry,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &models.AuthToken{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expiry,
	}, nil
}