package models

import "time"

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type SignInOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenOutput struct {
	AccessToken string `json:"access_token"`
}

type SignUpInput struct {
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=8,max=32"`
	Name           *string `json:"name"`
	Role           int     `json:"-"` // тоже можно отключить от JSON
	OrganizationID int     `json:"-"`
	DepartmentID   *int
	Code           string `json:"code"`
}

type Organization struct {
	OrganizationID int    `db:"organization_id" json:"organization_id"`
	Name           string `db:"name" json:"name"`
}

type RegistrationCode struct {
	Code             string     `db:"code" json:"code"`
	Prefix           string     `db:"prefix" json:"prefix"`
	OrganizationID   int        `db:"organization_id" json:"organization_id"`
	DepartmentID     *int       `db:"department_id" json:"department_id,omitempty"`
	MakeAdmin        bool       `db:"make_admin" json:"make_admin"`
	IsUsed           bool       `db:"used" json:"is_used"`
	ExpiresAt        *time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	Role             int        `db:"role" json:"role"`
	OrganizationName string     `db:"organization_name"`
	UsedAt           *time.Time `db:"used_at" json:"used_at"`
	DepartmentName   *string    `db:"department_name" json:"department_name,omitempty"`
}

type SignUpWithCodeInput struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8,max=32"`
	Name     *string `json:"name"`
	Code     string  `json:"code" validate:"required"` // Пример: MTI-WP93F6
}

type SignUpOutput struct {
	Message string `json:"message"`
}

type CreateRegistrationCodeInput struct {
	OrganizationID int  `json:"organization_id"`
	IsAdmin        bool `json:"is_admin"`
	DepartmentID   *int `json:"department_id,omitempty"`
}

type RegistrationCodeOutput struct {
	Code string `json:"code"`
}

type UserResponse struct {
	Email        string  `db:"email" json:"email"`
	Name         *string `db:"name" json:"name"`
	DepartmentID *int    `db:"department_id" json:"department_id"`
	Role         int     `db:"role" json:"role"`
}
