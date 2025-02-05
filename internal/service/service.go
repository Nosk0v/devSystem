package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
)

type Material interface {
	CreateMaterial(material models.Material) (int, error)
	GetMaterialByID(id int) (models.MaterialResponse, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.MaterialResponse, error)
	LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error
	GetMaterialTypeByID(id int) (models.MaterialType, error)
	GetAllMaterialTypes() ([]models.MaterialType, error)
	CreateMaterialType(materialType models.MaterialType) (int, error)
}

type Competency interface {
	CreateCompetency(comp models.Competency) error
	GetAllCompetencies() ([]models.CompetencyResponse, error)
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
	GetCompetencyByID(id int) (models.CompetencyResponse, error)
}

type Service struct {
	Material   Material
	Competency Competency
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Material:   NewMaterialService(repo.MaterialRepository),
		Competency: NewCompetencyService(repo.CompetencyRepository),
	}
}
