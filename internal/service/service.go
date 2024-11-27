package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
	"fmt"
)

type Material interface {
	CreateMaterial(material models.Material) (int, error)
	GetMaterialByID(id int) (models.Material, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.Material, error)
	LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error
}

type Competency interface {
	CreateCompetency(comp models.Competency) error
	GetAllCompetencies() ([]models.Competency, error)
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
}

type Service struct {
	Material
	Competency
}

func (s *Service) LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error {
	err := s.Material.LinkMaterialWithCompetencies(materialID, competencyIDs)
	if err != nil {
		return fmt.Errorf("error linking material with competencies: %w", err)
	}
	return nil
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Material:   NewMaterialService(repo),
		Competency: NewCompetencyService(repo),
	}
}
