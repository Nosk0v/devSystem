package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"fmt"
)

type MaterialService struct {
	repo *repository.Repository
}

func NewMaterialService(repo *repository.Repository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (m *MaterialService) CreateMaterial(material models.Material) (int, error) {

	materialID, err := m.repo.CreateMaterial(material)
	if err != nil {
		return 0, fmt.Errorf("error creating material: %w", err)
	}
	if len(material.CompetencyIDs) > 0 {
		err = m.repo.LinkMaterialWithCompetencies(materialID, material.CompetencyIDs)
		if err != nil {
			return 0, fmt.Errorf("error linking material with competencies: %w", err)
		}
	}
	return materialID, nil
}

func (m *MaterialService) GetMaterialByID(id int) (models.Material, error) {
	return m.repo.GetMaterialByID(id)
}

func (m *MaterialService) UpdateMaterial(material models.Material) error {
	return m.repo.UpdateMaterial(material)
}

func (m *MaterialService) DeleteMaterial(id int) error {
	return m.repo.DeleteMaterial(id)
}

func (m *MaterialService) GetAllMaterials() ([]models.Material, error) {
	return m.repo.GetAllMaterials()
}

func (m *MaterialService) LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error {
	// Вызываем репозиторий для связывания материала с компетенциями
	err := m.repo.LinkMaterialWithCompetencies(materialID, competencyIDs)
	if err != nil {
		return fmt.Errorf("error linking material with competencies: %w", err)
	}
	return nil
}
