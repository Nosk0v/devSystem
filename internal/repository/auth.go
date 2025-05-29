package repository

import (
	"devSystem/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strings"
	"time"
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

func (r *AuthPostgres) MarkRegistrationCodeAsUsed(code string) error {
	query := `UPDATE "InviteCode" SET used = TRUE WHERE code = $1`
	_, err := r.db.Exec(query, code)
	return err
}

func (r *AuthPostgres) GetOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	query := `SELECT organization_id, name FROM "Organization"`
	err := r.db.Select(&orgs, query)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func encodeWithMaskedPrefix(prefix string, n int) string {
	const letters = "ABCDEFGHJKMNPQRSTUVWXYZ23456789"
	code := make([]rune, n)
	maskPositions := []int{2, 5, 9}

	// заполняем случайным
	for i := 0; i < n; i++ {
		code[i] = rune(letters[rand.Intn(len(letters))])
	}

	// вставляем префикс по позициям
	for i, pos := range maskPositions {
		if i < len(prefix) && pos < n {
			code[pos] = rune(prefix[i])
		}
	}

	// форматируем с дефисами по 5
	return formatCode(string(code))
}

func formatCode(code string) string {
	var result []string
	for i := 0; i < len(code); i += 5 {
		end := i + 5
		if end > len(code) {
			end = len(code)
		}
		result = append(result, code[i:end])
	}
	return strings.Join(result, "-")
}

func (r *AuthPostgres) CreateRegistrationCode(prefix string, isAdmin bool) (string, error) {
	code := encodeWithMaskedPrefix(prefix, 30)
	role := 1
	if isAdmin {
		role = 0
	}
	query := `
	INSERT INTO "InviteCode" (code, prefix, role, used, expires_at, created_at)
	VALUES ($1, $2, $3, FALSE, NULL, $4)
	`
	_, err := r.db.Exec(query, code, prefix, role, time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to create registration code: %w", err)
	}
	return code, nil
}

func (r *AuthPostgres) GetRegistrationCodeInfo(code string) (*models.RegistrationCode, error) {
	var info models.RegistrationCode
	query := `
		SELECT rc.code, rc.prefix, rc.role, rc.used, rc.expires_at, rc.created_at,
		       o.organization_id, o.name AS organization_name
		FROM "InviteCode" rc
		JOIN "RegistrationPrefix" rp ON rc.prefix = rp.prefix
		JOIN "Organization" o ON rp.organization_id = o.organization_id
		WHERE rc.code = $1
	`
	err := r.db.Get(&info, query, code)
	if err != nil {
		return nil, fmt.Errorf("registration code not found: %w", err)
	}
	return &info, nil
}

func (r *AuthPostgres) DeleteRegistrationCode(code string) error {
	query := `DELETE FROM "InviteCode" WHERE code = $1`
	_, err := r.db.Exec(query, code)
	if err != nil {
		return fmt.Errorf("failed to delete registration code: %w", err)
	}
	return nil
}

func (r *AuthPostgres) GetPrefixByOrgID(orgID int) (string, error) {
	var prefix string
	err := r.db.Get(&prefix, `SELECT prefix FROM "RegistrationPrefix" WHERE organization_id = $1 LIMIT 1`, orgID)
	if err != nil {
		return "", fmt.Errorf("prefix not found for org: %w", err)
	}
	return prefix, nil
}
