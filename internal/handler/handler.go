package handler

import (
	"devSystem/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		AllowOrigins:     []string{"http://localhost:25502", "http://localhost:82", "http://localhost:3000", "http://localhost:25504", "https://b.service-to.ru"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/refresh", h.refresh)
	}

	materials := api.Group("/materials", h.UserIdentityMiddleware)
	{
		materials.GET("/:id", h.getMaterial)
		materials.GET("", h.getAllMaterials)
		materials.POST("", h.createMaterial)
		materials.PUT("/:id", h.updateMaterial)
		materials.DELETE("/:id", h.deleteMaterial)
	}

	registration := api.Group("/registration", h.UserIdentityMiddleware)
	{
		registration.GET("/organizations", h.getOrganizations)
		registration.POST("/code", h.createRegistrationCode)
		registration.GET("/departments", h.getDepartments)
		registration.DELETE("/code/:code", h.deleteRegistrationCode)
		registration.POST("/organization", h.createOrganization)
		registration.DELETE("/organization/:id", h.deleteOrganization)
	}

	materialsType := api.Group("/materialsType")
	{
		materialsType.GET("/:id", h.getMaterialType)
		materialsType.GET("", h.getAllMaterialTypes)
		materialsType.POST("", h.createMaterialType)
		materialsType.DELETE("/:id", h.deleteMaterialType)
	}

	competencies := api.Group("/competencies")
	{
		competencies.GET("", h.getAllCompetencies)
		competencies.POST("", h.createCompetency)
		competencies.PUT("/:id", h.updateCompetency)
		competencies.DELETE("/:id", h.deleteCompetency)
		competencies.GET("/:id", h.getCompetency)
	}
	courses := api.Group("/courses")
	{
		courses.GET("", h.getAllCourses)
		courses.POST("", h.createCourse)
		courses.GET("/completed", h.getCompletedCourses)

		course := courses.Group("/:course_id")
		{
			course.GET("", h.getCourse)
			course.PUT("", h.updateCourse)
			course.DELETE("", h.deleteCourse)
			courses.GET("/progress/organization", h.getOrganizationCourseProgress)

			course.POST("/complete", h.completeCourse)
			course.GET("/progress", h.getCourseProgress)
			course.GET("/completed", h.isCourseCompleted)
			course.POST("/materials/:material_id/complete", h.markMaterialAsCompleted)
		}

	}
	users := api.Group("/users", h.UserIdentityMiddleware)
	{
		users.GET("/organization", h.getOrganizationUsers)
		users.DELETE("/:email", h.deleteUser)
	}
	return router
}

type ErrorResponse struct {
	Error string `json:"error"`
}
