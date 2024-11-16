package repository

import (
	"devSystem/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type MaterialRepository struct {
	db *sqlx.DB
}

func NewMaterialRepository(db *sqlx.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) CreateMaterial(material models.Material) error {
	query := `INSERT INTO material (title, description, type, content, create_date) 
              VALUES ($1, $2, $3, $4, $5)`

	log.Printf("Executing query: %s with params: %+v\n", query, material)

	_, err := r.db.Exec(query, material.Title, material.Description, material.Type, material.Content, material.CreateDate)
	if err != nil {
		return fmt.Errorf("error creating material: %w", err)
	}
	return nil
}

func (r *MaterialRepository) GetMaterialByID(id int) (models.Material, error) {
	var material models.Material
	query := `SELECT * FROM material WHERE material_id = $1`
	err := r.db.Get(&material, query, id)
	if err != nil {
		return models.Material{}, fmt.Errorf("error fetching material by ID: %w", err)
	}
	return material, nil
}

func (r *MaterialRepository) UpdateMaterial(material models.Material) error {
	query := `UPDATE material SET title = $1, description = $2, type = $3, content = $4 WHERE material_id = $5`
	_, err := r.db.Exec(query, material.Title, material.Description, material.Type, material.Content, material.MaterialID)
	if err != nil {
		return fmt.Errorf("error updating material: %w", err)
	}
	return nil
}

func (r *MaterialRepository) DeleteMaterial(id int) error {
	query := `DELETE FROM material WHERE material_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting material: %w", err)
	}
	return nil
}

func (r *MaterialRepository) GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	query := `SELECT * FROM material`
	err := r.db.Select(&materials, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}
	return materials, nil
}
