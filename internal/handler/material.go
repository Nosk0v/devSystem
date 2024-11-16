package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

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
	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now()
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
