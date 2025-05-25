package repository

import (
	"devSystem/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) AccountExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM "Account" WHERE email = $1);`
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *AuthPostgres) CreateAccount(input *models.SignUpInput) error {

	query := `
	INSERT INTO "Account" ("email", "password", "name", "role", "organization_id") 
	VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, input.Email, input.Password, input.Name, input.Role, input.OrganizationID)
	return err
}
