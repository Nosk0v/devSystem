package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	Email          string `db:"email" json:"email"`
	Password       string `db:"password" json:"password"`
	Name           string `db:"name" json:"name"`
	Role           int    `db:"role" json:"role"`
	OrganizationID *int   `db:"organization_id" json:"organization_id"`
}

type JWTClaims struct {
	Email          string `json:"email"`
	TokenType      string `json:"token_type"`
	Role           int    `json:"role"`
	OrganizationID *int   `json:"organization_id,omitempty"`
	jwt.RegisteredClaims
}
