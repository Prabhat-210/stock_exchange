package models

type AuthToken struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
}