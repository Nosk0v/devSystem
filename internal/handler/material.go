package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetMaterial godoc
// @Summary Получить материал по ID
// @Description Получение сведений о материале по его ID.
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "ID материала"
// @Success 200 {object} models.MaterialResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials/{id} [get]
func (h *Handler) getMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid material ID from request: %v", c.Param("id"))
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material ID"})
		return
	}

	material, err := h.usecases.GetMaterial(id)
	if err != nil {
		log.Printf("Error fetching material with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching material"})
		return
	}

	if material == nil {
		log.Printf("Material with ID %d not found", id)
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Material not found"})
		return
	}

	log.Printf("Returning material with ID %d: %+v", id, material)
	c.JSON(http.StatusOK, material)
}

// GetAllMaterials godoc
// @Summary Получить все материалы
// @Description Получение списка со всеми материалами.
// @Tags materials
// @Accept json
// @Produce json
// @Success 200 {array} models.MaterialResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials [get]
func (h *Handler) getAllMaterials(c *gin.Context) {
	log.Println("Fetching all materials request received")
	materials, err := h.usecases.GetAllMaterials()
	if err != nil {
		log.Printf("Error fetching all materials: %v", err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching materials"})
		return
	}

	log.Printf("Returning list of materials: %d items", len(materials))
	c.JSON(http.StatusOK, materials)
}

// CreateMaterial godoc
// @Summary Создать материал
// @Description Создание нового материала по входным данным.
// @Tags materials
// @Accept json
// @Produce json
// @Param material body models.CreateMaterialRequest true "Входные данные"
// @Success 201 {object} models.MaterialResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials [post]
func (h *Handler) createMaterial(c *gin.Context) {
	var input models.CreateMaterialRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	material := models.Material{
		Title:        input.Title,
		Description:  input.Description,
		Type:         input.TypeID,
		Content:      input.Content,
		Competencies: input.Competencies,
	}

	materialID, err := h.usecases.CreateMaterial(material)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating material"})
		return
	}

	materialResponse, err := h.usecases.GetMaterial(materialID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching material details"})
		return
	}

	c.JSON(http.StatusCreated, materialResponse)
}

// UpdateMaterial godoc
// @Summary Обновить материал
// @Description Обновление материала по его ID.
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "ID материала"
// @Param material body models.CreateMaterialRequest true "Входные данные"
// @Success 204 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials/{id} [put]
func (h *Handler) updateMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material ID"})
		return
	}

	var input models.CreateMaterialRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	material := models.Material{
		MaterialID:   id,
		Title:        input.Title,
		Description:  input.Description,
		Type:         input.TypeID,
		Content:      input.Content,
		Competencies: input.Competencies,
	}

	err = h.usecases.UpdateMaterial(material)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating material"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteMaterial godoc
// @Summary Удалить материал
// @Description Удаление материала по его ID.
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "ID материала"
// @Success 204 {object} nil
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials/{id} [delete]
func (h *Handler) deleteMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material ID"})
		return
	}

	if err := h.usecases.DeleteMaterial(id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error deleting material"})
		return
	}

	c.Status(http.StatusNoContent)
}
