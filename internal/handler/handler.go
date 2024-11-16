package handler

import (
	"devSystem/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	materials := router.Group("/materials")
	{
		materials.GET("/:id", h.getMaterial)
		materials.GET("", h.getAllMaterials)
		materials.POST("", h.createMaterial)
		materials.PUT("/:id", h.updateMaterial)
		materials.DELETE("/:id", h.deleteMaterial)
	}

	competencies := router.Group("/competencies")
	{
		competencies.GET("", h.getAllCompetencies)
		competencies.POST("", h.createCompetency)
		competencies.PUT("/:id", h.updateCompetency)
		competencies.DELETE("/:id", h.deleteCompetency)
	}

	return router
}
