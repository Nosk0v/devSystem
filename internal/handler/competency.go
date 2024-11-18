package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ErrorResponse описывает ошибку в ответах API
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetAllCompetencies godoc
// @Summary Get all competencies
// @Description Get a list of all competencies
// @Tags competencies
// @Accept json
// @Produce json
// @Success 200 {array} models.Competency
// @Failure 500 {object} ErrorResponse
// @Router /competencies [get]
func (h *Handler) getAllCompetencies(c *gin.Context) {
	competencies, err := h.usecases.GetAllCompetencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching competencies"})
		return
	}

	c.JSON(http.StatusOK, competencies)
}

// CreateCompetency godoc
// @Summary Create a new competency
// @Description Create a new competency with the input payload
// @Tags competencies
// @Accept json
// @Produce json
// @Param competency body models.Competency true "Competency to create"
// @Success 201 {object} models.Competency
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /competencies [post]
func (h *Handler) createCompetency(c *gin.Context) {
	var input models.Competency
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	log.Printf("Received input in handler: %+v\n", input)

	if input.CreateDate.IsZero() {
		input.CreateDate = time.Now()
	}
	if err := h.usecases.CreateCompetency(input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating competency"})
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateCompetency godoc
// @Summary Update an existing competency
// @Description Update competency details by its ID
// @Tags competencies
// @Accept json
// @Produce json
// @Param id path int true "Competency ID"
// @Param competency body models.Competency true "Competency to update"
// @Success 200 {object} models.Competency
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /competencies/{id} [put]
func (h *Handler) updateCompetency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid competency ID"})
		return
	}

	var input models.Competency
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	input.CompetencyID = id

	if err := h.usecases.UpdateCompetency(input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error updating competency"})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteCompetency godoc
// @Summary Delete a competency
// @Description Delete a competency by its ID
// @Tags competencies
// @Accept json
// @Produce json
// @Param id path int true "Competency ID"
// @Success 204 {object} nil
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /competencies/{id} [delete]
func (h *Handler) deleteCompetency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid competency ID"})
		return
	}

	if err := h.usecases.DeleteCompetency(id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error deleting competency"})
		return
	}

	c.Status(http.StatusNoContent)
}
