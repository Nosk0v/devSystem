package repository

import (
	"devSystem/models"
	"github.com/jmoiron/sqlx"
)

type MaterialRepositoryInterface interface {
	CreateMaterial(material models.Material) (int, error)
	LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error
	GetMaterialByID(id int) (models.MaterialResponse, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.MaterialResponse, error)
}

type CompetencyRepositoryInterface interface {
	CreateCompetency(comp models.Competency) error
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
	GetAllCompetencies() ([]models.Competency, error)
}

type Repository struct {
	MaterialRepository   *MaterialRepository
	CompetencyRepository *CompetencyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MaterialRepository:   NewMaterialRepository(db),
		CompetencyRepository: NewCompetencyRepository(db),
	}
}
