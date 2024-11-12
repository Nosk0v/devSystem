package repository

import (
	"devSystem/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateMaterial(material models.Material) error {
	query := `INSERT INTO material (title, description, type, content, create_date) 
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, material.Title, material.Description, material.Type, material.Content, material.CreateDate)
	if err != nil {
		return fmt.Errorf("error creating material: %w", err)
	}
	return nil
}

func (r *Repository) GetMaterialByID(id int) (models.Material, error) {
	var material models.Material
	query := `SELECT * FROM material WHERE material_id = $1`
	err := r.db.Get(&material, query, id)
	if err != nil {
		return models.Material{}, fmt.Errorf("error fetching material by ID: %w", err)
	}
	return material, nil
}

func (r *Repository) UpdateMaterial(material models.Material) error {
	query := `UPDATE material SET title = $1, description = $2, type = $3, content = $4 WHERE material_id = $5`
	_, err := r.db.Exec(query, material.Title, material.Description, material.Type, material.Content, material.MaterialID)
	if err != nil {
		return fmt.Errorf("error updating material: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMaterial(id int) error {
	query := `DELETE FROM material WHERE material_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting material: %w", err)
	}
	return nil
}

func (r *Repository) CreateCompetency(comp models.Competency) error {
	query := `INSERT INTO competency (name, description, parent_id, create_date) 
              VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, comp.Name, comp.Description, comp.ParentID, comp.CreateDate)
	if err != nil {
		return fmt.Errorf("error creating competency: %w", err)
	}
	return nil
}

func (r *Repository) UpdateCompetency(comp models.Competency) error {
	query := `UPDATE competency SET name = $1, description = $2, parent_id = $3 WHERE competency_id = $4`
	_, err := r.db.Exec(query, comp.Name, comp.Description, comp.ParentID, comp.CompetencyID)
	if err != nil {
		return fmt.Errorf("error updating competency: %w", err)
	}
	return nil
}

func (r *Repository) DeleteCompetency(id int) error {
	query := `DELETE FROM competency WHERE competency_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting competency: %w", err)
	}
	return nil
}

func (r *Repository) GetAllMaterials() ([]models.Material, error) {
	var material []models.Material
	query := `SELECT * FROM material`
	err := r.db.Select(&material, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}
	return material, nil
}

func (r *Repository) GetAllCompetencies() ([]models.Competency, error) {
	var competencies []models.Competency
	query := `SELECT * FROM competency`
	err := r.db.Select(&competencies, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all competencies: %w", err)
	}
	return competencies, nil
}
