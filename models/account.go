package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
	Role     int    `db:"role" json:"role"`
}

type JWTClaims struct {
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}
