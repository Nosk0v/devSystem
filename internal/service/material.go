package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"fmt"
)

type MaterialService struct {
	repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (m *MaterialService) CreateMaterial(material models.Material) (int, error) {
	materialID, err := m.repo.CreateMaterial(material)
	if err != nil {
		return 0, fmt.Errorf("error creating material: %w", err)
	}
	return materialID, nil
}

func (m *MaterialService) GetMaterialByID(id int) (models.MaterialResponse, error) {
	material, err := m.repo.GetMaterialByID(id)
	if err != nil {
		return models.MaterialResponse{}, fmt.Errorf("error fetching material by ID: %w", err)
	}
	return material, nil
}

func (m *MaterialService) GetAllMaterials() ([]models.MaterialResponse, error) {
	materials, err := m.repo.GetAllMaterials()
	if err != nil {
		return nil, fmt.Errorf("error fetching all materials: %w", err)
	}
	return materials, nil
}

func (s *MaterialService) UpdateMaterial(material models.Material) error {
	return s.repo.UpdateMaterial(material)
}

func (m *MaterialService) DeleteMaterial(id int) error {
	err := m.repo.DeleteMaterial(id)
	if err != nil {
		return fmt.Errorf("error deleting material: %w", err)
	}
	return nil
}

func (m *MaterialService) LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error {
	err := m.repo.LinkMaterialWithCompetencies(materialID, competencyIDs)
	if err != nil {
		return fmt.Errorf("error linking material with competencies: %w", err)
	}
	return nil
}
