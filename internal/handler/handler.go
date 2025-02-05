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

	// Middleware для CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Разрешенные домены
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешенные HTTP-методы
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Разрешенные заголовки
		ExposeHeaders:    []string{"Content-Length"},                          // Заголовки, доступные клиенту
		AllowCredentials: true,                                                // Разрешить отправку cookies
		MaxAge:           12 * time.Hour,                                      // Кэширование CORS настроек
	}))

	materials := router.Group("/materials")
	{
		materials.GET("/:id", h.getMaterial)
		materials.GET("", h.getAllMaterials)
		materials.POST("", h.createMaterial)
		materials.PUT("/:id", h.updateMaterial)
		materials.DELETE("/:id", h.deleteMaterial)

	}
	materialsType := router.Group("/materialsType")
	{
		materialsType.GET("/:id", h.getMaterialType)
		materialsType.GET("", h.getAllMaterialTypes)
		materialsType.POST("", h.createMaterialType)
	}

	// Группы маршрутов для компетенций
	competencies := router.Group("/competencies")
	{
		competencies.GET("", h.getAllCompetencies)
		competencies.POST("", h.createCompetency)
		competencies.PUT("/:id", h.updateCompetency)
		competencies.DELETE("/:id", h.deleteCompetency)
		competencies.GET("/:id", h.getCompetency)
	}

	return router
}

type ErrorResponse struct {
	Error string `json:"error"`
}
