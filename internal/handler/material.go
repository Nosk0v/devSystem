package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
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

// GetMaterialType godoc
// @Summary Получить тип материала по ID
// @Description Получение сведений о типе материале по его ID.
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "ID типа материала"
// @Success 200 {object} models.MaterialType
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materialsType/{id} [get]
func (h *Handler) getMaterialType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid material type ID from request: %v", c.Param("id"))
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material type ID"})
		return
	}

	materialType, err := h.usecases.GetMaterialType(id)
	if err != nil {
		log.Printf("Error fetching material type with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching material type"})
		return
	}

	if materialType == nil {
		log.Printf("Material type with ID %d not found", id)
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Material type not found"})
		return
	}

	log.Printf("Returning material type with ID %d: %+v", id, materialType)
	c.JSON(http.StatusOK, materialType)
}

// GetAllMaterials godoc
// @Summary Получить все материалы
// @Description Получение списка со всеми материалами.
// @Tags materials
// @Accept json
// @Produce json
// @Success 200 {array} models.MaterialResponse
// @Failure 500 {object} ErrorResponse
// @Router /materialsType [get]
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

// GetAllMaterialTypes godoc
// @Summary Получить все типы материалов
// @Description Получение списка со всеми типами материалов.
// @Tags materials
// @Accept json
// @Produce json
// @Success 200 {array} models.MaterialType
// @Failure 500 {object} ErrorResponse
// @Router /materials [get]
func (h *Handler) getAllMaterialTypes(c *gin.Context) {
	materials, err := h.usecases.GetAllMaterialTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching material types"})
		return
	}

	c.JSON(http.StatusOK, materials)
}

// CreateMaterial godoc
// @Summary Создать материал
// @Description Создание нового материала по входным данным с указанием компетенций.
// @Tags materials
// @Accept json
// @Produce json
// @Param material body models.CreateMaterialRequest true "Входные данные"
// @Success 201 {object} models.MaterialResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /materials [post]
func (h *Handler) createMaterial(c *gin.Context) {
	var input models.Material
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now().UTC()
	}

	materialID, err := h.usecases.CreateMaterial(input)
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

	var input models.Material
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	input.MaterialID = id
	err = h.usecases.UpdateMaterial(input)
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
