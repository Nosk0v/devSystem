package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetMaterial godoc
// @Summary Get a material by ID
// @Description Get material details by its ID
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 200 {object} models.Material
// @Failure 404 {object} ErrorResponse
// @Router /materials/{id} [get]
func (h *Handler) getMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid material ID"})
		return
	}

	material, err := h.usecases.GetMaterial(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching material"})
		return
	}

	if material == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Material not found"})
		return
	}

	c.JSON(http.StatusOK, material)
}

// GetAllMaterials godoc
// @Summary Get all materials
// @Description Get a list of all materials
// @Tags materials
// @Accept json
// @Produce json
// @Success 200 {array} models.Material
// @Failure 500 {object} ErrorResponse
// @Router /materials [get]
func (h *Handler) getAllMaterials(c *gin.Context) {
	materials, err := h.usecases.GetAllMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching materials"})
		return
	}

	c.JSON(http.StatusOK, materials)
}

// CreateMaterial godoc
// @Summary Create a new material
// @Description Create a new material with the input payload
// @Tags materials
// @Accept json
// @Produce json
// @Param material body models.Material true "Material to create"
// @Success 201 {object} models.Material
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /materials [post]
func (h *Handler) createMaterial(c *gin.Context) {
	var input models.Material
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}
	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now()
	}
	if err := h.usecases.CreateMaterial(input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating material"})
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateMaterial godoc
// @Summary Update an existing material
// @Description Update material details by its ID
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Param material body models.Material true "Material to update"
// @Success 200 {object} models.Material
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

	if err := h.usecases.UpdateMaterial(input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating material"})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteMaterial godoc
// @Summary Delete a material
// @Description Delete a material by its ID
// @Tags materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 204 {object} nil
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
