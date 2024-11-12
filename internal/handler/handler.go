package handler

import (
	"devSystem/internal/usecase"
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) getMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	material, err := h.usecases.GetMaterial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching material"})
		return
	}

	if material == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *Handler) getAllMaterials(c *gin.Context) {
	materials, err := h.usecases.GetAllMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching materials"})
		return
	}

	c.JSON(http.StatusOK, materials)
}

func (h *Handler) createMaterial(c *gin.Context) {
	var input models.Material
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.usecases.CreateMaterial(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating material"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) updateMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	var input models.Material
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	input.MaterialID = id

	if err := h.usecases.UpdateMaterial(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating material"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	if err := h.usecases.DeleteMaterial(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting material"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) getAllCompetencies(c *gin.Context) {
	competencies, err := h.usecases.GetAllCompetencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching competencies"})
		return
	}

	c.JSON(http.StatusOK, competencies)
}

func (h *Handler) createCompetency(c *gin.Context) {
	var input models.Competency
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.usecases.CreateCompetency(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating competency"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) updateCompetency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid competency ID"})
		return
	}

	var input models.Competency
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	input.CompetencyID = id

	if err := h.usecases.UpdateCompetency(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating competency"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteCompetency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid competency ID"})
		return
	}

	if err := h.usecases.DeleteCompetency(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting competency"})
		return
	}

	c.Status(http.StatusNoContent)
}
