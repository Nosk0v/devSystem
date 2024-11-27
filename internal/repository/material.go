package repository

import (
	"devSystem/models"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
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

func (r *MaterialRepository) GetAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	query := `
		SELECT 
			m.material_id,
			m.title,
			m.description,
			m.type,
			m.content,
			m.create_date,
			COALESCE(array_agg(mc.competency_id)::TEXT, '{}') AS competency_ids
		FROM material m
		LEFT JOIN MaterialCompetency mc ON m.material_id = mc.material_id
		GROUP BY m.material_id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var material models.Material
		var competencyIDs string // Временное хранение массива как строки

		err := rows.Scan(
			&material.MaterialID,
			&material.Title,
			&material.Description,
			&material.Type,
			&material.Content,
			&material.CreateDate,
			&competencyIDs,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Преобразуем строку компетенций в массив []int
		if err := parseStringToIntSlice(competencyIDs, &material.CompetencyIDs); err != nil {
			return nil, fmt.Errorf("error parsing competency IDs: %w", err)
		}

		materials = append(materials, material)
	}

	return materials, nil
}

func (r *MaterialRepository) GetMaterialByID(id int) (models.Material, error) {
	var material models.Material
	var competencyIDs string // Временное хранение массива как строки

	query := `
		SELECT 
			m.material_id,
			m.title,
			m.description,
			m.type,
			m.content,
			m.create_date,
			COALESCE(array_agg(mc.competency_id)::TEXT, '{}') AS competency_ids
		FROM material m
		LEFT JOIN MaterialCompetency mc ON m.material_id = mc.material_id
		WHERE m.material_id = $1
		GROUP BY m.material_id
	`

	err := r.db.QueryRow(query, id).Scan(
		&material.MaterialID,
		&material.Title,
		&material.Description,
		&material.Type,
		&material.Content,
		&material.CreateDate,
		&competencyIDs, // Получаем компетенции в виде строки
	)
	if err != nil {
		return models.Material{}, fmt.Errorf("error fetching material by ID: %w", err)
	}

	// Преобразуем строку компетенций в массив []int
	if err := parseStringToIntSlice(competencyIDs, &material.CompetencyIDs); err != nil {
		return models.Material{}, fmt.Errorf("error parsing competency IDs: %w", err)
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

func parseStringToIntSlice(input string, output *[]int) error {
	input = strings.ReplaceAll(input, "{", "[")
	input = strings.ReplaceAll(input, "}", "]")

	err := json.Unmarshal([]byte(input), output)
	if err != nil {
		return fmt.Errorf("error unmarshaling array: %w", err)
	}

	return nil
}
