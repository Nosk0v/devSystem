package handler

import (
	"devSystem/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetAllCompetencies godoc
// @Summary Получить все компетенции
// @Description Получение списка с компетенциями.
// @Tags competencies
// @Accept json
// @Produce json
// @Success 200 {array} models.CompetencyResponse
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
// @Summary Создать компетенцию
// @Description Создает новую компетенцию по входным данным.
// @Tags competencies
// @Accept json
// @Produce json
// @Param competency body models.Competency true "Входные данные"
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
		input.CreateDate = time.Now().UTC()
	}
	if err := h.usecases.CreateCompetency(input); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating competency"})
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateCompetency godoc
// @Summary Обновить компетенцию
// @Description Обновление сведений о компетенции по её ID.
// @Tags competencies
// @Accept json
// @Produce json
// @Param id path int true "ID компетенции"
// @Param competency body models.Competency true "Входные данные"
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
// @Summary Удалить компетенцию
// @Description Удаление компетенции по её ID.
// @Tags competencies
// @Accept json
// @Produce json
// @Param id path int true "ID компетенции"
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

// GetCompetency godoc
// @Summary Получить компетенцию по ID
// @Description Получение сведений компетенции по её ID.
// @Tags competencies
// @Accept json
// @Produce json
// @Param id path int true "ID компетенции"
// @Success 200 {object} models.CompetencyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /competencies/{id} [get]
func (h *Handler) getCompetency(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid competency ID from request: %v", c.Param("id"))
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid competency ID"})
		return
	}

	competency, err := h.usecases.GetCompetencyByID(id)
	if err != nil {
		log.Printf("Error fetching competency with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching competency"})
		return
	}

	if competency == nil {
		log.Printf("Competency with ID %d not found", id)
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Competency not found"})
		return
	}

	log.Printf("Returning Competency with ID %d: %+v", id, competency)
	c.JSON(http.StatusOK, competency)
}
