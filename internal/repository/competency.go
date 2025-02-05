package repository

import (
	"devSystem/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CompetencyRepository struct {
	db *sqlx.DB
}

func NewCompetencyRepository(db *sqlx.DB) *CompetencyRepository {
	return &CompetencyRepository{db: db}
}

func (r *CompetencyRepository) CreateCompetency(comp models.Competency) error {
	query := `INSERT INTO competency (name, description, parent_id, create_date) 
              VALUES ($1, $2, $3, $4)`

	log.Printf("Executing query: %s with params: Name=%s, Description=%s, ParentID=%v, CreateDate=%v\n",
		query, comp.Name, comp.Description, comp.ParentID, comp.CreateDate)

	_, err := r.db.Exec(query, comp.Name, comp.Description, comp.ParentID, comp.CreateDate)
	if err != nil {
		return fmt.Errorf("error creating competency: %w", err)
	}

	return nil
}

func (r *CompetencyRepository) UpdateCompetency(comp models.Competency) error {
	query := `UPDATE competency SET name = $1, description = $2, parent_id = $3 WHERE competency_id = $4`
	_, err := r.db.Exec(query, comp.Name, comp.Description, comp.ParentID, comp.CompetencyID)
	if err != nil {
		return fmt.Errorf("error updating competency: %w", err)
	}
	return nil
}

func (r *CompetencyRepository) DeleteCompetency(id int) error {
	query := `DELETE FROM competency WHERE competency_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting competency: %w", err)
	}
	return nil
}

func (r *CompetencyRepository) GetAllCompetencies() ([]models.CompetencyResponse, error) {
	var competencies []models.CompetencyResponse
	query := `
		SELECT c.competency_id, c.name, c.description, COALESCE(p.name, '') AS parent_name, c.create_date
		FROM Competency c
		LEFT JOIN Competency p ON c.parent_id = p.competency_id
		ORDER BY c.competency_id
`
	err := r.db.Select(&competencies, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all competencies: %w", err)
	}
	return competencies, nil
}

func (r CompetencyRepository) GetCompetencyByID(id int) (models.CompetencyResponse, error) {
	var competency models.CompetencyResponse
	query := `
		SELECT c.competency_id, c.name, c.description, COALESCE(p.name, '') AS parent_name, c.create_date
		FROM Competency c
		LEFT JOIN Competency p ON c.parent_id = p.competency_id
		WHERE c.competency_id = $1
`
	err := r.db.Get(&competency, query, id)
	if err != nil {
		return models.CompetencyResponse{}, fmt.Errorf("error fetching competency by ID: %w", err)
	}
	return competency, nil
}
