package repository

import (
	"devSystem/models"
	"github.com/jmoiron/sqlx"
)

type Sources struct {
	Db *sqlx.DB
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

type CourseRepositoryInterface interface {
	CreateCourse(course models.Course) (int, error)
	LinkCourseWithMaterials(courseID int, materialIDs []int) error
	LinkCourseWithCompetencies(courseID int, competencyIDs []int) error
	GetCourseByID(id int) (models.CourseResponse, error)
	UpdateCourse(course models.Course) error
	GetAllCourses() ([]models.CourseResponse, error)
	GetCoursesByOrganization(organizationID int) ([]models.CourseResponse, error)
	DeleteCourse(id int) error
	GetUserCourseProgress(userEmail string, courseID int) ([]int, error)
	MarkMaterialAsCompleted(userEmail string, courseID int, materialID int) error
	CompleteCourse(userEmail string, courseID int) error
	IsCourseCompleted(userEmail string, courseID int) (bool, error)
	GetCompletedCourses(userEmail string) ([]models.CourseResponse, error)
}

type Auth interface {
	CreateAccount(input *models.SignUpInput) error
	AccountExists(email string) (bool, error)
	GetOrganizations() ([]models.Organization, error)
	DeleteRegistrationCode(code string) error
	GetRegistrationCodeInfo(code string) (*models.RegistrationCode, error)
	CreateRegistrationCode(prefix string, isAdmin bool) (string, error)
	MarkRegistrationCodeAsUsed(code string) error
	GetPrefixByOrgID(orgID int) (string, error)
}

type Account interface {
	Get(email string) (*models.Account, error)
}

type Repository struct {
	MaterialRepository   *MaterialRepository
	CompetencyRepository *CompetencyRepository
	Auth
	Account
	CourseRepository *CourseRepository
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		MaterialRepository:   NewMaterialRepository(sources.Db),
		CompetencyRepository: NewCompetencyRepository(sources.Db),
		Auth:                 NewAuthPostgres(sources.Db),
		Account:              NewAccountPostgres(sources.Db),
		CourseRepository:     NewCourseRepository(sources.Db),
	}
}
