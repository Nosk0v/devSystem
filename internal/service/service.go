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
	DeleteMaterialType(id int) error
}

type Competency interface {
	CreateCompetency(comp models.Competency) error
	GetAllCompetencies() ([]models.CompetencyResponse, error)
	UpdateCompetency(comp models.Competency) error
	DeleteCompetency(id int) error
	GetCompetencyByID(id int) (models.CompetencyResponse, error)
}

type Account interface {
	Get(email string) (*models.Account, error)
}

type Auth interface {
	SignIn(input *models.SignInInput, accountPassword string) error
}

type JWTToken interface {
	GenerateAccessToken(email string) (string, error)
	GenerateRefreshToken(email string) (string, error)
	ParseToken(tokenString string) (*models.JWTClaims, error)
	GenerateAccessFromRefresh(refreshToken string) (string, error)
}

type Course interface {
	CreateCourse(course models.Course) (int, error)
	LinkCourseWithMaterials(courseID int, materialIDs []int) error
	LinkCourseWithCompetencies(courseID int, competencyIDs []int) error
	GetCourseByID(id int) (models.CourseResponse, error)
	GetAllCourses() ([]models.CourseResponse, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(id int) error
}

type Service struct {
	Material   Material
	Competency Competency
	Account    Account
	Auth       Auth
	JWTToken   JWTToken
	Course     Course
}

func NewService(repo *repository.Repository, config *models.ConfigService) *Service {
	return &Service{
		Material:   NewMaterialService(repo.MaterialRepository),
		Competency: NewCompetencyService(repo.CompetencyRepository),
		Account:    NewAccountService(repo.Account),
		Auth:       NewAuthService(repo.Auth),
		JWTToken:   NewJWTTokenService(config.Server),
		Course:     NewCourseService(repo.CourseRepository),
	}
}
