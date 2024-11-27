package repository

import (
	"devSystem/models"
	"github.com/jmoiron/sqlx"
)

type Material interface {
	CreateMaterial(material models.Material) (int, error)
	LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error
	GetMaterialByID(id int) (models.Material, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.Material, error)
}

type Competency interface {
	CreateCompetency(comp models.Competency) error
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
	GetAllCompetencies() ([]models.Competency, error)
}

type Repository struct {
	Material
	Competency
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Material:   NewMaterialRepository(db),
		Competency: NewCompetencyRepository(db),
	}
}
