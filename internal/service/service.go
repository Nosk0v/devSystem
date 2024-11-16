package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
)

type Material interface {
	CreateMaterial(material models.Material) error
	GetMaterialByID(id int) (models.Material, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.Material, error)
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

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Material:   NewMaterialService(repo),
		Competency: NewCompetencyService(repo),
	}
}
