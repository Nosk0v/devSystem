package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
)

type MaterialService struct {
	repo *repository.Repository
}

func NewMaterialService(repo *repository.Repository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (m *MaterialService) CreateMaterial(material models.Material) error {
	return m.repo.CreateMaterial(material)
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
