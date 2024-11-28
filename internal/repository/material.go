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

func (r *MaterialRepository) CreateMaterial(material models.Material) (int, error) {
	query := `INSERT INTO material (title, description, type, content, create_date) 
              VALUES ($1, $2, $3, $4, $5) RETURNING material_id`

	log.Printf("Executing query: %s with params: %+v\n", query, material)

	var materialID int
	err := r.db.QueryRow(query, material.Title, material.Description, material.Type, material.Content, material.CreateDate).Scan(&materialID)
	if err != nil {
		return 0, fmt.Errorf("error creating material: %w", err)
	}

	return materialID, nil
}

func (r *MaterialRepository) LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error {
	query := `INSERT INTO MaterialCompetency (material_id, competency_id) VALUES ($1, $2)`
	for _, competencyID := range competencyIDs {
		_, err := r.db.Exec(query, materialID, competencyID)
		if err != nil {
			return fmt.Errorf("error linking material with competency ID %d: %w", competencyID, err)
		}
	}
	return nil
}

func (r *MaterialRepository) GetMaterialByID(id int) (models.MaterialResponse, error) {
	var material models.MaterialResponse
	query := `
		SELECT m.material_id, m.title, m.description, mt.type AS type_name, m.content, m.create_date,
		       array_agg(c.name) AS competencies
		FROM Material m
		LEFT JOIN MaterialType mt ON m.type = mt.type_id
		LEFT JOIN MaterialCompetency mc ON m.material_id = mc.material_id
		LEFT JOIN Competency c ON mc.competency_id = c.competency_id
		WHERE m.material_id = $1
		GROUP BY m.material_id, mt.type
	`
	err := r.db.Get(&material, query, id)
	if err != nil {
		return models.MaterialResponse{}, fmt.Errorf("error fetching material by ID: %w", err)
	}
	return material, nil
}

func (r *MaterialRepository) UpdateMaterial(material models.Material) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := `UPDATE material SET title = $1, description = $2, type = $3, content = $4 WHERE material_id = $5`
	_, err = tx.Exec(query, material.Title, material.Description, material.Type, material.Content, material.MaterialID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update material: %w", err)
	}

	_, err = tx.Exec("DELETE FROM materialcompetency WHERE material_id = $1", material.MaterialID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting old competency links: %w", err)
	}

	for _, competencyID := range material.Competencies {
		_, err := tx.Exec("INSERT INTO materialcompetency (material_id, competency_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", material.MaterialID, competencyID)
		if err != nil {
			tx.Rollback()
			return nil
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *MaterialRepository) GetAllMaterials() ([]models.MaterialResponse, error) {
	var materials []models.MaterialResponse
	query := `
		SELECT m.material_id, m.title, m.description, mt.type AS type_name, m.content, m.create_date,
		       array_agg(c.name) AS competencies
		FROM Material m
		LEFT JOIN MaterialType mt ON m.type = mt.type_id
		LEFT JOIN MaterialCompetency mc ON m.material_id = mc.material_id
		LEFT JOIN Competency c ON mc.competency_id = c.competency_id
		GROUP BY m.material_id, mt.type
	`
	err := r.db.Select(&materials, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}
	return materials, nil
}

func (r *MaterialRepository) DeleteMaterial(id int) error {
	query := `DELETE FROM material WHERE material_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting material: %w", err)
	}
	return nil
}
