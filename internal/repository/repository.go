package repository

import (
	"devSystem/models"
	"github.com/jmoiron/sqlx"
)

type Sources struct {
	db *sqlx.DB
}

type MaterialRepositoryInterface interface {
	CreateMaterial(material models.Material) (int, error)
	LinkMaterialWithCompetencies(materialID int, competencyIDs []int) error
	GetMaterialByID(id int) (models.MaterialResponse, error)
	UpdateMaterial(material models.Material) error
	DeleteMaterial(id int) error
	GetAllMaterials() ([]models.MaterialResponse, error)
	GetMaterialTypeByID(id int) (models.MaterialType, error)
	GetAllMaterialTypes() ([]models.MaterialType, error)
	CreateMaterialType(materialType models.MaterialType) (int, error)
	DeleteMaterialType(int) error
}

type CompetencyRepositoryInterface interface {
	CreateCompetency(comp models.Competency) error
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
	GetAllCompetencies() ([]models.CompetencyResponse, error)
	GetCompetencyById(id int) (models.CompetencyResponse, error)
}

type Auth interface {
}

type Account interface {
	Get(email string) (*models.Account, error)
}

type Repository struct {
	MaterialRepository   *MaterialRepository
	CompetencyRepository *CompetencyRepository
	Auth
	Account
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		MaterialRepository:   NewMaterialRepository(sources.db),
		CompetencyRepository: NewCompetencyRepository(sources.db),
		Auth:                 NewAuthPostgres(sources.db),
		Account:              NewAccountPostgres(sources.db),
	}
}
