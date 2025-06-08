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
	SignUp(input *models.SignUpInput) error
	GetOrganizations() ([]models.Organization, error)
	DeleteRegistrationCode(code string) error
	GetRegistrationCodeInfo(code string) (*models.RegistrationCode, error)
	CreateRegistrationCode(prefix string, isAdmin bool, departmentID *int) (string, error)
	MarkRegistrationCodeAsUsed(code string) error
	GetPrefixByOrgID(orgID int) (string, error)
	CreateOrganizationWithPrefix(name string, prefix string) error
	DeleteOrganization(orgID int) error
	GetUsersByOrganization(orgID int) ([]models.UserResponse, error)
	GetDepartments() ([]models.Department, error)
	DeleteUser(email string) error
}

type JWTToken interface {
	GenerateAccessToken(email string, role int, orgID *int, deptID *int) (string, error)
	GenerateRefreshToken(email string, role int, orgID *int, deptID *int) (string, error)
	ParseToken(tokenString string) (*models.JWTClaims, error)
	GenerateAccessFromRefresh(refreshToken string) (string, error)
}

type Course interface {
	CreateCourse(course models.Course) (int, error)
	LinkCourseWithMaterials(courseID int, materialIDs []int) error
	LinkCourseWithCompetencies(courseID int, competencyIDs []int) error
	GetCourseByID(id int) (models.CourseResponse, error)
	GetAllCourses() ([]models.CourseResponse, error)
	GetCoursesByOrganization(organizationID int) ([]models.CourseResponse, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(id int) error
	GetUserCourseProgress(userEmail string, courseID int) ([]int, error)
	MarkMaterialAsCompleted(userEmail string, courseID int, materialID int) error
	CompleteCourse(userEmail string, courseID int) error
	IsCourseCompleted(userEmail string, courseID int) (bool, error)
	GetCompletedCourses(userEmail string) ([]models.CourseResponse, error)
	GetCoursesByDepartment(orgID int, departmentID int) ([]models.CourseResponse, error)
	GetCourseProgressByOrganization(orgID int) ([]models.UserCourseProgress, error)
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
