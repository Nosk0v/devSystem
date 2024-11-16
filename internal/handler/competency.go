package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

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

	log.Printf("Received input in handler: %+v\n", input)

	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now()
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
